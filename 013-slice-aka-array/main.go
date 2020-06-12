package main

import "fmt"

//! Array is also known as a slice. Slice has a limit in capacity, array doesn't.
//* a SLICE allows you to group together values of the same TYPE

func main() {
	array()
	example()
	loop()
	sliceTheSlice()
	addToSlice()
	deleteFromSlice()
	makeExample()
	multiDimensional()
}

func endSection() {
	fmt.Println("-----------------")
}

func array() {
	//* Initiate a new array of 5 length
	//* The length is part of its t ype
	var x [5]int
	fmt.Println(x)
	//* setting index 3 to 42
	x[3] = 42
	fmt.Println(x)
	//* Showing Index position
	fmt.Println(x[2])
	//* Showing length of array
	fmt.Println(len(x))
	//*
	fmt.Println(cap(x))

	endSection()
}

//! Composite literal
func example() {
	//* x :] type{values} // composite literal
	x := []int{4, 5, 6, 8, 42}
	fmt.Println(x)

	endSection()
}

//* loop over with range clause, 2 easy methods
func loop() {
	x := []int{4, 5, 6, 8, 42}
	for i, v := range x {
		fmt.Println(i, v)
	}
	//* Another way
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	endSection()
}	

func sliceTheSlice() {
	x := []int{4, 5, 6, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	//* startings from 1
	fmt.Println(x[1:])
	//* starts from 1 & up to but not including 3
	fmt.Println(x[1:3])

	endSection()
}

//! Append
func addToSlice() {
	x := []int{4, 5, 6, 8, 9}
	b := append(x, 10, 11, 12)
	fmt.Println(b)

	x = append(x, 10, 11, 12)
	y := []int{13, 14, 15, 16}
	z := append(x, y...)

	fmt.Println(z)

	endSection()
}

//! Using Append to Delete
func deleteFromSlice() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7}
	x = append(x[:2], x[5:]...)
	fmt.Println(x)

	endSection()
}

//! Built in function `make`
//! make allows us to create slices with predertermined lengths  & capacity
func makeExample() {
	x := make([]int, 10, 11)
	fmt.Println("Slice", x)
	fmt.Println("Length", len(x))
	fmt.Println("Capacity", cap(x))
	x[0] = 42
	x[9] = 99

	fmt.Println("Slice", x)
	fmt.Println("Length", len(x))
	fmt.Println("Capacity", cap(x))

	
	//* x[10] = 3423 // This won't work because it's out of our range
	x = append(x, 3423)

	fmt.Println("Slice", x)
	fmt.Println("Length", len(x))
	fmt.Println("Capacity", cap(x))

	//* This increases the length to 12 but doubles the capacity from 11 to 22. 
	//* this step also takes more processing power, so it's better to know your capacity 
	//* when declaring it rather than changing it later
	x = append(x, 1000)
	fmt.Println("Slice", x)
	fmt.Println("Length", len(x))
	fmt.Println("Capacity", cap(x))

	endSection()
}

func multiDimensional() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "MoneyPenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	endSection()
}	