package main 

import "fmt"

func main(){

	// binary representation of a number 
    fmt.Println("Binary Representation : ")
	decimal_1 := 1020 
	fmt.Printf("The Binary is : %b",decimal_1) // use Printf function 


	fmt.Println()
	float_1 := 1023.2201
	fmt.Printf("the binary is: %b",float_1)  
	fmt.Println()


	// octal representation of a number 

	fmt.Println("Octal Representation :")
    fmt.Println()
	decimal_2 := 22511
	fmt.Printf("The Octal Number is :%o",decimal_2)

	// hexa decimal 

	fmt.Println()
	decimal_3 := 55221
	fmt.Printf("The hexa Number is :  %X",decimal_3)

	// Decimal of a Number 

	fmt.Println("Decimal Number Convertion")
	decimal_4 := 5522
	fmt.Printf("The Decimal number is %d ",decimal_4) // decimal to decimal no convert   

	// in engineering Format 

	fmt.Println()
	
	decimal_5 := 6.221152222322
	fmt.Printf("The Engineering format is := %e" ,decimal_5 ) 

	// for string 

	string_1 := "this is a string 1 "
	fmt.Printf("the value is %s ",string_1)
	fmt.Println()

	
	// for a qoute 

	
	string_2 := " This is string 2"
	fmt.Printf( "%q" , string_2 )
     
	var string_3 string = fmt.Sprintf("Number : %07d is coooll",45)
	fmt.Println(string_3)


	


	





	 

}