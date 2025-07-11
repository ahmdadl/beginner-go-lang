package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var sum int64

	for _, v := range m {
		sum += v
	}

	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64

	for _, v := range m {
		sum += v
	}

	return sum
}

func SumIntsOrFloats[K comparable , V int64 | float64](m map[K]V) V {
	var sum V

	for _, v := range m {
		sum += v
	}

	return sum
}

func SumNumbers[Key comparable, Value Number](nums map[Key]Value) Value {
	var sum Value

	for _, v := range nums {
		sum += v
	}

	return  sum
}

func main() {
	ints := map[string]int64{
		"first":  5,
		"second": 10,
	}

	fmt.Printf("Non-Generic ints sum => %v\n", SumInts(ints))

	floats := map[string]float64{
		"first": 22.5,
		"second": 23.5,
	}

	fmt.Printf("Non-Generic floats sum => %v\n", SumFloats(floats))

	fmt.Println("------------------Generic------------")

	fmt.Printf("Genric ints => %v\n", SumIntsOrFloats[string, int64](ints))
	fmt.Printf("Genric floats => %v\n", SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Genric Numbers ints => %v\n", SumNumbers(ints))
	fmt.Printf("Genric Numbers floats => %v\n", SumNumbers(floats))

}