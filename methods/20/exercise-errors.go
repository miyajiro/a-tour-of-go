package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	prev := 0.0
	z := 1.0
	eps := 1e-6

	if x < 0 {
		err := ErrNegativeSqrt(x)
		return x, err
	}

	for math.Abs(z-prev) > eps {
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
