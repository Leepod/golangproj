package main

import "fmt"

func main() {

	myMap := map[string]string{"foo": "bar", "foo1": "bar1"}
	fmt.Println(myMap)
	fmt.Println(myMap["foo1"])

}
