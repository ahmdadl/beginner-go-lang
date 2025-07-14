package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if (t == nil) {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chnl1 := make(chan int)
	chnl2 := make(chan int)

	go func() {
		Walk(t1, chnl1)
		defer close(chnl1)
	}()

	go func() {
		Walk(t2, chnl2)
		defer close(chnl2)
	}()

	for {
		v1, ok1 := <- chnl1
		v2, ok2 := <- chnl2

		if ok1 != ok2 || v1 != v2 {
			return  false
		}

		if ok1 {
			fmt.Println("v1: ", v1)
		}
		
		if ok2 {
			fmt.Println("v2: ", v2)
		}

		if !ok1 && !ok2 {
			break
		}
	}

	return true
}
