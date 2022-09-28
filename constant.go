package main

import "fmt"

func main() {
	// constant

	const DISCOUNT int = 10
	const VAT int = 15

	fmt.Println("Congrats, you got ", DISCOUNT, "taka discount and your price will be counted after ", DISCOUNT, "and ", VAT, "deduction")
}