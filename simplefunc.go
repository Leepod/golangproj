package main

import "fmt"

func main() {
	howdy("Jane")
	howdy("Caleb")
}

func howdy(name string) {
	fmt.Println(name)
}
