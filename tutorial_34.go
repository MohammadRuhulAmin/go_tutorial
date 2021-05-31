package main

import (
	"fmt"
	"time"
)

func sendtoChannel_1(inf chan string, info string) {
	inf <- info
	time.Sleep(time.Millisecond * 10)
}

func main() {
	fmt.Println("https://golangdocs.com/select-statement-in-golang")
	channel_1 := make(chan string)
	go sendtoChannel_1(channel_1, "Md Ruhul Amin")
	// rsv_1 := <-channel_1
	// fmt.Println(rsv_1)
	channel_2 := make(chan string)
	go sendtoChannel_1(channel_2, "ruhulamin@cs.aiub.edu")
	// rsv_2 :=  <- channel_2
	// fmt.Println(rsv_2)

	select {
	case rsv_1 := <-channel_1:
		fmt.Println(rsv_1)
	case rsv_2 := <-channel_2:
		fmt.Println(rsv_2)
	default:
		fmt.Println("This is default! ")
	}

}
