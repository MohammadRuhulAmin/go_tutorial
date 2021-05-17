package main

import "fmt"

func main(){

	// finding out the type 
	var first_number = 232
	fmt.Printf("%T",first_number) // printf is a data formet checker 

	fmt.Println()

	var date = "6-july-2021"
	fmt.Printf("%T",date) // printf is a data formet checker  & must use %T 
	var height = 5.90
	fmt.Println()
	fmt.Printf("%T",height)
	

	// another way to declear variable
	
	myAge := 26
	fmt.Println()
	fmt.Printf(" My Age Type := %T ", myAge) 

	student := true 
	fmt.Println()
	fmt.Printf(" The data type :=  %T" ,student) 

	// declearing a variable 

	// var variable_name variable_type = "  "

}