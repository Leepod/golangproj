package main

import "fmt"

func main() {
	i := 0
	for {
		if i >= 100000000000 {
			fmt.Println("All Finished")
			break
		}
		i++
	}
}
