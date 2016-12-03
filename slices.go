package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[1:3])
	fmt.Println(mySlice[2])
	fmt.Println("myString"[2])

}
