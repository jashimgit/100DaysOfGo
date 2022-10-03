package main

import "fmt"

var amount int
var chargePercentage = 10

func calculateDecimalValue(charge int) int { 
	charge % 100
} 


func main(){
	fmt.Scan("Enter total amount:", &amount)
	// convert chargePercentage to decimal value
	var chargeOf100Taka = calculateDecimalValue(chargePercentage)

	fmt.Println(chargeOf100Taka)
}