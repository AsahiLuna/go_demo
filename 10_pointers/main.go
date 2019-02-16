package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(a, *b)
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}