package main

import "fmt"


type User struct{
	Name string
	Email string
	Age int
	Status bool
}

func main()  {
	fmt.Println("welcome to struct keyword in goLang")

	prakash := User{"prakash rawat", "praksh@gmai.com", 24, true}
	fmt.Printf("user values is %+v\n", prakash)
}