package web

import (
	"WebCrawler/parser"
	"net/http"
	"strings"
)

// GetPage - Function to get a web page as a string
func GetPage(url string, allLinks map[string]bool, domain string) {

	url = formatDomain(url)

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	parser.GetLinks(resp.Body, allLinks, domain)
}

func formatDomain(domain string) string {
	if !strings.Contains(domain, "http://") {
		domain = "http://" + domain
	}

	return domain
}
