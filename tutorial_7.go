package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Println("This is Tutorial 7")
	// comparison
	x := 522
	check_1 := x > 45
	check_2 := x < 88
	check_3 := x == 522
	fmt.Println(check_1, check_2, check_3) 

	//--------------------------------------------

	y := 6.330
	check_4 := float64(x) == y
	fmt.Println(check_4) 


	// ---------------------------  

	check_5 := float64(x) !=y 
	fmt.Println(check_5) 

	// checking by charecter 

	variable_1 := "a"
	variable_2 := "b"
	check_6 := variable_1 > variable_2
	fmt.Println(check_6) 

	check_7 := "a" == "a"
	fmt.Println(check_7) 

	check_8 := "abba"  == "abba"
	fmt.Println(check_8) 
    
	fmt.Println(" Check The two String are Palindrom or Not ")
	
	fmt.Println("Enter The First Number ")
	 scanner_str1 := bufio.NewScanner(os.Stdin)
	 scanner_str1.Scan()
	 str_1 := scanner_str1.Text()
	 
	 fmt.Println("Enter The Second Number")
	 scanner_str2 := bufio.NewScanner(os.Stdin)
	 scanner_str2.Scan()
	 str_2 := scanner_str2.Text()

	 palindrom := str_1 == str_2 
	 fmt.Println("The Two String is " , palindrom)
	 


}
