package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for num := range randNumsGenerator(10) {
		fmt.Println(num)
	}
}

func randNumsGenerator(n int) (ch <-chan int) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	ch1 := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch1 <- rnd.Int()
		}
		close(ch1)
	}()
	return ch1
}

func randNumsGeneratorBuffered(n int) (ch <-chan int) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	ch1 := make(chan int, n)
	for i := 0; i < n; i++ {
		ch1 <- rnd.Int()
	}
	close(ch1)
	return ch1
}
