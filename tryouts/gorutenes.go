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
	nums1 := []int{5, 7, 9, 10, 15, 25}
	nums2 := []int{3, 4, 8, 6, 9, 12}

	chnl := make(chan int, len(nums1))

	go Sum(nums1, chnl)
	go Sum(nums2, chnl)

	x, y := <-chnl, <-chnl

	fmt.Println(x)
	fmt.Println(y)

}