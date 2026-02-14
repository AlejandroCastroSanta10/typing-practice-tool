package wordlist

import (
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed words.txt
var wordsFile string

var words []string

func init() {
	for _, w := range strings.Split(wordsFile, "\n") {
		w = strings.TrimSpace(w)
		if w != "" {
			words = append(words, w)
		}
	}
}

// Generate returns a string of random words from the embedded list,
// with at least `minChars` total characters.
func Generate(minChars int) string {
	var result []string
	total := 0
	for total < minChars {
		w := words[rand.Intn(len(words))]
		result = append(result, w)
		total += len(w) + 1
	}
	return strings.Join(result, " ")
}
