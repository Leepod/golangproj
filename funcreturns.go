package main

import "fmt"

func main() {
	fmt.Println(greet("Caleb", "Johnson"))

}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
