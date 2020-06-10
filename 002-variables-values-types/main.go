package main

import "fmt"

func main() {
	fmt.Println("Start of Notes")
	basic()
	outsideFunc()
	types()
	stringDeclaration()
	zeroes()
	fmtBasicsPrinting()
	customType()
	fmt.Println("End of Notes")
}

// ! Variables declared need to be used
//* JS => const x = 42
func basic() {
	fmt.Println(":= <= short declaration operator")

	x := 42
	fmt.Println(x)

	x = 99
	fmt.Println(x)

	//* Expression => 100+20
	//* Statement => y := 100 + 20
	y := 100 + 20
	fmt.Println(y)

	endSection()
}

//! Global variables
//* "LIMIT SCOPE!" But var is used to declare outside of function body.
var outside = "I'm outside of the function"

func outsideFunc() {
	fmt.Println("Global Variables")

	fmt.Println("This string was declared outside:", outside)

	endSection()
}

//! Zero variables for each types
//* booleans => false, integers => 0, floats => 0.0, strings => ""
//* (functions | pointers | interfaces | slices | channels | maps ) => nil

func zeroes() {
	fmt.Println("Zero variables")

	var true_false bool
	var i int
	var words string

	fmt.Println("boolean zero:", true_false)
	fmt.Println("integer zero:", i)
	fmt.Println("string zero:", words)

	endSection()
}

//! Data Types
//* Can't reassign variables of assigned TYPE to a different TYPE
//* This makes GoLang a STATIC programming language & not a DYNAMIC one like JS
func types() {
	fmt.Println("Data Types")

	y := 42
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	z := "I don't want to be the next batman!"
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	endSection()
}

//! String declaration
func stringDeclaration() {
	fmt.Println(`Declaring strings`)

	str1 := "This is a string"
	str2 :=
		`This 
		is 
			also 
				a 
					string`

	fmt.Println(str1)
	fmt.Println(str2)

	endSection()

}

//! fmt basics [https://golang.org/pkg/fmt/#hdr-Printing]
func fmtBasicsPrinting() {
	y := 42

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Printf(s)

	fmt.Printf("%v\n", y)
	// * \n is an escape characters, more here [https://golang.org/ref/spec#Rune_literals]

	endSection()
}

//! Creating your own types

var a int

//* Created batman type
type batman int

var b batman

func customType() {

	a = 40
	b = 80
	fmt.Println("a:", a)
	fmt.Printf("%T\n", a)
	fmt.Println("b:", b)
	fmt.Printf("%T\n", b)

	//* you can't assign a = b because they are two different types, to make it work, we'll do this...
	//! Conversion
	a = int(b)
	fmt.Println("a:", a)
	fmt.Printf("%T\n", a)

	b = batman(a)
	fmt.Println("b:", b)
	fmt.Printf("%T\n", b)

	endSection()
}

func endSection() {
	fmt.Println("----------------------------------------")
}
