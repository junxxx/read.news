package extractor

import (
	app "github.com/junxxx/read.news/internal"
)

type voaExtractor struct{}

func init() {
	var extractor voaExtractor
	app.Register("voa", extractor)
}

func (voa voaExtractor) Parse(url string) ([]*app.Result, error) {
	var results []*app.Result
	results = append(results, &app.Result{
		Title:   "voa test",
		Content: "voa test content",
	})
	return results, nil
}
