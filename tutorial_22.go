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
	
	arr :=[]int{1,2,3,4,5,6}
	

	for i,value := range arr{
		fmt.Println(i,value)
	} 

		brr:= []int{1,2,3,4,5,6,7,8}
	for x,myval :=range brr{
		fmt.Println(x,myval*2)
	}

	aboutMe := make(map[string]interface{})
	aboutMe["name"]="ruhulamin"
	aboutMe["age"]=200
	aboutMe["email"]="alfa"
	aboutMe["xd"]=22.030

	for key,val :=  range aboutMe{
		fmt.Println(key,val)
	} 
	myIncome := []interface{}{12,12,"asdjfh",225.52}
	for key,val := range myIncome{
		fmt.Println(key,val)
	}






}