package main

import (
	"fmt"
	"strconv"
	"time"
)

func fanOutStart() {
	// eg: https://gist.github.com/mchirico/df9fad3e7a5ea0c4527a
	// generator
	p1 := generate("p1", time.Millisecond*800)

	// fan-out, one channel to two channel
	out1, out2 := fanOut(p1)

	go func() {
		for {
			select {
			case in1 := <-out1:
				fmt.Printf("consumed from out 1 :%s \n", in1)
			case in2 := <-out2:
				fmt.Printf("consumed from out 2 :%s \n", in2)
			}
		}
	}()

	time.Sleep(time.Second * 8) // wait for few second output only
}

func fanOut(in <-chan string) (<-chan string, <-chan string) {
	out1 := make(chan string)
	out12 := make(chan string)
	go func() {
		for data := range in {
			select {
			case out1 <- data + "- out1":
			case out12 <- data + "- out2":
			}
		}
	}()
	return out1, out12
}

func generate(name string, sleep time.Duration) chan string {
	intChan := make(chan string)
	go func() {
		for i := 0; ; i++ { //  send forevers
			time.Sleep(sleep)
			intChan <- name + "-" + strconv.Itoa(i)
		}
	}()
	return intChan
}
