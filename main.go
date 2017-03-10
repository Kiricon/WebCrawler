package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println(getWebPage("valenciana.me"))
}

func getWebPage(domain string) string {

	if !strings.Contains(domain, "http://") {
		domain = "http://" + domain
	}

	resp, err := http.Get(domain)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}
