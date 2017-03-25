package parser

import (
	"fmt"
	"os"
	"strings"
)

// IsLocal - Check if a link is foreign or local
func IsLocal(link string, domain string) bool {
	if link[0] == '/' {
		return true
	}

	if strings.Contains(link, "://"+domain) {
		return true
	}

	return false
}

// AppendToFile - Append a url to the file
func AppendToFile(url string) {
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	n, nerr := f.WriteString(url + "\n")
	if nerr != nil {
		fmt.Println(nerr)
		fmt.Println(n)
	}

	f.Close()
}

// SetupFile - Setup our data output file if it's missing
func SetupFile() {
	if _, err := os.Stat("data.txt"); os.IsNotExist(err) {

		f, ferr := os.Create("data.txt")
		if ferr != nil {
			fmt.Println(ferr)
		}
		fmt.Println("Data File Created!")
		f.Close()

	}
}
