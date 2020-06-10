package main

import "fmt"

func main() {

	// scene_1()
	// scene_2()
	// scene_3()
	bitShifting()
	bitShiftingIota()

}

func endSection() {
	fmt.Println("=============================")
}

//* Iota
func scene_1() {
	const (
		a = iota
		b
		c
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func scene_2() {
	const (
		a = iota
		b
		c
	)
	const (
		d = iota
		e
		f
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func scene_3() {
	const (
		a = iota
		b
		c
		d = iota
		e
		f
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func bitShifting() {
	//! Basic idea of bits [https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827]
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	endSection()
	//! each one has 10 more zeroes than the one before
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	endSection()
}

func bitShiftingIota() {
	const (
		_ = iota
		// kb := 1024
		kb = 1 << (iota * 10)
		// mb := 1024 * kb
		mb = 1 << (iota * 10)
		// gb := 1024 * mb
		gb = 1 << (iota * 10)
	)
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
