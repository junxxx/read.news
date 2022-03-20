package main

import (
	"log"

	"github.com/jasonlvhit/gocron"
	"github.com/junxxx/read.news/cache"
	"github.com/junxxx/read.news/deliver"
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
	c := cache.GetInstance()
	return c.Exist(date)
}

func task() {
	log.Println("start job")
	cache.GetInstance().Expire()
	date := util.Today()
	if jobDone(date) {
		log.Println("date ", date, "send successfully")
		return
	}
	deliver.DeliverDoc(parser.Parse())
	log.Println("job done!")
}
