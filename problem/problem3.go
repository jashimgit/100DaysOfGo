package main

import "fmt"

func main() {
	var inputValue int

	fmt.Print("enter a number: ")
	fmt.Scan(&inputValue);
	fmt.Printf("you have entered %d \n", inputValue)

	// check for negetive value
	if inputValue <=0 {
		fmt.Println("Input integer can not be negetive")
	}

	if ( inputValue & (inputValue - 1 ) == 0){
		fmt.Println("%d is power of 2", inputValue)
	} else {
		fmt.Println("%d is not power of 2", inputValue)
	}
}

/*
	n int = 20

	(n & (n-1) == 0 )
	20 & (20 - 1 ) == O


*/