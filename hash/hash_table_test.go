package hash

import (
	"fmt"
	"testing"
)

func TestHashTable(t *testing.T) {
	ht := NewHashTable(10)

	// Insert key-value pairs
	ht.Insert("apple", 100)
	fmt.Println("NewHashTable:", ht)
	ht.Insert("banana", 200)
	fmt.Println("NewHashTable:", ht)
	ht.Insert("orange", 300)
	fmt.Println("NewHashTable:", ht)
	ht.Insert("grape", 400)
	fmt.Println("NewHashTable:", ht)

	// Search for keys
	val, found := ht.Search("banana")
	if found {
		fmt.Println("Found banana:", val)
	} else {
		fmt.Println("Banana not found")
	}

	val, found = ht.Search("apple")
	if found {
		fmt.Println("Found apple:", val)
	} else {
		fmt.Println("Apple not found")
	}

	// Delete a key
	ht.Delete("banana")
	val, found = ht.Search("banana")
	if found {
		fmt.Println("Found banana:", val)
	} else {
		fmt.Println("Banana not found after deletion")
	}
}
