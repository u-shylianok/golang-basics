package main

import (
	"fmt"
	"time"
)

func main() {
	timeToSleep := time.Millisecond * 500
	timeToClose := time.Second * 4
	ch1 := make(chan int)
	ch2 := make(chan int)
	var i int // iterator

	// start recievers before timer
	go valueChecker(ch1, ch2)
	go valuePrinter(ch2)

	done := make(chan struct{}) // just to mark that time is up
	go func() {
		time.Sleep(timeToClose)
		done <- struct{}{}
		close(done)
	}()

	for {
		select {
		case <-done:
			close(ch1)
			return
		default:
			ch1 <- i
			i++
			time.Sleep(timeToSleep)
		}
	}
}

func valueChecker(ch1 <-chan int, ch2 chan<- int) {
	for {
		value, ok := <-ch1
		if !ok {
			break
		}
		if value%2 == 0 {
			ch2 <- value
		}
	}
}

func valuePrinter(ch <-chan int) {
	for {
		value, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(value)
	}
}
