package main 

import (
	"fmt"
)

func main(){

	allTypeDataInArray := []interface{}{"RuhulAmin","CSE",148,3.23,"American International University-Bangladesh","01322-352864"}
	
	for key,val := range allTypeDataInArray{
		fmt.Println(key," -> ",val)
	}




}