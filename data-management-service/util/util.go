package util

import "github.com/microcosm-cc/bluemonday"

func Contains(s string, arr []string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func ExtractRawStringFromHTMLTags(s string) string {
	p := bluemonday.StrictPolicy()
	return p.Sanitize(s)
}
