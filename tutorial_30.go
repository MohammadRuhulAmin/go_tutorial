package main  

import(
	"fmt"
)

type Information struct{
	Name string 
	Age int 
	contacts []string
}

func SendInformation (info chan Information,inf Information){
	info <-inf 
}
func main(){
	fmt.Println("Tutorial Link :=> ","https://golangdocs.com/channels-in-golang")
	info_1 := Information{
		Name: "Md Ruhul Amin",
		Age: 26,
		contacts: []string{"01322-352864","ruhulamin.cs.dev@gmail.com"},
	}
	
	infoSend := make(chan Information)
	go SendInformation(infoSend,info_1)
	infoReceive := <- infoSend 
	fmt.Println(infoReceive)
	
}
