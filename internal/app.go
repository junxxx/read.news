package internal

import (
	"log"

	"github.com/jasonlvhit/gocron"
	"github.com/junxxx/read.news/internal/cache"
	"github.com/junxxx/read.news/internal/deliver"
	_ "github.com/junxxx/read.news/internal/env"
	"github.com/junxxx/read.news/internal/parser"
	"github.com/junxxx/read.news/internal/util"
)

func Run() {
	// UTC time
	gocron.Every(1).Day().At("00:00").Do(task)
	gocron.Every(1).Day().At("00:15").Do(task)
	gocron.Every(1).Day().At("00:30").Do(task)
	<-gocron.Start()
}

func test() {
	log.Println("test")
}

func task() {
	cache.GetInstance().Expire()
	log.Println("start job")
	cache.GetInstance().Expire()
	date := util.Today()
	if taskDone(date) {
		log.Println("date ", date, "send successfully")
		return
	}
	deliver.DeliverDoc(parser.Parse())
	log.Println("job done!")
}

func taskDone(date string) bool {
	if len(date) <= 0 {
		date = util.Today()
	}
	c := cache.GetInstance()
	return c.Exist(date)
}
