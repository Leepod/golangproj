package main

import "fmt"

func main() {

	var crazybool bool = (true && false) || (false && true) || !(false && false)
	fmt.Printf("Is of type: %T \n and has a value of: %b\n", crazybool, crazybool)
}
