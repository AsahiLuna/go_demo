package main

import "fmt"

func main() {
	ids := []int{32, 123, 43, 6342, 324, 2, 34, 123, 512, 3, 51, 34, 5, 233, 314}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Luna": "luna@alpgo.cc", "Mea": "mea@alpgo.cc", "Asahi": "asahi@alpgo.cc"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
