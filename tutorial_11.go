package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func main() {
	fmt.Println("[____________Tutorial 11 ____________]")
	var arr_1 [5]int 
	arr_1[0] =  1
	arr_1[1]=  11
	arr_1[2]= 111
	arr_1[3] =1111
	arr_1[4]=11111
    fmt.Println(arr_1)

	myBook :=[]string{"Math_1st Paper","Math_2nd Paper","English_1st Paper","English_2nd Paper"}
	fmt.Println(myBook) 

	myEmail :=[]string{"alfa@gmail.com","ruhulamin.cs.dev@gmail.com"}
	fmt.Println(myEmail)
	
	// Getting the Length of an Array 

	fmt.Println("The length of my Mail Array List is : " , len(myEmail))
	fmt.Println()

	//inputing marks and adding them all
	// One dimentional Array!
	fmt.Println("Enter All the marks Length")
	scanner_markLen := bufio.NewScanner(os.Stdin)
	scanner_markLen.Scan()
	markLen,_ := strconv.ParseInt(scanner_markLen.Text(),10,64)
	
	mark :=[5]int64{}
	var markSum int64 
	for i:=0;i<int(markLen);i++{
		fmt.Printf("Enter %d th Number \n",i)
		scanner_stdMark := bufio.NewScanner(os.Stdin)
		scanner_stdMark.Scan()
		temp,_ := strconv.ParseInt(scanner_stdMark.Text(),10,64)
		mark[i]=temp 
		
	}
	fmt.Println("All the Marks in a List ")
	
	fmt.Println(mark)
	for i:=0;i<len(mark);i++{
		markSum += mark[i]
	}

	
	fmt.Println()
	fmt.Println("All Marks Summation : ", markSum)


	// 2 Dimentional Array 

	fmt.Println("2D Array--------->")
	arr2D := [2][2]int {
		{1,2},
		{3,4},
	}
	for i :=0;i<2;i++{
		for j:=0;j<2;j++{
			fmt.Printf("%d ",arr2D[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Friens Name ===>")
	var names [2][2]string
	for i :=0;i<2;i++{
		for j:=0;j<2;j++{
			scanner_nm := bufio.NewScanner(os.Stdin)
			scanner_nm.Scan()
			temp_nm := scanner_nm.Text()
			names [i][j] = temp_nm 
		}
	}

	fmt.Println(names)

	

}
