package main 

import (
	"fmt"
	"sync"
	
)





func main(){
	
	
	wg := &sync.WaitGroup{}
	wg.Add(2)
	fmt.Println("fine!")
	go func(){
		fmt.Println("hi")
		wg.Done()
	}() // annonimus / lamda function  
	fmt.Println("fine!")
	go func(){
		fmt.Println("world")
		wg.Done()
	}()

	wg.Wait() 
	fmt.Println("fine!")






}