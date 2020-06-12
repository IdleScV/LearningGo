package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	basic()
	embedded()
	anonymous()

}

func endSection() {
	fmt.Println("--------------------")
}

func basic() {
	//* always use the field names for clarity
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Bat",
		last:  "Man",
		age:   35,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	endSection()
}

//* This is an embedded struct
type secretAgent struct {
	//* `person` is an unqualified type
	person
	first string
	ltk bool
}

func embedded() {
	sa1 := secretAgent{
		//* unqualified type name acts as the field name 
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
		first: "Michael",
	}

	//* Name collisions will be overwritten & will have to be specified like `sa1.person.fist`
	//* but we can use `sa1.last` if there is no collision
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk, sa1.person.first)
	endSection()
}




//* This is an anonymous struct because we're declaring the structure right as we're defining a variable
//* This is anonymous because it DOESN'T HAVE A NAME
//* Good way to keep code clean
func anonymous(){
	d1 := struct {
		name string
		age int
	}{
		name: "Bob",
		age: 20,
	}


	fmt.Println(d1, d1.name, d1.age)
}

