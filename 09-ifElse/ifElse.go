package main

import "fmt"

func main()  {
	fmt.Println("welcome ot if else statement")


	values := 10

	var result string

	if values < 10 {
		result = "values in lass than 10"
	}else if values > 10{
		result = "values begger then 10"
	}else {
		result = "value equel to 10"
	}
	fmt.Println(result)
}