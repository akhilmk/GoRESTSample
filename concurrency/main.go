package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var ready bool

func main() {
	fmt.Println("main start")
	doSomeWorkWithCondSignal()
	//doSomeWorkWithCondBroadcast()
	fmt.Println("main end")

}
func doSomeWorkWithCondBroadcast() {
	var wg sync.WaitGroup
	wg.Add(5)
	syncCall := sync.NewCond(&sync.Mutex{})

	// make all gouroutine wait until it get broadcast signal to start working.
	waitForSignal(func() {
		getHttpStatus("https://www.google.com")
		wg.Done()
	}, syncCall)
	waitForSignal(func() {
		getHttpStatus("https://kaviraj.me/understanding-condition-variable-in-go/")
		wg.Done()
	}, syncCall)
	waitForSignal(func() {
		getHttpStatus("https://www.reddit.com")
		wg.Done()
	}, syncCall)
	waitForSignal(func() {
		getHttpStatus("https://www.facebook.com")
		wg.Done()
	}, syncCall)
	waitForSignal(func() {
		getHttpStatus("https://www.youtube.com/watch?v=NIvSQCwcots")
		wg.Done()
	}, syncCall)

	// do anything more here...

	fmt.Println("Website status checkin starting....")
	syncCall.Broadcast() // start all waiting go routine to check website status
	wg.Wait()            // wait for all go routine to complete.
	fmt.Println("Website status checkin completed....")
}
func waitForSignal(fn func(), syncCalls *sync.Cond) {
	go func() {
		syncCalls.L.Lock()
		syncCalls.Wait()
		syncCalls.L.Unlock()
		fn()
	}()
}
func getHttpStatus(url string) {
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("HTTP call failed:", err)
		return
	}

	fmt.Println("HTTP URL:", url, ", Status:", response.StatusCode)

}

func doSomeWorkWithCondSignal() {
	cond := sync.NewCond(&sync.Mutex{})
	go someWorkWithCond(cond)
	workIteration := 0
	cond.L.Lock()
	for !ready {
		workIteration++
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Printf("work done cond, iterations taken:%d \n", workIteration)
}
func someWorkWithCond(cond *sync.Cond) {
	// doing work
	rand.Seed(time.Now().UnixNano())
	waitTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(waitTime)

	ready = true
	cond.Signal()
}
