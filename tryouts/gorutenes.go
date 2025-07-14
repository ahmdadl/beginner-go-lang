package main

import (
	"fmt"
	"time"
)

func Sum(nums []int, chnl chan int) {
	sum := 0
	for _, v := range nums {
		sum += v
		time.Sleep(500 * time.Millisecond)
	}

	chnl <- sum
}

func Gorotenses() {
	nums1 := []int{5, 7, 9, 10, 25}
	nums2 := []int{3, 4, 8, 6, 9, 12}
	nums3 := []int{5, 7, 10, 2, 8, 25}

	chnl := make(chan int, len(nums1))
	chnl3 := make(chan int, len(nums3))

	go Sum(nums1, chnl)
	go Sum(nums2, chnl)
	go Sum(nums3, chnl3)

	// x, y := <-chnl, <-chnl

	select {
		case y := <- chnl:
			fmt.Println("y1: ", y)
		case x := <- chnl3:
			fmt.Println("x3: ", x)
		case <- time.After(5 * time.Second):
			fmt.Println("ennnnnnd")
	}

	// fmt.Println(x)
	// fmt.Println(y)

}