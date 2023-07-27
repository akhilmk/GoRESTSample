package main

import (
	"github.com/akhilmk/go-samples/concurrency"
	"github.com/akhilmk/go-samples/ds"
	"github.com/akhilmk/go-samples/rest"
)

func main() {
	concurrency.Run()
	ds.Run()
	rest.Run()
}
