package main

import "fmt"

// ! Ninja Level 2 [https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#]
func main() {
	exercise_1()
	exercise_2()
	exercise_3()
	exercise_4()
	exercise_5()
	exercise_6()
}

func endSection() {
	fmt.Println("------------------")
}

//* Decimal, Binary, Hex
func exercise_1() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	endSection()
}

//* Operators
func exercise_2() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)

	endSection()
}

//* Typed vs UnTyped
func exercise_3() {
	a := 4
	b := int(5)
	fmt.Println(a, b)

	endSection()
}

//* shift bits
func exercise_4() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)

	endSection()
}

//* reinforce raw string literal
func exercise_5() {
	x :=
		`This is a string
	a
	raw
	string`
	fmt.Println(x)

	endSection()
}

//* reinforce iota
func exercise_6() {
	const (
		a = 2017 + iota
		b
		c
		d
	)

	fmt.Println(a, b, c, d)
	endSection()
}

//* Quiz [https://docs.google.com/forms/d/e/1FAIpQLSfjhxXjo0r_OsVys58B1lVs35CLPpneVcjiEKTPsLuQs4mftA/viewform]
func exercise_7() {

	fmt.Println("21/21")

}
