package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func whyAtomicIsImportant(){
	fmt.Println("Atomic operations are used when we need to have a shared variable between different goroutines which will be updated by them. If the updating operation is not synchronized then it will create a problem that we saw.Atomic operations solve that problem by synchronizing access to the shared variable and making the output correct.")
}

func main(){
	fmt.Println("https://golangdocs.com/concurrency-in-golang")
	var wg sync.WaitGroup 
	var counter int64
	
	for x:=1;x<=1000;x++{
		wg.Add(1)
	  go func(){
			// counter++
			atomic.AddInt64(&counter, 1) // counter++ 
			wg.Done()
	  }()
	  
	}
	whyAtomicIsImportant()
	wg.Wait()

	fmt.Println(counter)

}