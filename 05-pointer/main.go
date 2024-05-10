package main

import "fmt"

func main()  {
	fmt.Println("welcome to pointer")

	var ptr *int
	fmt.Println(ptr)

	myNumber := 4
	ptrValue :=  &myNumber

	// first line direct rafrence in mamory addres 
	fmt.Println("value of my number value is : ", ptrValue)
	// actual value of ptr store 
	fmt.Println("value of my number value is : ", *ptrValue)

	*ptrValue = *ptrValue *2

	fmt.Println("values of myNumber is :", myNumber)

}