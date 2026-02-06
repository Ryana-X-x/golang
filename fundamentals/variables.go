package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// the *var* statement can be at package or function level
var (
	c, python, java bool
	ToBe bool = false
	// shifts 1 to 64 positions to the left, the subtracts 1
	// to make all the 0's to 1's. then we get 64 1's bits
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func zero_values() {
	var i int
	var f float64
	var b bool
	var s string
	// %v for values(any type), %q for string and slices also adds ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func type_conversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y^y))
	fmt.Println(f)
}

func type_inference() {
	v := "Hello"
	fmt.Printf("v is of the type %T\n", v)
}

func main() {
	var i int
	
	// := is for short assignment statements, can only be used inside a function
	k := 13
	fmt.Println(i, c, python, java, k)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	zero_values()
	type_conversion()
	type_inference()
}
