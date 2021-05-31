package main 

import (
	"fmt"
	"time"
)

func counter(){
	var i int 
	for i =0;i<5;i++{
		time.Sleep(time.Millisecond * 10)
		fmt.Print(i," ")
	}
}


func main(){ 
	fmt.Println("Go Routine Tutorial Link : ","https://golangdocs.com/goroutines-in-golang#what-is-a-goroutine")
  go	counter()
	fmt.Print("\n")
	counter()

	
	go func(){
		for x:= 0 ;x<10;x++{
			fmt.Print(x," ")
			time.Sleep(time.Millisecond * 10)
		}
	}()

	time.Sleep(time.Millisecond * 10) 
	fmt.Print("John", "Doe")

}