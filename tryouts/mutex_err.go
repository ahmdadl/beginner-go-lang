package main

import (
	"fmt"
	"time"
)

// Counter is safe to use concurrently.
type Counter struct {
	v  map[string]int
}

func (counter *Counter) Increment(key string) {
	counter.v[key]++
}

func (counter *Counter) Value(key string) int {
	return counter.v[key]
}

func GoWithoutMutex() {
	counter := Counter{v: make(map[string]int)}

	for i:= 0; i < 200; i++ {
		go counter.Increment("withoutMutex")
	}

	time.Sleep(time.Second)
	fmt.Println("without mutex: ", counter.Value("withoutMutex"))
}