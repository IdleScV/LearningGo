package main

import "fmt"

func main() {
	// basic()
	// add()
	remove()
}

func endSection() {
	fmt.Println("---------------------")
}

func basic() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])

	//* Value not stored in key, so it gets a zero
	fmt.Println(m["Barnabas"])

	//* To check if value exist
	value, doesExist := m["Barnabas"]
	fmt.Println(value)
	fmt.Println(doesExist)

	//* Common method of checking
	if _, ok := m["Barnabus"]; ok {
		fmt.Println("This will print if Barnabus exist")
	}

	if _, ok := m["James"]; ok {
		fmt.Println("This will print if James exist")
	}

	endSection()
}

func add() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	m["Barnabus"] = 33

	//* Now it exists
	if v, ok := m["Barnabus"]; ok {
		fmt.Println(v)
	}

	//* now we're printing everything
	for k, v := range m {
		fmt.Println(k, v)
	}

	endSection()
}

func remove() {
	k := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(k)


	//* delete(mapName, key)
	delete(k, "James")
	
	fmt.Println(k)

	if v, ok := k["Miss Moneypenny"]; ok {
		fmt.Println("value", v)
		delete(k, "Miss Moneypenny")
	}

	fmt.Println("This is the final", k)

}
