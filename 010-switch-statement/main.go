package main

import "fmt"

func main() {
	// switch_1()
	// switch_2()
	switch_3()
}

func switch_1() {
	//! unlike JS, default isn't neccessary
	//* we can use fallthrough to have multiple cases
	//* generally don't use this
	switch {
	//* ^ a missing expression defaults to true
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("This should not print2")
	case (3 == 3):
		fmt.Println("This prints")
		//* Try to comment out the line bellow
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does print?")
	case (7 == 9):
		fmt.Println("Not True 1")
		fallthrough
	case (11 == 9):
		fmt.Println("Not True 2")
		fallthrough
	case (15 == 15):
		fmt.Println("True, but doesn't pass through")
		fallthrough
	//! not neccessary
	default:
		fmt.Println("this is default")
	}
}

func switch_2() {

	switch "Bond" {
	case "Not Bond":
		fmt.Println("should not print")
	case "Bond":
		fmt.Println("Should Print")
	}
}

//! allow multiple cases
func switch_3() {
	n := "Bond"
	switch n {
	case "Not Bond", "Bond":
		fmt.Println("Multiple Cases")
	case "m":
		fmt.Println("Should Print")
	default:
		fmt.Println("No print")
	}
}
