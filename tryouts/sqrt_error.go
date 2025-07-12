package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(err))
}

func Sqrt(x float64) (float64, ErrNegativeSqrt) {
	if x < 0 {
		return  0, ErrNegativeSqrt(x)
	}

	return math.Sqrt(x), ErrNegativeSqrt(0)
}
