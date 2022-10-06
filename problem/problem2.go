package main

import "fmt"

func main() {

	var inputPositiveNumber int

	fmt.Printf("Enter a positive number: ")
	fmt.Scan(&inputPositiveNumber)
	
	
	// check number is negetive

	if inputPositiveNumber <=0 {
		fmt.Println("Number can not be negetive or Zero")
	}

	// procecs 

	if inputPositiveNumber / 3 == 3 {
		fmt.Printf("%v is power of 3", inputPositiveNumber)
	} else {
		fmt.Printf("%v is not power of 3", inputPositiveNumber)
	}

}
