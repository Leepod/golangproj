package main

import "fmt"

func main() {
	n := average(32, 54, 89, 100, 84, 120)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("Type is %T\n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
