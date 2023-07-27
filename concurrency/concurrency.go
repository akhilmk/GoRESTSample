package concurrency

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var ready bool

func Run() {
	fmt.Println("concurrency start")
	fanInMain()
	//fanOutMain()
	//doSomeWorkWithCondSignal()
	//doSomeWorkWithCondBroadcast()
	//differentSpeed()
	fmt.Println("concurrency end")
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

func differentSpeed() {
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

func fanInMain() {
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

func fanOutMain() {
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
