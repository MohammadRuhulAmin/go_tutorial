package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	name := "ruhulamin"
	age := 26
	profession := "student"
	if name == "ruhulamin" && age == 26 && profession == "student" {
		fmt.Println("This is Mohammad Ruhul Amin")
	} else {
		fmt.Println("This is Not Ruhul Amin")
	}

	fmt.Println("Enter User Name")
	scanner_name := bufio.NewScanner(os.Stdin)
	scanner_name.Scan()
	userName := scanner_name.Text()

	fmt.Println("Enter Password")
	scanner_password := bufio.NewScanner(os.Stdin)
	scanner_password.Scan()
	password := scanner_password.Text()

	if userName == "RuhulAmin" && password == "25144747" {
		fmt.Println("Loading...")
		fmt.Println("You are Successfully Loggedin")
	} else {
		if userName == "RuhulAmin" && password != "25144747" {
			fmt.Println("Your User Name is Correct But Password is incorrect")
		} else if userName != "Ruhulamin" && password == "25144747" {
			fmt.Println("Your UserName is not Correct & Password is Correct")
		} else {
			fmt.Println("Your User Name and Password Both are Incorrect")

		}
	}

}
