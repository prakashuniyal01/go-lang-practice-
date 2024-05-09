package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main()  {
		fmt.Println("welcome to using input like user enter string to number convtor")
		
		reader := bufio.NewReader(os.Stdin)

		input, _ := reader.ReadString('\n')


		fmt.Println("input ",input)

		numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

		if err != nil{
			fmt.Println(err)
		}else{
			fmt.Println( "adding number" ,numRating + 1)
		}


}