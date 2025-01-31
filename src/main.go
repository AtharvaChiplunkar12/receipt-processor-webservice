package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const url1 = ""

func main() {
	fmt.Println(url1)
	res, err := http.Get(url1)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	// _, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(string(data))

	result, _ := url.Parse(url1)
	fmt.Println()

}
