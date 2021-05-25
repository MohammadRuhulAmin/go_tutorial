package main 

import (
	"fmt"
)

type Point struct{
	x int32 
	y int32 
	isOnBox bool 
}
type aboutMe struct{
	userName string 
	userEmail string 
	userContact string 
	userRoll int32 
} 

type Origin struct{
	x int32 
	y int32 
}

type Circle struct{
	radius float32 
	center *Origin 
}


 
func main(){
	var point_1 Point = Point{1,2,true}
	var point_2 Point = Point{-3,-4,false}
	fmt.Println(point_1.x+point_2.y)
	fmt.Println(point_2,point_2.isOnBox) 
	point_3 := Point{225,2125,false}
	fmt.Println(point_3)  

	student_1 := aboutMe{"ruhul amin","ruhul@gmail.com","01322-352864",6}
	fmt.Println("About First Student , " ,student_1)

	// by spleating every information 
	fmt.Println("Student Name :=> ", student_1.userName )
	fmt.Println("Student Email :=> ",student_1.userEmail)
	fmt.Println("Student Contact :=> ",student_1.userContact)
	fmt.Println("Student Roll :=> ",student_1.userRoll)  

	c1 := Circle{4.33,&Origin{1,2}} 
	fmt.Println("Circle Information :=> ",c1) 

	fmt.Println(c1.center.x , c1.center.y)


}
