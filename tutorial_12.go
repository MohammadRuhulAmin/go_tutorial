package main

import (
	"fmt"
)

func main() {
	fmt.Println(".........Array Slices..........")
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var s_arr_1 []int = arr[:]
	//[:] indecating slice syntax and s_arr[] is a slice of another array
	// and dont initialize any array length for slicing

	fmt.Println(s_arr_1)
	var s_arr_2 []int = arr[0:4]
	fmt.Println(s_arr_2)

	var s_arr_3 []int = arr[2:4]
	fmt.Println(s_arr_3)
	var s_arr_4 []int = arr[0:2]
	fmt.Println(s_arr_4)
	var s_arr_5 []int = arr[:3]
	fmt.Println(s_arr_5)

	fmt.Println(cap(s_arr_5)) // capacity of a slice array to hold the array value

	// adding element in a slice

	s_arr_6 := append(s_arr_5, 101)
	fmt.Println("After Append of 101 : ", s_arr_6)
	s_arr_7 := append(s_arr_6, 1002)
	fmt.Println("After Append of  1002 : ", s_arr_7)
	a := make([]int, 10)
	fmt.Println(a)

}
