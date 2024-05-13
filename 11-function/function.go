package main

import "fmt"

func main() {
	fmt.Println("welcom to function in go lang")

	// result := addTowNumber(2,4)

	// fmt.Println("value of " , result)

	second()

	result, _ := proFunf(2,3,5,6,7 )
	fmt.Println("values of pro is ", result)
}

func second() {
	fmt.Println("my second func")

	result := addTowNumber(2, 4)

	fmt.Println("value of ", result)

}

func proFunf(values ...int) (int ,string ) {  
	// isse ham kite bhi int or str pass kr sakte h 
	total := 0

	for _, val := range values{ 
		// _ denote index ko by pass krne ke liye 
		total += val
	}

	return total, "hello"
}

func addTowNumber(valOne int, valTwo int) int {
	return valOne * valTwo
}
