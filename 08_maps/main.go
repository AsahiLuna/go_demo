package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign kv
	// emails["Asahi"] = "asahi@alpgo.cc"
	// emails["Luna"] = "luna@alpgo.cc"
	// emails["Mea"] = "mea@alpgo.cc"

	// Decalre map and add kv
	emails := map[string]string{"Luna": "luna@alpgo.cc", "Mea": "mea@alpgo.cc", "Asahi": "asahi@alpgo.cc"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Asahi"])

	// Delete from map
	delete(emails, "Asahi")
	fmt.Println(emails)
}
