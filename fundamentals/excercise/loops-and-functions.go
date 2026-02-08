package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0 

	for i := 0; i < 5; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("iteration", i, ":", z)
	}

	return z
}

func main() {
	fmt.Println(
		"Final: ", sqrt(5),
		math.Sqrt(5),
	)
}
