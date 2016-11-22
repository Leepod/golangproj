package main

import "fmt"

func factorial(x int64) int64 {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(10))
}
