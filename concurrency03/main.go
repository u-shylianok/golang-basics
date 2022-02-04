package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	a := randNumsGenerator(10, 10)
	b := randNumsGenerator(10, 100000)
	c := randNumsGenerator(10, 100000*100000)

	for num := range joinChannels(a, b, c) {
		fmt.Println(num)
	}
}

func joinChannels(chs ...<-chan int) <-chan int {
	mergeChan := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for _, ch := range chs {
			ch := ch
			wg.Add(1)
			go func() {
				for num := range ch {
					mergeChan <- num
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(mergeChan)
	}()
	return mergeChan
}

func randNumsGenerator(n, max int) (ch <-chan int) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	ch1 := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch1 <- rnd.Intn(max)
		}
		close(ch1)
	}()
	return ch1
}
