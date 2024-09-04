package main

import "fmt"

func main() {
	// Arrays
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// Strings
	str := "Hello, Go!"
	fmt.Println("String:", str)

	// Iterate over string
	for i, ch := range str {
		fmt.Printf("Character %d: %c\n", i, ch)
	}
}
