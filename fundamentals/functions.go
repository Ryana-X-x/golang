package main

import "fmt"

// when two or more parameters share a type we can omit the type from all but the last
func add(x , y int) int {
	return x + y
}

func swap(x, y string) (string, string){
	return y, x
}

// naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 123))
	a, b := swap("luffy", "zoro")
	fmt.Println(a,b)
	fmt.Println(split(17))
}