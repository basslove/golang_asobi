package main

import (
	"fmt"
	"sync"
)

type Cache interface {
	Get(k string) string
	Set(k string, v string)
}

type myCache struct {
}

func (mc myCache) Get(k string) string {
	return "default v"
}

func (mc myCache) Set(k, v string) {
}

// private
var cache *myCache

// var mutex sync.Mutex
var once sync.Once

func GetCache() Cache {
	once.Do(func() {
		fmt.Println("instance is nil")
		cache = &myCache{}
		fmt.Println("instance create")
	})
	return *cache
}

func exe(ch chan<- interface{}, key, value string) {
	c := GetCache()
	ch <- c
	fmt.Println(c, key, value)
}

func main() {
	fmt.Println("singleton")

	ch := make(chan interface{})
	go exe(ch, "key1", "value1")
	go exe(ch, "key2", "value2")
	go exe(ch, "key3", "value3")
	go exe(ch, "key4", "value4")
	go exe(ch, "key5", "value5")
	<-ch
	<-ch
	<-ch
	<-ch
	<-ch

	fmt.Scanln()
}
