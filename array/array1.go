package main

import "fmt"

func main() {
	var arr1 = [2]int{1, 2}
	var arr2 = [...]int { 3, 5, 6, 10}


	fmt.Println(arr1)
	fmt.Println(arr2)
	// access array elements
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])

	// change array elements
	arr1[0] = 50
	fmt.Println("Array 1 elements now:", arr1)

	// find array lenght
	fmt.Println("length of arr1 is: ", len(arr1))
	fmt.Println("lenght of arr2 is: ", len(arr2))
}
