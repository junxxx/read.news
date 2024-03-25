package app

import (
	"log"
	"sync"
)

type Result struct {
	Title   string
	Content string
}

type Extractor interface {
	Parse(url string) ([]*Result, error)
}

var extractors = make(map[string]Extractor)

func Extract(extractor Extractor, url string, results chan<- *Result) {
	contents, err := extractor.Parse(url)
	if err != nil {
		log.Println(err)
	}
	for _, result := range contents {
		results <- result
	}
}

func Register(site string, extractor Extractor) {
	if _, exists := extractors[site]; exists {
		log.Fatalln(site, "extractor already registered")
	}
	log.Println("Register", site, "extractor")
	extractors[site] = extractor
}

// 1. get Extractor
// 2. run Extractor.parse concurrently
// 3. wait the result of goruntine
// 4. dispatch the result to destination
func Run() {
	sites, err := RetrieveSites()
	if err != nil {
		log.Print("RetrieveSites err", err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(sites))

	for _, site := range sites {
		// get extractor
		extractor, exists := extractors[site.Name]
		if !exists {
			log.Fatalln(site, "extractor doesn't exists")
		}

		go func(extractor Extractor, url string) {
			Extract(extractor, url, results)
			waitGroup.Done()
		}(extractor, site.URI)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Dispatch(results)
}

func Dispatch(results chan *Result) {
	for result := range results {
		log.Printf(result.Title, result.Content)
	}
}
