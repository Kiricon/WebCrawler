package crawler

import (
	"WebCrawler/parser"
	"WebCrawler/stats"
	"WebCrawler/web"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// Crawler - A crawler object for each web page.
type Crawler struct {
	Domain   string
	Path     string
	AllLinks map[string]bool
	Wg       *sync.WaitGroup
	Stats    *stats.Stats
}

// Crawl - Recursively crawls through unexplored links on pages
func (c *Crawler) Crawl() {
	fmt.Println(c.Stats.TotalLinks)
	c.GetPage()
	c.Wg.Done()
}

// GetPage - Function to get a web page as a string
func (c *Crawler) GetPage() {

	url := web.FormatUrl(c.Domain + c.Path)

	resp, respErr := http.Get(url)
	if respErr != nil {
		fmt.Println(respErr)
	}
	defer resp.Body.Close()
	c.getLinks(resp.Body)
}

//MakeNewCrawler - Make a new webcrawler for the sub page
func makeNewCrawler(domain string, path string, allLinks map[string]bool, wg *sync.WaitGroup, stats *stats.Stats) Crawler {

	if path[0] != '/' {
		//path = strings.Replace(path, domain, "", 1)
		path = strings.Split(path, domain)[1]
	}
	return Crawler{Domain: domain, Path: path, AllLinks: allLinks, Wg: wg, Stats: stats}
}

func (c *Crawler) getLinks(body io.Reader) {

	page := html.NewTokenizer(body)
	c.Stats.TotalRoutines--
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					c.handelLink(attr.Val)
				}
			}
		}
	}

}

func (c *Crawler) handelLink(linkURL string) {
	if linkURL != "" {
		if !c.AllLinks[linkURL] && parser.IsLocal(linkURL, c.Domain) {
			c.AllLinks[linkURL] = true
			parser.AppendToFile(linkURL)
			c.Stats.TotalLinks++
			newCrawler := makeNewCrawler(c.Domain, linkURL, c.AllLinks, c.Wg, c.Stats)
			c.Wg.Add(1)
			c.Stats.TotalRoutines++
			newCrawler.Crawl()
		}
	}
}
