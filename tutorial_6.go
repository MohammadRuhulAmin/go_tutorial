package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "math" // this is a package
)

func main() {
	fmt.Println("--------tutorial 6-----")
	fmt.Println("Arithmatic Operation")
	fmt.Println("//// Enter first Number for Arithmatic Operation \\\\")
	scanner_num1 := bufio.NewScanner(os.Stdin)
	scanner_num1.Scan()
	FirstNumber, _ := strconv.ParseInt(scanner_num1.Text(), 10, 64)

	fmt.Println("Enter The Second Number")
	scanner_num2 := bufio.NewScanner(os.Stdin)
	scanner_num2.Scan()
	SecondNumber, _ := strconv.ParseInt(scanner_num2.Text(), 10, 64)

	multi := FirstNumber * SecondNumber
	sum := FirstNumber * SecondNumber
	sub := FirstNumber - SecondNumber
	div := FirstNumber / SecondNumber
	rem := FirstNumber % SecondNumber

	fmt.Printf("The First Number = %d\n", FirstNumber)
	fmt.Printf("The Second Number = %d\n", SecondNumber)

	fmt.Printf("Multiplication of %d & %d is  %d\n", FirstNumber, SecondNumber, multi)
	fmt.Printf("Subtruction of %d & %d is %d\n", FirstNumber, SecondNumber, sub)
	fmt.Printf("Divition of %d & %d is %d\n", FirstNumber, SecondNumber, div)
	fmt.Printf("Summation of %d & %d is %d\n", FirstNumber, SecondNumber, sum)
	fmt.Printf("Remainder of %d & %d is  %d \n", FirstNumber, SecondNumber, rem)
	

}
