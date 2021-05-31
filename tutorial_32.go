package main

import (
	"fmt"
)

func SendCalculation(inf chan int, val int) {
	inf <- val * 2
	inf <- val * 4
	inf <- val * 8

	close(inf) // closing a channel is must !!!
}

func main() {

	channel := make(chan int)
	go SendCalculation(channel, 2)
	
	for val := range channel {
		fmt.Println(val)
	}

}
