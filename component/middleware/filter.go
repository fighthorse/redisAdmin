package middleware

import "strings"

func FilterUrl(url string) bool {
	if strings.Contains(url, "/assets/") {
		return true
	}

	return false
}
