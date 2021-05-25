package main

import (
	"fmt"
)

//------------------------------
type Student struct {
	name  string
	marks []int
	age   int
}

// function for accessing the property of the struct Student

func (s Student) getAge() int {
	return s.age
}
func (s Student) getMarks() []int {
	return s.marks
}
func (s Student) getEveryThing() Student {

	return s

}
func (s Student) setAge(age int) {
	s.age = age
}
func (s Student) getAverageMarks() float64 {
	arr := s.marks
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr))

}
func (s Student) getAvgByRange() float64 {
	sum := 0
	for _, v := range s.marks {
		sum += v
	}
	return float64(sum) / float64(len(s.marks))
}

type UserVarification struct {
	userName       string
	userEmail      string
	userContact    string
	userLocationId []int
}

func (varification UserVarification) getAllInfoOfUser() UserVarification {
	return varification
}
func (varification UserVarification) getMaxLocation() int {
	location := varification.userLocationId
	max := -1
	for i := 0; i < len(location); i++ {
		if location[i] > max {
			max = location[i]
		}
	}
	return max
}

func main() {
	s1 := Student{"Ruhul Amin", []int{12, 22, 385, 44}, 26}
	fmt.Println(s1.getAge())
	fmt.Println(s1.getMarks())
	fmt.Println(s1.getEveryThing())
	s1.age = 25
	fmt.Println(s1.getEveryThing())
	fmt.Printf("Average  Marks of S1 is : %f\n", s1.getAverageMarks())
	fmt.Printf("Average getting by Range Key : %f\n", s1.getAvgByRange())

	userVar_1 := UserVarification{"MdRuhulAmin", "ruhulamin.cs.dev@gmail.co,", "01322-352864", []int{101, 102, 103, 104}}
	fmt.Println("All the Information of First User :=> ", userVar_1.getAllInfoOfUser())
	fmt.Println("User Name :=> ", userVar_1.getAllInfoOfUser().userName)
	fmt.Println("User Highest Location :=> ", userVar_1.getMaxLocation())

}
