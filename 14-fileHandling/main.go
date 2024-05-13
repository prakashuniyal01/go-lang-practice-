package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)


func main()  {
	fmt.Println("wecome to file handling")


	content := "how to create file in go lang"

	file , err := os.Create("./handleFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length", length)
	file.Close()

	readFile("./handleFile.txt")
}

func readFile(filename string)  {
 	databyte, err := ioutil.ReadFile(filename)
	 checkNilErr(err)

	fmt.Println("Text data inside the databyte", string(databyte))
}


func checkNilErr(err error)  {
	if err != nil {
		panic(err)
	}
}