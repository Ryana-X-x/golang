package main

import "fmt"

func main() {
	sum := 0 
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	// the init and post statements are optional
	count := 0
	for ; count < 1000 ; {
		count += 1
	}

	fmt.Println(count)

	// for is go's while
	cnt := 1
	for cnt < 1000 {
		cnt += cnt
	}
	fmt.Println(cnt)	// 1024

	// forever
	// for {
	// 	fmt.Println("forever")
	// }
}
