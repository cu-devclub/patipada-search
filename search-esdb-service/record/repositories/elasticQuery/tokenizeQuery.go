package elasticQuery

import "fmt"

func BuildAnalyzeQuery(index, query string) string {
	return fmt.Sprintf(`{
        "tokenizer": "icu_tokenizer",
        "text": "%s"
    }`, query)
}
