package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("welcome to time in golang")
	correntTime := time.Now()
	fmt.Println(correntTime)
	// formate time 
	fmt.Println("time format", correntTime.Format("01-02-2006"))
}