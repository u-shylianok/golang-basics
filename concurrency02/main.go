package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go numsGenerator(ch1, 10)
	go numsMultiplier(ch1, ch2)

	for num := range ch2 {
		fmt.Println(num)
	}
}

func numsGenerator(ch1 chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch1 <- i
	}
	close(ch1)
}

func numsMultiplier(ch1 <-chan int, ch2 chan<- int) {
	for num := range ch1 {
		ch2 <- num * num
	}
	close(ch2)
}
