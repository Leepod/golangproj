package main

import "fmt"

func main() {

	theSlice := make([]int, 0, 4)

	fmt.Println("---------------------")
	fmt.Println(theSlice)
	fmt.Println(len(theSlice))
	fmt.Println(cap(theSlice))
	fmt.Println("---------------------")

	for i := 0; i < 100; i++ {
		theSlice = append(theSlice, i)
		fmt.Println("Len:", len(theSlice), "Capacity:", cap(theSlice), "Value: ", theSlice[i])
	}
}
