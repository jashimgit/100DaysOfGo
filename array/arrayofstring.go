package main

import "fmt"

func main() {
	var cars = [4]string{"volvo", "BMW", "Ford", "Mazda"}

	fmt.Println(cars)
	// access array elements

	fmt.Println(cars[0])
	fmt.Println(cars[1])
	fmt.Println(cars[2])

	// change elements of an array
	cars[1] = "Ferrari"
	fmt.Println(cars)

	// array initialization
	arr1 := [4]int{}     // not initialized
	arr2 := [4]int{2, 4} // partially initialized
	arr3 := [3]string{}
	arr4 := [4]int{1:20, 3:40} // specific initialization



	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	// find the length of an array
	fmt.Println("the length of arr1",len(arr1))
	fmt.Println("the length of arr2",len(arr2))
	
	fmt.Println("the length of arr3",len(arr3))
	fmt.Println("the length of arr4",len(arr4))
}
