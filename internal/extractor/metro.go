package extractor

import (
	app "github.com/junxxx/read.news/internal"
)

type metroExtractor struct{}

func init() {
	var extractor metroExtractor
	app.Register("metro", extractor)
}

func (metro metroExtractor) Parse(url string) ([]*app.Result, error) {
	var results []*app.Result
	results = append(results, &app.Result{
		Title:   "metro test title",
		Content: "metro test content",
	})
	return results, nil
}
