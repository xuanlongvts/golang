package main

import (
	"fmt"
	"sync"
	"time"
)

// Way 3. Use Mutex

func main() {
	toilet := Toilet{}
	fmt.Println("1. ", toilet.Get())

	go toilet.Increase()
	time.Sleep(time.Second)

	fmt.Println("2. ", toilet.Get())
}

type Toilet struct {
	i int
	m sync.Mutex
}

func (t *Toilet) Get() int {
	t.m.Lock()
	defer t.m.Unlock()
	return t.i
}

func (t *Toilet) Increase() {
	t.m.Lock()
	defer t.m.Unlock()

	t.i++
}
