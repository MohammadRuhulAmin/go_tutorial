package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	arr  := []string {"Ruhul","Amin","Sakib","Sajid"}
	wg.Add(1)
	go func(){
		for key,val :=range arr{
			fmt.Println(key," ",val)
		}
		wg.Done()
	}()
	
	wg.Wait()


}