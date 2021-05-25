package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printName() {
	fmt.Println("Mohammad Ruhul Amin")
}

func check_userName(name string) {
	if name == "ruhulamin" {
		fmt.Println("The User Name is Correct")
	} else {
		fmt.Println("The User Name is Not Correct")
	}
}

func summation(num1 int64, num2 int64) int64 {
	return num1 + num2
}
func multiplication(num1 int64, num2 int64) int64 {
	return num1 * num2
}
func division(num1 int64, num2 int64) float64 {
	return float64(num1) / float64(num2)
}
func substruction(num1, num2 int64) int64 {
	return num1 - num2
}

// summation , substruction , division , multiplication in one function

func calculator(num1 int, num2 int) (int, int, int, int) {
	return num1 + num2, num1 - num2, num1 * num2, num1 / num2
}

//----------------------------------------------------

// labeling the return type using some variables

func myCalculator(num1 int, num2 int) (add int, sub int, mul int, dev int) {
	add = num1 + num2
	sub = num1 - num2
	mul = num1 * num2
	dev = num1 / num2
	return add, sub, mul, dev
}

//------------------------------------------------------------------
func main() {
	fmt.Println("Tutorial 15 __ function ")
	printName()
	check_userName("skaib")
	check_userName("ruhulamin")
	fmt.Println("Enter First Number : ")
	scanner_num1 := bufio.NewScanner(os.Stdin)
	scanner_num1.Scan()
	num1, _ := strconv.ParseInt(scanner_num1.Text(), 10, 64)
	fmt.Println("Enter Second Number : ")
	scanner_num2 := bufio.NewScanner(os.Stdin)
	scanner_num2.Scan()
	num2, _ := strconv.ParseInt(scanner_num2.Text(), 10, 64)

	fmt.Printf("The Summation of %d %d is %d\n", num1, num2, summation(num1, num2))
	fmt.Printf("The Multiplication of %d %d is %d\n", num1, num2, multiplication(num1, num2))
	fmt.Printf("The Division of %d %d is %f\n", num1, num2, division(num1, num2))
	fmt.Printf("The Substruction of %d %d is %d\n", num1, num2, substruction(num1, num2))

	add, sub, mul, div := calculator(123, 5)
	fmt.Printf("Addition %d \n Substruction %d \n Division %d \n Substruction %d \n", add, sub, mul, div)

	ad_c, sub_c, mul_c, dev_c := myCalculator(1002, 52)
	fmt.Printf("Addition %d \n Substruction %d \n Multiplication %d \n Division %d \n", ad_c, sub_c, mul_c, dev_c)
}
