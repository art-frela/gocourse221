package main

import "fmt"

func main() {
	type foo [2]bool
	var options = [4]foo{{false, false}, {true, true}, {true, false}, {false, true}}

	for _, pq := range options {
		fmt.Printf("p=%t\tq=%t\tp & q = %t\tp | q = %t\tp ^ q = %t\n", pq[0], pq[1], (pq[0] && pq[1]), (pq[0] || pq[1]), (pq[0] != pq[1]))

	}
}
