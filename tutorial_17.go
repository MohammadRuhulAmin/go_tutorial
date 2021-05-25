package main 

import (
	"fmt"
)

func changeMyName(name *string){
	*name = "Mohammad"
}
func encryptInformation(userName *string, userEmail *string , userContact *string){
	*userName = "xxxxx-xxx"
	*userEmail = "xxx@xx.xxx"
	*userContact = "xxxx-xxxxxxx"
}

 func main(){
	myNumber := 1101
	fmt.Printf("My Number is %d\n",myNumber) 
	fmt.Println("My Location  is ",&myNumber)// & ampercent 

	myName :="RuhulAmin"
	changeMyName(&myName)
	fmt.Println("Applying Pointer :=>  ",myName) 
	
	// practise 

	userName :="ruhul amin"
	userEmail :="ruhulamin.cs.dev@gmail.com"
	userContact := "01322-352864"
	encryptInformation(&userName,&userEmail,&userContact) 
	fmt.Println(userName)
	fmt.Println(userEmail)
	fmt.Println(userContact) 
	

	






 }