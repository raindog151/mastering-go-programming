package main

import "fmt"

type Set map[string]struct{}

func main() {

	set := make(Set)

	set["one"] = struct{}{}
	set["two"] = struct{}{}

	printSet(set)

}

func printSet (set Set) {
	for k, _ := range set {
		fmt.Println(k)
	}
}
