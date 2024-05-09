package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "welcome to go lang input flield"
	fmt.Println(welcome)


	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" welcome sir enter the rating for your pizza ")



	// comma oK // err oK 

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating me", input)
}