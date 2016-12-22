package main

import (
	"fmt"
	"sort"
)

func main() {

	ns := []int{1, 3, 4, 100, 93, 27, 83, 34, 22}
	fmt.Println(ns)
	sort.Sort(sort.IntSlice(ns))
	fmt.Println(ns)
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	fmt.Println(ns)
}
