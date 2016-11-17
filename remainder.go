package main

import "fmt"

func main() {
	for x := 1; x <= 100; x++ {
		if x%2 != 1 {
			fmt.Printf("%d - is odd", x)
		} else {
			fmt.Printf("%d - is even", x)
		}
	}
}
