package main

import (
	"bufio"
	"fmt"
	"os"
)

func message(name string) {
	fmt.Println("This is a Message ", name)
}

func main() {
	test_message := message
	test_message("Ruhul amin")

	myTotalEducation_fee := func() int {
		school_fee := 1500000
		college_fee := 252233
		versity_fee := 5211212
		total_fee := school_fee + college_fee + versity_fee
		return total_fee
	}
	fmt.Printf("Total Cost for My Education : %d\n", myTotalEducation_fee())

	// another way

	myIncome := func(income_1, income_2, income_3 int) int {
		tutioni := income_1
		youtube := income_2
		others := income_3
		return tutioni + youtube + others

	}(2522, 5522, 2223)

	fmt.Printf("The Total Income is : %d ", myIncome)

	fmt.Println("User Name ")
	scanner_name := bufio.NewScanner(os.Stdin)
	scanner_name.Scan()
	uname := scanner_name.Text()
	myName := func(name string) string {
		return name + " is validated by Authore !"
	}(uname)

	fmt.Println("The Final Message is : ", myName)
	
	







}
