package main

import (
	"fmt"
	"net/url"
)

const myurl = "http://localhost:8888/tree?"

func main() {
	fmt.Println("URL in golang")

	results, err1 := url.Parse(myurl)
	fmt.Println("Scheme:", results.Scheme)
	fmt.Println("Opaque:", results.Opaque)
	fmt.Println("User:", results.User)
	if results.User != nil {
		username := results.User.Username()
		password, _ := results.User.Password()
		fmt.Println("Username:", username)
		fmt.Println("Password:", password)
	}
	fmt.Println("Host:", results.Host)
	fmt.Println("Hostname:", results.Hostname())
	fmt.Println("Port:", results.Port())
	fmt.Println("Path:", results.Path)
	fmt.Println("RawPath:", results.RawPath)
	fmt.Println("RawQuery:", results.RawQuery)
	fmt.Println("Fragment:", results.Fragment)

	tryCatch(err1)

	querypara := results.Query()
	fmt.Printf("The type of Query Parameter is:%T \n", querypara)
	fmt.Println(querypara["query"])

	for _, val := range querypara {
		fmt.Println("The Parameters are:", val)
	}

	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "\tutcss",
		RawPath: "user=ali",
	}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)

}

func tryCatch(err error) {
	if err != nil {
		panic(err)
	}
}
