package main

import "fmt"

func main() {

	var val int
	fmt.Printf("Enter a positive number: ")
	fmt.Scan(&val)

	if val%2 != 0 || val <= 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}

}






























// func divideIntoTwo(param int) string {
// 	if param % 2 == 0 && param <=0 {
// 		return "yes"
// 	} else {
// 		return "no"
// 	}
// }

// if val % 2 == 0 && val != 0 {
// 	fmt.Println("yes")
// } else {
// 	fmt.Println("no")
// }
