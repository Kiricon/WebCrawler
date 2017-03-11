package crawler

import (
	"WebCrawler/web"
)

// Crawler - A crawler object for each web page.
type Crawler struct {
	Domain string
	Path   string
	links  []string
}

// Crawl - Recursively crawls through unexplored links on pages
func (c *Crawler) Crawl() {
	web.GetPage(c.Domain + c.Path)
}
