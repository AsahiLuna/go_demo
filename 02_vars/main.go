package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Asahi"
	var age int = 37
	const isCool = true

	// Short hand
	// name := "Asahi"
	// email := "asahi@alpgo.cc"

	name, email := "Asahi", "asahi@alpgo.cc"

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", email)
}
