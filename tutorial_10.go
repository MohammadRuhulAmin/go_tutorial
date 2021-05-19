package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"

)

func main(){
	fmt.Println("Tutorial_10....switch statement")
	
		fmt.Println("Insert The Number of Try: ")
		scanner_try := bufio.NewScanner(os.Stdin)
		scanner_try.Scan()
		nTry,_ := strconv.ParseInt(scanner_try.Text(),10,64)
		
		for i:=1;i<=int(nTry);i++{
			fmt.Println("Insert Money~~~~~~~~~!")
		scanner_money := bufio.NewScanner(os.Stdin)
		scanner_money.Scan()
		money,_ := strconv.ParseInt(scanner_money.Text(),10,64)
		switch {
		case money>100:
			fmt.Println("you Have inserted more than 100  taka !")
		case money<100:
			fmt.Println("you Have inserted less than 100 taka")
		case money ==100:
			fmt.Println("You have successfully inserted Money")
			
		default:
			fmt.Println("None of the Case matched ! ")
		}
		if money == 100 {
			break
		}else{
			continue
		}

	}

	

	

	
	


}