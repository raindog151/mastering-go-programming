package main

import "fmt"

func main() {

	c := make(chan int)

	go receiver(c)

	for x := range c {
		fmt.Println(x)
	}

}

func receiver(c chan int) {

	var ints = []int{1, 3, 5, 7, 9, 11, 13, 15, 17}

	for _, v := range ints {
		c <- v
	}

	close(c)

}