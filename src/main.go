package main

import (
	"github.com/junxxx/read.news/deliver"
	_ "github.com/junxxx/read.news/env"
	"github.com/junxxx/read.news/parser"
)

const docPath = "./doc/"

func main() {
	filenames := parser.Parse()
	deliver.DeliverDoc(filenames)
	// fmt.Println(filenames)
}
