package main

import (
	"fmt"
)

func main()  {
	fmt.Println("welcome to map in go lang")


	language := make(map[string]string)

	language["js"] = "javaScript"
	language["py"] = "python"
	language["go"] = "goLang"
	fmt.Println( language)

	// seprate value in go lang
	// isme []me key deni h jo bnai h 
	fmt.Println( "sort the value ",language["js"])

	// delete 
	delete(language, "go")
	fmt.Println( "delete the value ",language)

	// loop are interesting in golang

	//  _ se kuchh bhi ignore kr sakte h go lang me 
	for key, values  := range language{
		fmt.Printf("for key %v value is %v\n", key, values)
	}
}