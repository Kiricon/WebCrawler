package web

import "strings"

// FormatDomain - Formate the domain so that it will work
func FormatUrl(domain string) string {
	if !strings.Contains(domain, "http://") {
		domain = "http://" + domain
	}

	return domain
}
