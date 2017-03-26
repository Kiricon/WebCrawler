package main

import (
	"WebCrawler/crawler"
	"WebCrawler/parser"
	"fmt"
	"sync"
)

func main() {

	linkCount := 1

	parser.SetupFile()
	var wg sync.WaitGroup

	allLinks := make(map[string]bool)
	domain := "wikipedia.org"
	path := "/"
	start := crawler.Crawler{Domain: domain, Path: path, AllLinks: allLinks, Wg: &wg, Count: &linkCount}

	wg.Add(1)
	go start.Crawl()

	wg.Wait()
	fmt.Println("Done")

}
