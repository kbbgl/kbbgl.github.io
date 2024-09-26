package main

import "fmt"

func main() {
	integers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, integer := range integers {
		if integer%2 == 0 {
			fmt.Printf("%d is even\n", integer)
		} else {
			fmt.Printf("%d is odd\n", integer)
		}
	}
}
