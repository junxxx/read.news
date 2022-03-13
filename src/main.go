package main

import (
	"fmt"
	"log"

	"github.com/jasonlvhit/gocron"
	"github.com/junxxx/read.news/deliver"
	_ "github.com/junxxx/read.news/env"
	"github.com/junxxx/read.news/parser"
)

const docPath = "./doc/"

func main() {
	fmt.Println("register job")
	// UTC time
	gocron.Every(1).Day().At("01:30").Do(task)
	<-gocron.Start()
}

func task() {
	log.Println("start job")
	deliver.DeliverDoc(parser.Parse())
	log.Println("job done!")
}
