package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://example.com/"

func main() {
	fmt.Println("Web Requests in Golang")

	response, err := http.Get(url)
	tryCatch(err)

	
	fmt.Printf("Response is of type %T \n", response)

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Status:", response.Status)
	// fmt.Println("Headers:")
	// for key, values := range response.Header {
	// 	fmt.Printf("  %s: %v\n", key, values)
	// }

	bodyBytes, err := io.ReadAll(response.Body)
	tryCatch(err)
	fmt.Println("Body:")
	fmt.Println(string(bodyBytes))

	fmt.Println("Trailer Headers (if any):", response.Trailer)
	defer response.Body.Close()
}

func tryCatch(err error) {
	if err != nil {
		panic(err)
	}
}
