package main 

import (
	"fmt"
)

func SendDataToChannel(ch chan int , value int){
	ch <-value
}
func SendMyName(myName chan string , name string ){
	myName <- name
}

func SendInformation(myInformation chan interface{} , information interface{}){
	myInformation <- information
}

func main(){
	// firstChannel := make(chan int)  // a channel is decleared which sends only int 
	// firstChannel <- 622451  // send the number to the channel 
	// firstChannelReceive := <-firstChannel  // get int from the channel  // <- is a channel operator 

	var v int 
	ch := make(chan int) // create channel 
	go SendDataToChannel(ch,101) // send data via go routine 
	v  = <-ch // receive data from the channel 
	fmt.Println(v) // print 101 

	
	nameChannel := make(chan string)
	go SendMyName(nameChannel,"Md Ruhul Amin")
	receiveMyName := <- nameChannel 
	fmt.Println(receiveMyName) 

	myInformation := make(chan interface{})
	go SendInformation(myInformation , 3.32221222352)
	receiveInformation := <- myInformation
	fmt.Println(receiveInformation)




	

}