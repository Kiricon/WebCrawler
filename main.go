package main

import "WebCrawler/crawler"

func main() {
	domain := "valenciana.me"
	path := "/"
	start := crawler.Crawler{Domain: domain, Path: path}
	start.Crawl()
}
