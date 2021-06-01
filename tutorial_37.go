package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup 

func main(){
	wg.Add(3)
	arr  := []string {"Ruhul","Amin","Sakib","Sajid"}
	brr :=[]int{1,2,3,4,5}
	go func(){
		for key,val := range brr{
			fmt.Printf("%d => %d =>",key+1,val)
		}
		wg.Done()
		fmt.Println()
	}()
	go func(){
		for key,val :=range arr{
			fmt.Println(key," ",val)
		}
		wg.Done()
		
	}()
	go func(){
		for key,val :=range arr{
			fmt.Println(key," ",val)
		}
		wg.Done()
		
	}()
	
	wg.Wait()


}