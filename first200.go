package main

import "fmt"

func main() {
	for counter := 0; counter <= 200; counter++ {
		fmt.Printf("Decimal: %d \t Binary: %b \t Hexadecimal: %#X \t UTF-8:%q \n ", counter, counter, counter, counter)
	}
}
