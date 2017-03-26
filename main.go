package main

import (
	"WebCrawler/crawler"
	"WebCrawler/parser"
	"WebCrawler/stats"
	"fmt"
	"sync"
)

func main() {

	stats := stats.Stats{TotalLinks: 0, TotalRoutines: 0}

	parser.SetupFile()
	var wg sync.WaitGroup

	allLinks := make(map[string]bool)
	domain := "usfigureskating.org"
	path := "/"
	start := crawler.Crawler{Domain: domain, Path: path, AllLinks: allLinks, Wg: &wg, Stats: &stats}

	wg.Add(1)
	stats.TotalRoutines++
	go start.Crawl()

	wg.Wait()
	fmt.Println()
	fmt.Println("Done")

}
