package extractor

import (
	app "github.com/junxxx/read.news/internal"
	"github.com/junxxx/read.news/internal/result"
)

type voaExtractor struct{}

func init() {
	var extractor voaExtractor
	app.Register("voa", extractor)
}

func (voa voaExtractor) Parse(url string) ([]*result.Result, error) {
	var results []*result.Result
	results = append(results, &result.Result{
		Title:   "voa test",
		Content: "voa test content",
	})
	return results, nil
}
