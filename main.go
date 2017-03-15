package main

import "WebCrawler/crawler"
import "fmt"

func main() {
	allLinks := make(map[string]bool)
	domain := "usfigureskating.org"
	path := "/"
	start := crawler.Crawler{Domain: domain, Path: path, AllLinks: allLinks}
	start.Crawl()

	fmt.Println(len(allLinks))
}
