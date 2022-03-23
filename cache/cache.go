package cache

import (
	"fmt"

	"github.com/junxxx/read.news/util"
)

type cacheValue bool

type SendLog struct {
	log map[string]cacheValue
}

var sendLog SendLog

func GetInstance() *SendLog {
	if sendLog.log == nil {
		sendLog.log = make(map[string]cacheValue)
	}
	return &sendLog
}

func (c *SendLog) Clean() *SendLog {
	c.log = make(map[string]cacheValue)
	return c
}

func (c *SendLog) Get(k string) cacheValue {
	return c.log[k]
}

func (c *SendLog) Exist(k string) bool {
	_, ok := c.log[k]
	return ok
}

func (c *SendLog) Set(k string) {
	c.log[k] = cacheValue(true)
}

func (c *SendLog) Expire() {
	today := util.Today()
	for k := range c.log {
		if k < today {
			delete(c.log, k)
		}
	}
}

func (c *SendLog) D() {
	for k, v := range c.log {
		fmt.Println(k, v)
	}
}
