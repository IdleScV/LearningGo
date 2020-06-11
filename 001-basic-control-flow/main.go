package main

import "fmt"

func main() {
	// main is always called to start the application
	fmt.Println("Starting Application")

	app()

	// application ends when it hits the last line of main
	fmt.Println("Closing Application")
}

func app() {
	fmt.Println("This is the app")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("This is the end of the app")
}
