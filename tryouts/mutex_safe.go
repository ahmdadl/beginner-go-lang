package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v map[string]int
}

func (sf *SafeCounter) Inc(key string) {
	sf.mu.Lock()

	sf.v[key]++
	defer sf.mu.Unlock()
}

func (sf *SafeCounter) Value(key string) int {
	sf.mu.Lock()

	defer sf.mu.Unlock()

	return  sf.v[key]
}


func GoWithMutex() {
	counter := SafeCounter{v: make(map[string]int)}
	counterKey := "withMutex"

	for i := 0; i < 2000; i++ {
		go counter.Inc(counterKey)
	}

	time.Sleep(time.Second)
	fmt.Println("with mutex: ", counter.Value(counterKey))
}