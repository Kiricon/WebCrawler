package main

import "WebCrawler/crawler"

func main() {
	start := crawler.Crawler{"valenciana.me", "/"}
	start.Crawl()
}
