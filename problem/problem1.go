
package main

import ("fmt")

func main(){

	// fmt.Println("Hello World!")
	var inputNumber int

	fmt.Print("Enter a number: ")
	fmt.Scan(&inputNumber)

	// check input number is negetive or not
	if inputNumber <= 0 {
		fmt.Println("Number can not be negetive or zero")
	}
	if inputNumber / 2 == 2 {
		fmt.Printf("%v is power of 2", inputNumber)
	} else {
		fmt.Printf("%v is not power of 2", inputNumber)
	}
}