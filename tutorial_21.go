package main 

import (
	"fmt"

)

type Human interface{
	religion() string 
	country()string 
	prayerQuantity()int
}

type Muslin struct{
	relig string 
	con string 
	parayer int 
}

type Hindu struct{
	relig string 
	con string 
	prayer int 
}

func (m  Muslin )religion()string{
	return m.relig
}
func (h Hindu)religiion()string{
	return h.relig
}

func main(){

	M1 := Muslin{
		relig: "ISLAM",
		con: "Bangladesh",
		parayer: 5,
	}
	H1 := Hindu{
		relig: "HINDU",
		con: "India",
		prayer: 0,
	}
	fmt.Println("The Religion for Muslim  is :=> ",M1.religion()) 
	fmt.Println("The Religion for Hindu is  :=> ",H1.religiion())


}