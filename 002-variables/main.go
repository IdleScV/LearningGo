package main

import "fmt"



func main() {
	//*---------------------------------------------
	// ! Variables declared need to be used
	//* := short declaration operator (looks like)
	//* JS => const x = 42
	x := 42
	fmt.Println(x)

	x = 99
	fmt.Println(x)

	//* Expression => 100+20
	//* Statement => y := 100 + 20
	y := 100 + 20
	fmt.Println(y)
	//*---------------------------------------------	
	outsideFunc()
	zeroes()

}


//! Global variables
//* "LIMIT SCOPE!" But var is used to declare outside of function body.
var outside = "I'm outside of the function"

func outsideFunc(){
	fmt.Println("again:", outside)
}

//! Zero variables for each types
//* booleans => false, integers => 0, floats => 0.0, strings => ""
//* (functions | pointers | interfaces | slices | channels | maps ) => nil

func zeroes() {
	fmt.Println("Zero variables")
	var true_false bool
	var i int
	var words string
	fmt.Println("boolean", true_false)
	fmt.Println("integer", i)
	fmt.Println("string", words)
}

