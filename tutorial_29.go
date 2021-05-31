package main 

import (
	"fmt"
)

func SendInformation(information chan interface{} , demoInfo interface{}){
	information <-demoInfo 
}

func main(){

	information := make(chan interface{})
	go SendInformation(information,[]string{"Md Ruhul Amin","16-31438-1","ruhulamin.cs.dev@gmail.com"})
	receiveInformation := <- information 
	fmt.Println("The Information are :=> ", receiveInformation) 
	

	
	


}