package main

import "fmt"

func main(){
	fmt.Println("Starting Application")

	app()
	

	fmt.Println("Closing Application")
}

func app(){
	fmt.Println("This is the app")
	for i:= 0; i<100;i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("This is the end of the app")
}