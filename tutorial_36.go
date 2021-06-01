package main

import (
	"fmt"
	"sync"
	
)

func myFunc(wg *sync.WaitGroup){
	fmt.Println("go go go") 
	wg.Done()
	
}

func main(){
	fmt.Println("Go waitGroup Tutorial link :=> ","https://www.youtube.com/watch?v=0BPSR-W4GSY") 
	var wg sync.WaitGroup 
	wg.Add(1)

	go myFunc(&wg)
	wg.Wait()
	fmt.Println("Finished ")
	wg.Wait()

}