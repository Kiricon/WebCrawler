package main

import (
	"WebCrawler/crawler"
	"WebCrawler/parser"
	"WebCrawler/stats"
	"WebCrawler/terminal"
	"fmt"
	"sync"
)

func main() {

	stats := stats.Stats{TotalLinks: 0, TotalRoutines: 0}
	display := terminal.NewDisplay(&stats)

	parser.SetupFile()
	var wg sync.WaitGroup

	allLinks := make(map[string]bool)
	domain := "usfigureskating.org"
	path := "/"
	start := crawler.Crawler{Domain: domain, Path: path, AllLinks: allLinks, Wg: &wg, Stats: &stats}

	wg.Add(1)
	stats.TotalRoutines++
	go start.Crawl()
	go display.Draw()

	wg.Wait()
	fmt.Println("Done")

}
