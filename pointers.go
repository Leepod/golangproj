package main

import "fmt"

func main() {

	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *string = &a

	fmt.Println(b)
	fmt.Println(*b)
}
