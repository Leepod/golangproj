package main

import "fmt"

func main() {
	h, even := evenhalf(178)
	fmt.Println(h, even)
}

func evenhalf(check int) (int, bool) {
	return check / 2, check%2 == 0
}
