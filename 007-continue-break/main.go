package main

import "fmt"

func main() {
	x := 0

	for {
		if x > 100 {
			break
		}

		if x%2 != 0 {
			x++

			break
		}

		fmt.Println(x)
		x++

	}

	fmt.Println("Done.")

}
