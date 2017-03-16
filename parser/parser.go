package parser

import (
	"io"
	"os"

	"strings"

	"fmt"

	"golang.org/x/net/html"
)

// GetLinks - Get all links on a web page.
func GetLinks(body io.Reader, allLinks map[string]bool, domain string) []string {
	links := []string{}
	page := html.NewTokenizer(body)

	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
					if !allLinks[attr.Val] && isLocal(attr.Val, domain) {
						allLinks[attr.Val] = true
						f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0666)
						if err != nil {
							fmt.Println(err)
						}
						n, nerr := f.WriteString(attr.Val + "\n")
						if nerr != nil {
							fmt.Println(nerr)
							fmt.Println(n)
						}

						f.Close()
					}
				}
			}
		}
	}

}

func isLocal(link string, domain string) bool {
	if link[0] == '/' || strings.Contains(link, domain) {
		return true
	}

	return false
}
