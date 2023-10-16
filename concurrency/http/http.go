package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var ready bool

func main() {
	http.HandleFunc("/slow1", slow1)
	http.HandleFunc("/slow2", slow2)
	log.Println("server started at 8080")
	http.ListenAndServe(":8080", nil)
}

func slow1(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	query := func(q string) string {
		time.Sleep(time.Second * 5)
		return "db data " + q + "\n"
	}

	// slow query without context
	msg := query("1")
	log.Println("query 1 completed")

	// check http request cancelled or not ?
	select {
	case <-ctx.Done():
		log.Println("server returned:", ctx.Err())
		return
	default:
	}

	// slow query without context
	msg += query("2")
	log.Println("query 2 completed")

	log.Println("result returned")
	w.Write([]byte(time.Now().Local().String() + "\n" + msg))
}

func slow2(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	chanDb := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)

	qContext := func(q string, sec time.Duration) {
		time.Sleep(time.Second * sec)
		chanDb <- "db data " + q + "\n"
		wg.Done()
	}

	// slow query with context
	go qContext("1", 4)

	// slow query with context
	go qContext("2", 6)

	msg := ""
	go func() {
		for {
			select {
			case qryVal := <-chanDb:
				msg += qryVal
			case <-ctx.Done():
				log.Println("server:", ctx.Err())
				return
			}
		}
	}()

	wg.Wait()
	w.Write([]byte(time.Now().Local().String() + "\n" + msg))
}
