package main

import (
	"fmt"
)
 
const LoggedInToken string = "trjfkkfkdkdkdkdkdkdkdknffkdjfnfkmdmfm" // is a public

func main() {
	// string
	var username string = "prakash"
	fmt.Println(username)
	fmt.Printf("value of  %T: \n", username)

	// booln
	var isLogIn bool = false
	fmt.Println(isLogIn)
	fmt.Printf("value of  %T: \n", isLogIn)

	// int normal
	var isNumber int = 4
	fmt.Println(isNumber)
	fmt.Printf("value of : %T \n", isNumber)

	// uint 8 uint 16 uint 32 uint 64
	var isUint8 uint16 = 25500
	fmt.Println(isUint8)
	fmt.Printf("value of : %T \n", isUint8)

	// float 32 or float 64
	var isFloat float32 = 466.777777765434456565
	fmt.Println(isFloat)
	fmt.Printf("value of : %T \n", isFloat)

	// default values and some aliases
	var anotherVaraible int 
	fmt.Println(anotherVaraible)
	fmt.Printf("value of : %T \n", anotherVaraible)

	// implicit  no var style

	Myweb := "https://github.com/prakashuniyal01?tab=repositories"
	fmt.Println(Myweb)
	fmt.Printf("value of : %T \n",Myweb)

	Mynum := 223
	fmt.Println(Mynum)
	fmt.Printf("value of : %T \n",Mynum)

	fmt.Println(LoggedInToken)
	fmt.Printf("value of : %T \n",LoggedInToken)
}
