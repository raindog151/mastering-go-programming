package main

import (
	"fmt"
	"time"
)

func main() {

	var slice = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}

	fmt.Println(binarySearch(slice, 13, 0, len(slice) - 1))
}

func binarySearch(slice []int, target, lowIndex, highIndex int) int {

	if highIndex < lowIndex || highIndex == lowIndex{
		return -1
	}

	middle := int((lowIndex + highIndex) / 2)

	fmt.Printf("Target: %d - Middle: %d - Low: %d - High: %d\n", target, middle, lowIndex, highIndex)

	time.Sleep(time.Second * 1)

	if slice[middle] > target {
		return binarySearch(slice, target, lowIndex, middle)
	} else if slice[middle] < target {
		return binarySearch(slice, target, middle + 1, highIndex)
	} else {
		return middle
	}

}
