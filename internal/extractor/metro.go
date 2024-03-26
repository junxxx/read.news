package extractor

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	app "github.com/junxxx/read.news/internal"
)

type metroExtractor struct{}

func init() {
	var extractor metroExtractor
	app.Register("metro", extractor)
}

func (metro metroExtractor) Parse(url string) ([]*app.Result, error) {
	var results []*app.Result
	var wg sync.WaitGroup
	content := make(chan *app.Result)

	categories, err := category(url)
	if err != nil {
		return results, err
	}

	for _, cateUrl := range categories {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			Content(url, content)
		}(cateUrl)
	}

	go func() {
		wg.Wait()
		close(content)
	}()

	for c := range content {
		results = append(results, c)
	}

	return results, nil
}

func getHttpBody(url string) (*goquery.Document, error) {
	// Request the HTML page.
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	return goquery.NewDocumentFromReader(resp.Body)
}

func category(homepage string) ([]string, error) {
	var ret []string
	body, err := getHttpBody(homepage)
	if err != nil {
		return nil, err
	}
	body.Find("#widget-category-trending li a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			ret = append(ret, href)
		}
	})
	return ret, nil
}

func Content(url string, content chan<- *app.Result) {
	var result app.Result

	body, err := getHttpBody(url)
	if err != nil {
		log.Println("Conent err:", err)
	}

	title := body.Find("article header h1").Text()
	result.Title = title
	body.Find("article div > p:not(:has(span))").Each(func(i int, s *goquery.Selection) {
		c := strings.Trim(s.Text(), "")
		if len(c) > 0 {
			c = c + "\n\n"
			result.Content = result.Content + c
		}
	})

	content <- &result
}
