package crawler

import (
	"WebCrawler/web"
	"fmt"
)

// Crawler - A crawler object for each web page.
type Crawler struct {
	Domain string
	Path   string
}

// Crawl - Recursively crawls through unexplored links on pages
func (c *Crawler) Crawl() {
	links := web.GetPage(c.Domain + c.Path)
	fmt.Println(links)
}
