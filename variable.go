package main

import "fmt"


// syntax: var variableName type = value
var student1 string = "John Doe"	// type is string defined

func main() {
	// variable declaration with var keyword

	

	// another shorter syntax to declair variable
	// syntax variableName := value 
	// in this syntax value must be assigned otherwise error occured

	age := 24	// type is inffered

	var _address string = "Uttara, Dhaka"

	// output: Printf, Println, Print
	// input:  Scanf, Scan, Scanln


	/*
		1. var can be declaired inside and outside of functions, := only can be declaired within function
		2. var declaration and value assignment can be done seperately.
	*/



	fmt.Println(student1)
	fmt.Println(age)
	fmt.Print(_address, "\n")
	// using dynamic format

	fmt.Println("student1 is a good boy")



}