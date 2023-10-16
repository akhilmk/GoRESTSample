package main

import (
	"fmt"
	"time"
)

func fanInStart() {
	p1 := generate("p1", time.Second*1)
	p2 := generate("p2", time.Second*2)

	// with fan-in, receiving from one channel.
	oneChan := fanIn(p1, p2)
	go func() {
		for {
			fmt.Println(<-oneChan)
		}
	}()

	time.Sleep(time.Second * 8) // wait for few second output only
}

// taking output from two go routine and putting into one
func fanIn(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case in1 := <-ch1:
				out <- in1
			case in2 := <-ch2:
				out <- in2
			}
		}
	}()
	return out
}
