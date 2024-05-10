package main

import "fmt"

func main() {
	fmt.Println("welcome to array in go lang")

	fruitlist := [4]string{}
	fruitlist[0] = "papaya"
	fruitlist[2] = "loool"
	fruitlist[3] = "mango"

	fmt.Println("value of", fruitlist)
	fmt.Println("value of", len(fruitlist))

	marks := [6]int{}
	marks[0] = 1

	fmt.Println("marks of ", len(marks))

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	//this is slice method
	fruitList := []string{"apple", "banana", "peach"}
	// fmt.Printf("type of this %T\n", fruitList)

	fruitList = append(fruitList, "mango", "cheeku")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)
	// fmt.Println(len(fruitList))

	// romove value

	index := 2
	fruitList = append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println(fruitList)
}
