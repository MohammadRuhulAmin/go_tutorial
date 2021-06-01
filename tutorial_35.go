package main

import (
	"bufio"
	"fmt"
	"os"
	 "time"
)


func gettingInformation(channelName chan string , name string)(chan string){
	channelName <- name 
	time.Sleep(time.Millisecond*10)
	return channelName
}

func useOfSelectStatement(){
	fmt.Print("The select statement is used when multiple goroutines are sending data via channels then the select statement receives data concurrently and chooses the case randomly if all are ready.\n If no case is ready then it simply outputs the default case if the default case is already provided before.\n This shows the versatility of the select statement which is used to selectively get data from multiple provider channels.")
}

func main(){

	fmt.Println("Enter Your Name : ")
	scanner_name := bufio.NewScanner(os.Stdin)
	scanner_name.Scan()
	name :=  scanner_name.Text() 
	name_channel := make(chan string)
	go gettingInformation(name_channel,name)
	myChannelName := <- name_channel
	fmt.Println(myChannelName)
	 email_channel := make(chan string)
	 go gettingInformation(email_channel,"ruhulamin.cs.dev@gmail.com")
	 myEmailChannel := <- email_channel 
	 fmt.Println(myEmailChannel) 
	 go useOfSelectStatement()
	 time.Sleep(time.Millisecond*10)

	

	
	



}
