package main


import "fmt"

func plus ( a int, b int) int {
	return a + b 
}


func plusPlus ( a, b, c int) int {
	return a+b+c
}


func main(){
	res := plus(5,5)
	fmt.Println("5 + 5 = ", res)

	res2 := plusPlus(3,3,3)
	fmt.Println("3 + 3 + 3 = ", res2)
}