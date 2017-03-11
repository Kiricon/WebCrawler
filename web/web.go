package web

import (
	"WebCrawler/parser"
	"net/http"
	"strings"
)

// GetPage - Function to get a web page as a string
func GetPage(domain string) {

	domain = formatDomain(domain)

	resp, _ := http.Get(domain)
	defer resp.Body.Close()
	parser.GetLinks(resp.Body)

}

func formatDomain(domain string) string {
	if !strings.Contains(domain, "http://") {
		domain = "http://" + domain
	}

	return domain
}
