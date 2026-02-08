package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		// negative numbers becomes +ve and gets called with recurssion
		return sqrt(-x) + "i"
	}
	// sprint converts value into a string but does not itself print
	return fmt.Sprint(math.Sqrt(x))
}

// IF with a short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// IF and ELSE
func power(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, because variables declared inside an if short statement
	// are also only available inside any of the else blocks
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 4, 20),
	)
	fmt.Println(
		power(3, 2, 10),
		power(3, 4, 20),
	)
}
