package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------------Tutorial 9----/loop-------")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for x := 1; x < 10; x++ {
		fmt.Println(x)
	}

	inc := 1
	for inc < 5 {
		fmt.Println("from inc ", inc)
		inc++
	}
	for l := 0; l < 10; l++ {
		for n := 0; n < 10; n++ {
			fmt.Printf("%d ", l*n)
		}
		fmt.Println()
	}

	// break and continue
	for b := 10; b < 100; b++ {
		if b == 12 {
			continue
		} else if b == 25 {
			break
		} else {
			fmt.Println("Example of break and continue", b)
		}
	}

	// while loop
	





}
