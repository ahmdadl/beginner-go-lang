package main

import (
	"fmt"
)

func main() {
	// this is defer example with panic and recover
	// tryDeferOnly()

	
	//  this defered func will get recovered panic value
	tryPanicAndDefer()
}

func tryDeferOnly() {
	i := 0

	for {
		if i >= 5 {
			break
		}

		i++
		defer fmt.Println(i)
	}

	fmt.Println("Save new User")

	defer fmt.Println("Send mail")

	fmt.Println("new user saved")
}

func tryPanicOnly(i int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("recovered: %s\n", r)
		}
	}()

	if i > 5 {
		panic("i can not be above 5")
	}

	fmt.Printf("%d\n", i * 2)
}

func tryPanicAndDefer() {
	// this will panic and recoverd
	tryPanicOnly(6)

	// this should work
	tryPanicOnly(3)
}