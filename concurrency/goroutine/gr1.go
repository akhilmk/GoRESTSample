package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func gr1_threeGoroutine() {

	log.Printf("start")
	wg.Add(2)

	stringChan := make(chan string)
	go hello(stringChan)
	go hi(stringChan)

	go func() {
		for {
			s := <-stringChan
			log.Printf("s =%s", s)
		}
	}()

	wg.Wait()
}

func hello(st chan string) {
	for i := 1; i <= 20; i++ {
		time.Sleep(1 * time.Second)
		//log.Printf("hello")
		st <- "hello"
	}

	wg.Done()
}

func hi(st chan string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(3 * time.Second)
		//	log.Printf("hi")
		st <- "hi"
	}
	wg.Done()
}
