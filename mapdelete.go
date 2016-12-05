package main

import "fmt"

func main() {

	myMap := map[string]string{
		"zero":  "Hello",
		"one":   "You",
		"two":   "Smelly",
		"three": "Ape!",
	}

	fmt.Println(myMap)
	fmt.Println(myMap["three"])
	delete(myMap, "three")
	fmt.Println(myMap["three"])
}
