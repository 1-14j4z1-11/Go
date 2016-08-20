package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := makeRune(rng)
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNotPalindrome(rng *rand.Rand) string {
	runes := []rune(randomPalindrome(rng))
	length := len(runes)

	if length < 1 {
		return string(runes)
	}

	idx := 0
	if length > 3 {
		rng.Intn(length / 2 - 1)
	}
	replace := runes[idx]
	for unicode.ToLower(replace) == unicode.ToLower(runes[idx]) {
		replace = makeRune(rng)
	}
	runes[idx] = replace

	return string(runes)
}

func insertRune(word string, insert rune, rng *rand.Rand) string {
	runes := []rune(word)
	length := len(runes)
	idx := 0

	if length > 2 {
		rng.Intn(length)
	}

	var newRunes []rune
	var offset int
	for i := 0; i < length; i++ {
		if i + offset == idx {
			newRunes = append(newRunes, insert)
			offset++;
			i--;
		} else {
			newRunes = append(newRunes, runes[i])
		}
	}

	return string(newRunes)
}

func makeRune(rng *rand.Rand) rune {
	var r rune

	for !unicode.IsLetter(r) {
		r = rune(rng.Intn(0xFF))
	}

	return r
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		p = insertRune(p, ',', rng)
		p = insertRune(p, ' ', rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNotPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)
		if len([]rune(p)) < 2 {
			continue
		}

		p = insertRune(p, ',', rng)
		p = insertRune(p, ' ', rng)

		if IsPalindrome(p) {
			t.Errorf("[%4d] !IsPalindrome(%v) = true", i, []rune(p))
		}
	}
}
