package main 

import "fmt"
func main(){
	var students map [string]int = map[string]int{
		"ruhul":148,
		"sakib":155,
		"sajid":25,
	}
	fmt.Println(students)

	contacts := map[string]string{  // map[key]value
		"ruhul":"01322-352864",
		"sakib":"01521-433840",
		"nurul":"01715-316183",
		"atik":"0193-2033521",
	}
	contacts["sakib"] ="sdsdsds"
	 fmt.Println(contacts)

	 // delete an element from the map
	 delete(contacts,"atik")
	 fmt.Println(contacts)

	 no,yes := contacts["ruhul"]
	 fmt.Println(yes,no)
	 fmt.Printf("The Length of the contacts list :=  %d\n",len(contacts))







}