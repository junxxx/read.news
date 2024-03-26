package extractor

import (
	"log"
	"net/http"
	"strings"

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

	url = "https://metro.co.uk/2024/03/25/banksy-spotted-latest-tree-mural-london-20524766/?ico=trending-module_category_news_item-3"
	result, err := Content(url)
	if err != nil {
		log.Println(err)
	}

	results = append(results, result)

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

func Content(url string) (*app.Result, error) {
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
		log.Println("Content:", c)
	})

	return &result, nil
}
