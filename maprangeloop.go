package main

import "fmt"

func main() {

	myGreetings := map[int]string{
		0: "Good Morning!",
		1: "Bonjour!",
		2: "Buenos Dias!",
		3: "Bongiorno!",
	}
	for key, val := range myGreetings {
		fmt.Println(key, " - ", val)
	}
}
