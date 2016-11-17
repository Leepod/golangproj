package main

import "fmt"

func main() {
	for x := 1; x <= 100; x++ {
		if x%2 == 1 {
			fmt.Printf("%d - is odd\n", x)
		} else {
			fmt.Printf("%d - is even\n", x)
		}
	}
}
