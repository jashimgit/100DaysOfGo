package main

import "fmt"

func main(){

	// create slice by slicing an array
	arr1 := [5]int{10,20,30,40,50}
	mySlice := arr1[2:4]

	fmt.Printf("mySlice = %v\n", mySlice)
	fmt.Printf("length = %d\n", len(mySlice))
	fmt.Printf("capacity = %d\n", cap(mySlice))


	// create slice with make function
	// syntax slice_name := make([]type, length, capacity)
	
	mySlice1 := make([]int, 5, 10);
	fmt.Printf("mySlice1 = %v\n", mySlice1)
	fmt.Printf("mySlice1 = %d\n", len(mySlice1))
	fmt.Printf("mySlice1 = %d\n", cap(mySlice1))
}