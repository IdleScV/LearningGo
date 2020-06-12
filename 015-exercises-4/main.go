package main

import "fmt"

func main() {
	exercise_1()
	exercise_2()
	exercise_3()
	exercise_4()
	exercise_5()
	exercise_6()
	exercise_7()
	exercise_8()
	exercise_9()
	exercise_10()
}

func endSection() {
	fmt.Println("-----------------------")
}

//* Array
func exercise_1() {
	x := [5]int{1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Printf("%T\t", v)
		fmt.Println("index:", i, "value:", v)

	}

	fmt.Printf("%T\n", x)
	endSection()
}

//* Slice
func exercise_2() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
	endSection()
}

func exercise_3() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	endSection()
}

func exercise_4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	endSection()
}

func exercise_5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
	endSection()
}

func exercise_6() {
	x := make([]string, 50, 50)
	x = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
	endSection()
}

func exercise_7() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)
	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
	endSection()
}

func exercise_8() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range m {
		fmt.Println("record:", k)
		for j, val := range v {
			fmt.Printf("\t index position: %v \n \t\t value: %v \n", j, val )
		}
	}
	endSection()
}

func exercise_9(){
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["new_record"] = []string{"I'm Batman", "Kangaroos", "Physical Exams"}
	for k, v := range m {
		fmt.Println("record:", k)
		for j, val := range v {
			fmt.Printf("\t index position: %v \n \t\t value: %v \n", j, val )
		}
	}
	endSection()

}

func exercise_10(){
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	delete(m, `no_dr`)
	for k, v := range m {
		fmt.Println("record:", k)
		for j, val := range v {
			fmt.Printf("\t index position: %v \n \t\t value: %v \n", j, val )
		}
	}
	endSection()
}