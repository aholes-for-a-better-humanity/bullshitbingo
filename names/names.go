package names

import (
	"math/rand"
	"strings"
)

func coinFlip() bool { return rand.Float64() < .5 }

// syllab makes a single syllab
func syllab() string {
	if coinFlip() {
		return consonnant() + vowelSound()
	}
	return vowelSound() + consonnant()
}

var consonnants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t"}
var nconsonnants = len(consonnants)

func consonnant() string { return consonnants[rand.Intn(nconsonnants)] }

var vowels = []string{"a", "e", "i", "o", "u", "y"}
var nvowels = len(vowels)

func rv() string { return vowels[rand.Intn(nvowels)] }
func vowelSound() string {
	full := rv()
	last := full
	for v := rv(); ; v = rv() {
		if coinFlip() || len(full) > 2 {
			return full
		}
		if v == last && (v == "y" || v == "u") {
			continue
		} // do not double these
		last = v
		full += v
	}
}

func Word() string {
	word := syllab()
	for i := 0; i < 5 && coinFlip() && len(word) < 8; i++ {
		word += syllab()
	}
	return word
}

func Pseudo() string {
	wds := make([]string, 2)
	for i := 0; i < 2; i++ {
		chars := []rune(Word())
		wds[i] = strings.ToUpper(string(chars[0])) + string(chars[1:])
	}
	return strings.Join(wds, " ")
}
