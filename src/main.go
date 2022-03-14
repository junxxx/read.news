package main

import (
	"log"

	"github.com/jasonlvhit/gocron"
	"github.com/junxxx/read.news/deliver"
	"github.com/junxxx/read.news/env"
	_ "github.com/junxxx/read.news/env"
	"github.com/junxxx/read.news/parser"
	"github.com/junxxx/read.news/util"
)

const docPath = "./doc/"

func main() {
	// UTC time
	// fault tolerant run task on 9:00, 9:15, 9:30 Beijing timezone
	gocron.Every(1).Day().At("00:00").Do(task)
	gocron.Every(1).Day().At("00:15").Do(task)
	gocron.Every(1).Day().At("00:30").Do(task)
	<-gocron.Start()
}

func jobDone(date string) bool {
	if len(date) <= 0 {
		date = util.Today()
	}
	_, ok := env.SendLog[date]
	return ok
}

func task() {
	log.Println("start job")
	date := util.Today()
	if jobDone(date) {
		log.Println("date ", date, "send successfully")
		return
	}
	deliver.DeliverDoc(parser.Parse())
	log.Println("job done!")
}
