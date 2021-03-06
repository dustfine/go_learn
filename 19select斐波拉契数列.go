package main

import (
	"fmt"
	"runtime"
)

func fibonacci(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Println(num, " ")
		case <-quit:
			runtime.Goexit()
		}
	}
}

func main19() {
	ch := make(chan int)
	quit := make(chan bool)

	go fibonacci(ch, quit)

	x, y := 1, 1
	for i := 0; i < 20; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}
