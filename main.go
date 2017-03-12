package main

import "WebCrawler/crawler"
import "fmt"

func main() {
	allLinks := make(map[string]bool)
	domain := "valenciana.me"
	path := "/"
	start := crawler.Crawler{Domain: domain, Path: path, AllLinks: allLinks}
	start.Crawl()

	for k := range allLinks {
		fmt.Println(k)
	}
}
