package main

import "fmt"

type link struct {
	value int
	next *link
}

// Creates a linked list.

func main() {

	// This returns a pointer to the memory next where the generated list is created
	firstLink := newLink()
	firstLink.value = 0

	secondLink := newLink()
	secondLink.value = 1

	firstLink.next = secondLink

	// Based on the example below, it appears that next points to the address of the first field.

	fmt.Printf("First  - Root Address: %p - Value: %d - Value_Address: %p - Next_Value: %p - Next: %p\n", &firstLink, firstLink.value, &firstLink.value, &firstLink.next, firstLink.next)
	fmt.Printf("Second - Root Address: %p - Value: %d - Value_Address: %p - Next_Value: %p - Next: %p\n", &secondLink, secondLink.value, &secondLink.value, &secondLink.next, secondLink.next)
	fmt.Printf("%d - %p \n", firstLink.next.value, &firstLink.next.value)
}

func newLink() *link {
	return new(link)
}
