package main

import "fmt"

func main() {

	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	var n int
	fmt.Print("Please Enter Number: ")
	fmt.Scan(&n)
	fmt.Println(&n)
	fmt.Println(half(n))
}
