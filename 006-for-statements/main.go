package main

import "fmt"

//* Go By Example [https://gobyexample.com/for]

func main() {
	// loop_1()
	loop_2()
	// loop_3()
	// loop_4()
}

// * Basic Loops
func loop_1() {
	// for init; condition; post {}
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

}

// * Nested loops
func loop_2() {
	for i := 0; i <= 5; i++ {
		fmt.Printf("outer loop: %d\n", i)
		for j := 0; j < 5; j++ {
			if j%2 == 0 {
				continue
			}

			fmt.Printf("--> inner loop: %d\n", j)
		}
	}

}

//* "while" statement
func loop_3() {
	x := 1
	// for condition {}
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("Done")
}

//* another iteration of 'for'
func loop_4() {
	x := 1
	// for {}
	for {

		if x > 9 {
			break
		}

		fmt.Println(x)
		x++

	}
	fmt.Println("Done")
}
