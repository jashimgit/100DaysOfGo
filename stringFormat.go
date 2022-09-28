package main

import "fmt"

func main() {

	var txt = "Hello"

	fmt.Printf("%s\n", txt) // Hello as plain text
	fmt.Printf("%q\n", txt) // "Hello" as a double-quoted string

	fmt.Printf("%x\n", txt)	// prints hex value
	fmt.Printf("%x\n", txt)	// prints hex value

	fmt.Printf("% X\n", txt)
}
