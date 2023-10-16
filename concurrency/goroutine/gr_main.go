package main

import (
	"fmt"
)

func main() {
	fmt.Println("concurrency start")
	gr1_threeGoroutine()
	gr2_differentSpeed()
	fmt.Println("concurrency end")
}
