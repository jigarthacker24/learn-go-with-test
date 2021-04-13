package main

import "sync"

type Counter struct {
	l     sync.Mutex
	Count int
}

func (c *Counter) Inc() {
	c.l.Lock()
	defer c.l.Unlock()
	c.Count++
}

func (c *Counter) Value() int {
	return c.Count
}

func NewCounter() *Counter {
	return &Counter{}
}
