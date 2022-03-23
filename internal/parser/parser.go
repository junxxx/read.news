package parser

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/junxxx/read.news/internal/util"
)

const (
	homeUrl      = "https://learningenglish.voanews.com"
	userAgent    = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36"
	articlesUrl  = homeUrl + "/z/3521"
	listClass    = "media-block-wrap"
	titleClass   = "media-block__content"
	dateClass    = "date date--mb date--size-3"
	articleClass = "wsw"
	articleId    = "articleItems"
	dateFormt    = "January 2, 2006"
)

var list []article

type article struct {
	title    string
	date     string
	url      string
	contents []byte
}

func getArticleUrl(href string) string {
	return homeUrl + href
}

func sameDate(d1, d2 string) bool {
	return d1 == d2
}

func yesterday() string {
	return time.Now().AddDate(0, 0, -1).Format(dateFormt)
}

func getUrlDoc(url string) (*goquery.Document, error) {
	// Request the HTML page.
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	return goquery.NewDocumentFromReader(resp.Body)
}

func parseArticles() {
	doc, err := getUrlDoc(articlesUrl)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("#articleItems").Children().Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title, _ := s.Find("a").Attr("title")
		date := s.Find("span").Text()
		href, _ := s.Find("a").Attr("href")
		url := getArticleUrl(href)
		list = append(list, article{
			title: title,
			date:  date,
			url:   url})
	})
}

func (a *article) addContent() {
	doc, err := getUrlDoc(a.url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("#article-content .wsw").Children().Each(func(i int, s *goquery.Selection) {
		c := strings.Trim(s.Filter("p").Text(), "")
		if len(c) > 0 {
			c = c + "\n\n"
			a.contents = append(a.contents, []byte(c)...)
		}
	})
}

func Parse() []string {
	filenames := make([]string, 0)
	parseArticles()
	for _, a := range list {
		if sameDate(yesterday(), a.date) {
			a.addContent()

			folder := util.Today()
			path := "./" + folder
			os.Mkdir(path, 0755)
			filename := path + "/" + a.title + ".txt"
			err := os.WriteFile(filename, a.contents, 0644)
			if err != nil {
				log.Println(err)
			}
			filenames = append(filenames, filename)
		}
	}
	return filenames
}
