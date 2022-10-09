package main

import "fmt"
// create a slice with []datatype{values}

// slice syntax

//  slice_name := []datatype{values}

// print slice value

func main() {
	// example:

	mySlice := []int{}
	fmt.Printf("A new slice defined name %v\n", mySlice)

	// crate a new slice

	mySlice1 := []int{}
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
	fmt.Println(mySlice1)

	// mySlice2

	mySlice2 := []string{"Go", "Are", "Awesome"}

	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	fmt.Println(mySlice2)
}