package main

import (
	"log"
	"os"

	app "github.com/junxxx/read.news/internal"
	_ "github.com/junxxx/read.news/internal/extractor"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	app.Run()
}
