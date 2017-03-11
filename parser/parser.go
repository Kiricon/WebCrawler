package parser

import (
	"io"

	"fmt"

	"golang.org/x/net/html"
)

// GetLinks - Get all links on a web page.
func GetLinks(body io.Reader) {

	page := html.NewTokenizer(body)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
				}
			}
		}
	}

}
