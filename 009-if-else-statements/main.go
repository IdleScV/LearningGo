package main

import "fmt"

func main() {
	if true {
		fmt.Println("This is true")
	} else if false {
		fmt.Println("This is not true")
	} else {
		fmt.Println("This is neither true nor false")
	}
}
