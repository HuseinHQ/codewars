package sentencesmash

import "fmt"

func Smash(words []string) string {
	var result string
	for _, word := range words {
		result = fmt.Sprintf("%s %s", result, word)
	}

	return result
}