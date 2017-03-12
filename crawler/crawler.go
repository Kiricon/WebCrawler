package crawler

import "WebCrawler/web"

// Crawler - A crawler object for each web page.
type Crawler struct {
	Domain   string
	Path     string
	AllLinks map[string]bool
}

// Crawl - Recursively crawls through unexplored links on pages
func (c *Crawler) Crawl() {
	links := web.GetPage(c.Domain + c.Path)
	c.addToDomainList(links)
}

func (c *Crawler) addToDomainList(links []string) {
	for _, value := range links {
		if !c.AllLinks[value] {
			c.AllLinks[value] = true
		}
	}
}
