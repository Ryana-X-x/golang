package main

import (
	"fmt"
	"runtime"
	"time"
)

// switch evaulates cases from top to bottom
func switch_example() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today. ")
	case today + 1:
		fmt.Println("Tomorrow. ")
	case today + 2:
		fmt.Println("In two days. ")
	default:
		fmt.Println("Too far away. ")
	}
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	switch_example()
}
