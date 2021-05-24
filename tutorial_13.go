package main

import "fmt"

func main() {
	fmt.Println("Tutorial on Range ")
	var arr []int = []int{1, 2, 4, 4, 5, 6, 7}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	// Range :
	fmt.Println("Using Range")
	for i, element := range arr {
		fmt.Printf("%d: %d\n", i, element)
	}
	// not using i
	for _, elem := range arr { // _ is anonimus variable !
		fmt.Printf("%d\n", elem)
	}

	for i, _ := range arr { // _ is annonimus variable
		fmt.Printf("%d\n", i)
	}

	for _,element_1 := range arr{
		for _,element_2 :=range arr{
			if element_1 == element_2{
				fmt.Println(element_1)
			}
		}
	}



}
