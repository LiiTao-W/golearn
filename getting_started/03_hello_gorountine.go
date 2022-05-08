package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello goroutine : ", i)
}

func main() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}
