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

	domain = formatDomain(domain)

	resp, _ := http.Get(domain)
	body, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	return string(body)
}

func formatDomain(domain string) string {
	if !strings.Contains(domain, "http://") {
		domain = "http://" + domain
	}

	return domain
}
