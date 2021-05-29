package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func saySomeThing(s string) {
	for i := 1; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 5)
	}

	wg.Done()
}
func main() {

	wg.Add(2)
	go saySomeThing("Hi")
	go saySomeThing("There")
	wg.Wait()
	

}
