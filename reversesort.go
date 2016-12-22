package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Micah", "Zebulon", "Albert"}
	sort.Strings(s)
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
