package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type worker struct {
	data string
	// keep http client, DB client or any expensive object here for reuse
}

func (w *worker) doWork() {
	for i := 1; i <= 2; i++ {
		// time taking operation
		fmt.Printf("working on %v - %v \n", w.data, i)
		time.Sleep(time.Second * 1)
	}
}

var workerPool chan *worker

func main() {
	// workers make sure specified number of goroutines are running at a time.
	//startWorkersNormal()

	// workers keep pooled reusable data as well.
	startWorkersPool()

}

func startWorkersPool() {

	numberOfWorkers := 10
	workerPool = make(chan *worker, numberOfWorkers)
	for i := 1; i <= numberOfWorkers; i++ {
		workerPool <- &worker{}
	}

	dataToProcess := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	wg := sync.WaitGroup{}
	wg.Add(len(dataToProcess))

	for _, data := range dataToProcess {
		wkr := getWorker(data)
		go func(w *worker) {
			w.doWork()
			queueWorker(w)
			wg.Done()
		}(wkr)
	}
	wg.Wait()
}

func getWorker(data string) *worker {
	worker := <-workerPool
	worker.data = data
	return worker
}

func queueWorker(worker *worker) {
	workerPool <- worker
}

func startWorkersNormal() {
	go func() {
		// count number of goroutines.
		for range time.Tick(500 * time.Millisecond) {
			fmt.Println("goroutines=", runtime.NumGoroutine())
		}
	}()

	// make sure one time 20 goroutines are running.
	threadLimiter := make(chan struct{}, 20)
	for i := 1; true; i++ {
		// block send when threadLimiter size reached
		threadLimiter <- struct{}{}

		go func(i int) {
			time.Sleep(3 * time.Second)
			fmt.Println("work done", i)
			<-threadLimiter // reading, unblock sender if sender is waiting.
		}(i)
	}
}
