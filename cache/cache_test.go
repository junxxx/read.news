package cache

import (
	"fmt"
	"testing"

	"github.com/junxxx/read.news/util"
)

func TestInstance(t *testing.T) {
	k := util.Today()
	a := GetInstance()
	a.Set(k)
	b := GetInstance()
	fmt.Println(b.Get(k))
	if a != b {
		t.Error("GetInstance return two instances")
	}
}

func TestGet(t *testing.T) {
	if GetInstance().Get("01") != false {
		t.Error("Get wrong cache")
	}
}

func TestSet(t *testing.T) {
	GetInstance().Set(util.Today())
	if GetInstance().Get(util.Today()) != true {
		t.Error("Set cache does'nt work")
	}
}

func TestExpire(t *testing.T) {
	fmt.Println("TestExpire")
	c := GetInstance()
	c.Set("2022-01-01")
	c.Set(util.Today())
	c.D()
	c.Expire()
	c.D()
	fmt.Println("Clean")
	c.Clean()
	c.D()
}
