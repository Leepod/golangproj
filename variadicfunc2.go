package main

import "fmt"

func main() {
	data := []float64{24, 25, 26, 27, 28}
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
