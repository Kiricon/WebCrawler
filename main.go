package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(getWebPage("http://valenciana.me"))
}

func getWebPage(domain string) string {
	resp, err := http.Get(domain)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}
