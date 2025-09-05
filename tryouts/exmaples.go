package main

import (
	"fmt"
	"maps"
	"slices"
	"time"
)


func Examples() {
	fmt.Println("Hello Go")

	str := "I am String"
	fmt.Println(str)

	const PI = 3.14154545
	fmt.Println(PI)

	i := 0
	for i < 2 {
		fmt.Println("I is ", i)
		i++
	}

	x := 0
	for {
		fmt.Println("X is ", x)

		if x >= 2 {
			break
		}
		x++
	}

	
	if now := time.Now().Month(); now == time.August {
		fmt.Println("month is August")
	} else if now == time.July {
		fmt.Println("month is July")
	} else {
		fmt.Println("month is not known")
	}

	switch time.Now().Month() {
		case time.August :
			fmt.Println("month is August")
		case time.July:
			fmt.Println("month is July")
		default :
			fmt.Println("month is not known")
	}

	arr := [3]int{1, 2, 3}
	arr[1] = 5
	fmt.Println(arr, len(arr))

	sl := []int{5, 6, 7}
	sl[1] = 25

	sl = append(sl, 150, 152, 153)
	c := make([]int, 10, 50)
	copy(c, sl)
	fmt.Println("SL:", sl, " Len: ", len(sl), " Cap: ", cap(sl))
	fmt.Println("C:",c, " Len: ", len(c), " Cap: ", cap(c))

	frag := sl[3:5]
	fmt.Println("Frag", frag, " Len: ", len(frag), " Cap: ", cap(frag))
	frag2 := sl[2:]
	fmt.Println("Frag2", frag2, " Len: ", len(frag2), " Cap: ", cap(frag2))
	frag3 := sl[:4]
	fmt.Println("Frag3", frag3, " Len: ", len(frag3), " Cap: ", cap(frag3))

	if slices.Equal(frag2, frag3) {
		fmt.Println("Frag2 and Frag3 are equals")
	} else {
		fmt.Println("Frag2 and Frag3 are nooooooooot equals")
	}

	m1 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	m1[4] = "D"
	m1[5] = "E"
	fmt.Println("m1:", m1)
	v5, ok := m1[5]
	if ok {
		fmt.Println("V5:", v5)
	}
	delete(m1, 4)
	fmt.Println("m1[4]:", m1[4])

	clear(m1)
	fmt.Println("m1[0]:", m1[0])

	n := map[string]int{"foo": 1, "bar": 2}
    n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n, n2 are Equal")
	}

	SumNums := func (a, b int) (int, error) {
		if a > 5 {
			return 0, fmt.Errorf("Err: a can not be greater than 5")
		}

		return a + b, nil
	}

	sums1, err1 := SumNums(5, 6)
	if err1 == nil {
		fmt.Println("SumNums is 5,6 =", sums1)
	}

	sums2, err2 := SumNums(7, 6)
	if err2 == nil {
		fmt.Println("SumNums is 7,6 =", sums2)
	} else {
		fmt.Println(err2)
	}
	

	vard := func (all ...int) int {
		total := 0
		for i := range all {
			total += i
		}

		return total
	}
	fmt.Println("Sum Of 1, 5, 7, 6, 9, 4 =>", vard(1, 5, 7, 6, 9, 4))

	nextId := func () func() int {
		i := 0

		return func() int {
			i++
			return i
		}
	}

	getNextId := nextId()

	fmt.Println("NextID: ", getNextId())
	fmt.Println("NextID: ", getNextId())
	fmt.Println("NextID: ", getNextId())
	fmt.Println("NextID: ", getNextId())

	fmt.Println("Recr 5", Recr(5))

	str2 := "Hi from Go"
	for _, c := range str2 {
		fmt.Println("c is", string(c))
	}
	
	
}

func Recr(i int) int {
	if i == 0 {
		return 1
	}

	return i * Recr(i - 1)
}