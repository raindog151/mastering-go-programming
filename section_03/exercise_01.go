package main

import "fmt"

type link struct {
	next  *link
	value int
}

type list struct {
	head *link
	tail *link
}

// Creates a linked list.

func main() {

	var ints = []int{1, 3, 5, 7, 9, 11, 13, 15, 17}


	// This returns a pointer to the memory next where the generated list is created
	myList := newList()

	for _, v := range ints {
		myList.Add(v)
	}

	// fmt.Printf("Address: %v - Value: %d - Next: %v\n", *myList.head, myList.head.value, myList.head.next)

	for i := myList.head; i != nil; i = i.next {
		// note that the last value returned doesn't have the memory next for the next item
		// dig in to why
		fmt.Printf("Address: %v - Value: %d - Next: %v\n", *i, i.value, &i.next)
	}

}

func newList() *list {
	return new(list)
}

func (nlist *list) Add(x int) {

	debug := true

	newLink := &link{value: x}

	if debug { fmt.Printf("I'm a new link and my address is %v and my value is %d and i point to %v\n", &newLink, newLink.value, newLink.next) }

	if nlist.head == nil {
		nlist.head = newLink
		//if debug { fmt.Printf("I'm changing nlist.head next %v to point to %v because it's nil\n", nlist.head, &newLink) }
	} else if nlist.tail == nlist.head {
		nlist.head.next = newLink
		//if debug { fmt.Printf("I'm changing nlist.head next %v to point to %v\n because head == tail\n", nlist.head, &newLink) }
	} else if nlist.tail != nil {
		nlist.tail.next = newLink
		//if debug { fmt.Printf("I'm changing nlist.tail next %v to point to %v\n because tail != nil\n", nlist.tail, &newLink) }
	}

	nlist.tail = newLink
	if debug { fmt.Printf("I'm a list and my begin is %v and my end is %v\n", &nlist.head, &nlist.tail) }
}