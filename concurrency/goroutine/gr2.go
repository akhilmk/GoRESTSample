package main

import (
	"fmt"
	"sync"
	"time"
)

func gr2_differentSpeed() {

	intChan := make(chan int)
	wait := sync.WaitGroup{}
	wait.Add(1)

	// producer fast, but wait until consumer is ready to accept
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("producer waiting %d \n", i)
			intChan <- i
		}
		close(intChan)
	}()

	// consumer slow
	go func() {
		for i := range intChan {
			fmt.Printf("consumed %d \n", i)
			time.Sleep(time.Second * 4)
		}
		wait.Done()
	}()
	wait.Wait()
}
