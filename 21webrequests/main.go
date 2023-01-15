package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const myUrl = "http://localhost:8000/get"
const myPostUrl = "http://localhost:8000/post"
const myFormUrl = "http://localhost:8000/postform"

func main() {
	fmt.Println("Web verb with Golang")
	// PerformGetRequest(myUrl)
	// PerformPostJsonRequest(myPostUrl)
	PerformFormRequest(myFormUrl)
}

func PerformGetRequest(myurl string) {
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)
	fmt.Println("Content lenght is:", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("Content is:", string(content))
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is:", byteCount)
	fmt.Println("Response string: ", responseString.String())
}

func PerformPostJsonRequest(myurl string) {
	// Fake json payload

	requestBody := strings.NewReader(`
			{
				"coursename":"Golang",
				"price":0,
				"platform":"learnCodeOnline.in"
			}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content", string(content))
}

func PerformFormRequest(myurl string) {

	// formdata

	data := url.Values{}
	data.Add("name", "Aditya")
	data.Add("course", "Golang")
	data.Add("email", "aditya@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content", string(content))
}
