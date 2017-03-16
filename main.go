package main

import "WebCrawler/parser"

func main() {

	/*
		var wg sync.WaitGroup

		allLinks := make(map[string]bool)
		domain := "usfigureskating.org"
		path := "/"
		start := crawler.Crawler{Domain: domain, Path: path, AllLinks: allLinks}

		wg.Add(1)
		go start.Crawl(&wg)

		wg.Wait()
		fmt.Println("Done")
	*/

	parser.SetupFile()

}
