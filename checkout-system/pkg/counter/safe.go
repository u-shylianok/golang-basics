package counter

import "sync"

type SafeCounter struct {
	sync.Mutex
	count int
}

func NewSafeCounter() Counter {
	return &SafeCounter{}
}

func (c *SafeCounter) Inc() {
	c.Lock()
	defer c.Unlock()

	c.count++
}

func (c *SafeCounter) Dec() {
	c.Lock()
	defer c.Unlock()

	c.count--
}

func (c *SafeCounter) Get() int {
	c.Lock()
	defer c.Unlock()

	return c.count
}

func (c *SafeCounter) Set(value int) {
	c.Lock()
	defer c.Unlock()
	c.count = value
}
