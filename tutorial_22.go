package main 

import (
	"fmt"
)


func abstract (abs interface{}){
	fmt.Println(abs)
}

type Car struct{
	name string 
	quality string 
	details string 
	price float64
}
func main(){

	fmt.Println("Advance of Interface !")
	abstract("ruhul amin")
	abstract(12345)
	abstract(2.33)
	var myMarks []int = []int{120,121,122,123}
	abstract(myMarks) 
	abstract(Car{
		name: "Toyota",
		quality: "Excellent",
		details: "This is an old model car but very smart",
		price: 253.332,
	}) 
	myMap := make(map[string]interface{})
	myMap["name"]="ruhul amin"
	myMap["age"]=2332 
	fmt.Println(myMap) 
	myMap["code"]="22522"
	myMap["id"]=2233
	fmt.Println(myMap) 







}