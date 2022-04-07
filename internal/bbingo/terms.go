package bbingo

import (
	_ "embed"
	"math/rand"
	"strings"
	"time"
)

//go:embed terms.txt
var allTerms string
var terms []string

func init() {
	for _, t := range strings.Split(allTerms, "\n") {
		rt := trysplit(strings.TrimSpace(t))
		terms = append(terms, rt)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(terms), func(i, j int) { terms[i], terms[j] = terms[j], terms[i] })
	allTerms = ""
}

// MkRealWord recreates the one-line string from a word that had newlines
func MkRealWord(word string) string {
	return strings.ReplaceAll(word, "\n", " ")
}

func trysplit(s string) string {
	if len(s) < 12 {
		return s
	}
	if !strings.Contains(s, " ") {
		return s
	}
	// start at about the middle and try to split close
	le, ri := len(s)/2, len(s)/2
	for {
		if s[le] != ' ' {
			le--
		} else {
			left := s[:le]
			right := s[le+1:]
			return left + "\n" + right
		}
		if s[ri] != ' ' {
			ri++
		} else {
			left := s[:ri]
			right := s[ri+1:]
			return left + "\n" + right
		}
		if le == 0 || ri == len(s) {
			return s
		}
	}
}
