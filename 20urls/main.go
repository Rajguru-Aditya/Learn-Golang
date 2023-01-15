package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=12345"

func main() {
	fmt.Println("Handling URLs in Golang")
	fmt.Println("URL is: ", myUrl)

	// parsing url
	result, _ := url.Parse(myUrl)

	fmt.Println("Scheme is: ", result.Scheme)
	fmt.Println("Scheme is: ", result.Host)
	fmt.Println("Scheme is: ", result.Path)
	fmt.Println("Scheme is: ", result.Port())
	fmt.Println("Scheme is: ", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Types of Query params are: %T\n", qparams)
	fmt.Println("Query params are: ", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=aditya",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another URL is: ", anotherUrl)
}
