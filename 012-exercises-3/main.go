package main

import "fmt"

// ! Ninja Level 3 [https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#]
func main() {
	// exercise_1()
	// exercise_2()
	// exercise_3()
	// exercise_4()
	// exercise_5()
	// exercise_67()
	// exercise_8()
	// exercise_9()
	exercise_10()
}

func exercise_1() {
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exercise_2() {
	for i := 65; i < 91; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#x\t%#U\n", i, i)
		}
	}
}

func exercise_3() {
	x := 1996
	for x <= 2020 {
		fmt.Println(x)
		x++
	}
}

func exercise_4() {
	x := 1996
	for {
		if x > 2020 {
			break
		}
		fmt.Println(x)
		x++

	}
}

func exercise_5() {
	for i := 11; i < 100; i++ {
		fmt.Println(i % 4)
	}
}

func exercise_67() {
	x := "Bond"
	if x == "not" {
		fmt.Println("Not")
	} else if x == "Bond" {
		fmt.Println("Yes")
	} else {
		fmt.Println("Not 2")
	}
}

func exercise_8() {
	switch {
	case false:
		fmt.Println("Won't Print")
	case (2 == 2):
		fmt.Println("True")
	}
}

func exercise_9() {
	s := "favSport"
	switch s {
	case "favSport":
		fmt.Println("Golf")
	case "favTv":
		fmt.Println("Community")
	case "favCity":
		fmt.Println("Taipei")
	}
}

func exercise_10() {
	fmt.Println("true	false	true	true	false")
}
