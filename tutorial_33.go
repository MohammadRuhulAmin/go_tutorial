package main

import (
	"fmt"
	"time"
)

func Calculation(num1 , num2 int){

	for i:=1;i<=10;i++{
		fmt.Print(i*num1*num2, " ")
	}
	time.Sleep(time.Millisecond *10 )

}


func main() {

	go Calculation(1,2)
	fmt.Println("Tutorial link:=> ","https://golangdocs.com/multiple-goroutines-in-golang")
	go Calculation(1,3)

	time.Sleep(time.Second * 2)

}
