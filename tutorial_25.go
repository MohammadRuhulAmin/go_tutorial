package main

import (
	"fmt"
	"time"

)

func say(s string){
	for i:=0;i<3;i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond*5)
	}
}

func main(){
  
	go say("hi")
	go say("There") 
	fmt.Println("wait for the execution ends...")
	time.Sleep(time.Millisecond*500)
 	

}
