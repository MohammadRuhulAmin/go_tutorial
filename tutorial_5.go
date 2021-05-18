package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("This is tutorial 4")
	scanner := bufio.NewScanner(os.Stdin) // os is operating system and Stdin is input..
	scanner.Scan()                        // scanner.Scan waits for your respons through keyboard
	input := scanner.Text()               // this statement get the text as a string  // this is store in input
	//------------------------------------- variable as a string
	fmt.Printf("You Typed: %q ", input) // what you just writh will show as a message
	fmt.Println()
	//--------------------------------------------------------------------------------

	fmt.Println("what is your school Name")
	scanner_school := bufio.NewScanner(os.Stdin)
	scanner_school.Scan()
	school := scanner_school.Text()

	//---------------------------------------------

	fmt.Println("What is your College Name")
	scanner_college := bufio.NewScanner(os.Stdin)
	scanner_college.Scan()
	college := scanner_college.Text()

	fmt.Println()
	//----------------------------------------
	fmt.Println("What is your University Name")
	scanner_university := bufio.NewScanner(os.Stdin)
	scanner_university.Scan()
	university := scanner_university.Text()

	fmt.Println()
	//---------------------------------------

	// All resulst
	fmt.Printf("School :  %q\n", school)
	fmt.Printf("College : %q\n", college)
	fmt.Printf("University : %q\n", university)
	fmt.Println()

	//To convert a string to int we use strconv module and  its ParseInt object
	// input := strconv.ParseInt(scanner.Text())

	scanner_age := bufio.NewScanner(os.Stdin)
	scanner_age.Scan()
	age := scanner_age.Text()
	myAge, _ := strconv.ParseInt(age, 10, 64) // ParseInt(variable , which base to convert , integersize 32bit or 64bit)
	fmt.Printf("%d", myAge)                   // if  the text cannot be converted into integer  , the function strconv.ParseInt() will
	// retern an error ... 
    





}
