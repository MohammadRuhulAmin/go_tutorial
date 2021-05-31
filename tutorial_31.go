package main 

import (
	"fmt"
)


func SendDatatoChannel(inf chan string,  info string){
	inf <- info 
	close(inf ) // closing chennel
}

func main(){

	channel := make(chan string)
	go SendDatatoChannel(channel,"Md Ruhul Amin")
	err,recv := <- channel
	fmt.Println(recv)
	fmt.Println(err)
	if !recv{
		fmt.Println("Channel is not closed!")
	}

}