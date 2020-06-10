package main

import "fmt"
// ! Ninja Level 1 [https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#]
func main() {
	exercise_1()
	exercise_2()
	exercise_3()
	exercise_4()
	exercise_5()
	exercise_6()
}

func endSection() {
	fmt.Println("----------------------------------------")
}

func exercise_1() {
	x := 42
	y := "James Bond"
	z := true
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)

	endSection()
}

var x int
var y string
var z bool
func exercise_2() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	endSection()
}


func exercise_3(){
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

	endSection()
}

type chickenNugget int
var nugget chickenNugget 

func exercise_4(){
	fmt.Println(nugget)
	fmt.Printf("%T\n", nugget)
	nugget = 42
	fmt.Println(nugget)

	endSection()
}

var newNugget chickenNugget
var nonNugget int

func exercise_5(){
	fmt.Println(newNugget)
	fmt.Printf("%T\n", newNugget)

	fmt.Println(nonNugget)
	fmt.Printf("%T\n", nonNugget)

	newNugget = 23
	nonNugget = int(newNugget)

	fmt.Println(newNugget)
	fmt.Printf("%T\n", newNugget)

	fmt.Println(nonNugget)
	fmt.Printf("%T\n", nonNugget)

	

	endSection()
}

func exercise_6(){
	fmt.Println("Total points 32/36")
}