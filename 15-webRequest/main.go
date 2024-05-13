package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url    = "https://freeapi.app/"
	appID  = "663f55b8dd655e5b9e12e5e3" // Replace with your actual API key
)

func main()  {
	fmt.Println("wecome to handle web request")

 	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is type of %T\n", response)

	defer response.Body.Close()	//coller's responsibilty to allews close the connections

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	content :=	string(databytes)

	fmt.Println(content)
}


func checkNilErr(err error)  {
	if err != nil{
		panic(err)
	}

}