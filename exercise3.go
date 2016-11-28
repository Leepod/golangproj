package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {

	biggest := max(39, 104, 183, 32, 94, 11, 90)
	fmt.Printf("%v\n", biggest)
}
