package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"Tommy", "Tony", "Steve", "Tina", "Tornado"}
	fmt.Println(s)
	sort.StringSlice(s).Sort()
	fmt.Println(s)
}
