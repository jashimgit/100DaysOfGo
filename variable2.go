package main

import ("fmt")

var a int
var b int = 2
var c = 5


func main() {
	a = 1

	sum := a + b + c 

	// multiple variable assignment

	var x, y, z = 10, 20, 30

	// variable declaration in a block
	var (
		m int
		n int = 1
		test string = "test"
	)


	// multi-word variable names
	/*
		Camel Case

		myVariableName = "John"

		pascal case
		MyVariableName = "John"

		Snake case

		my_variable_name = "John"
	
	*/

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(sum)
	fmt.Println(x, y, z)

	fmt.Println("m value is: ", m)
	fmt.Println("n value is: ", n)
	fmt.Println("test value is: ", test)
}