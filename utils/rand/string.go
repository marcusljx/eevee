package rand

import "math/rand"

var (
	lowercase = []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	alphabet = append(lowercase, uppercase...)
)

// String returns a random string given a set of tokens
func String(n int, tokens ...rune) string {
	if len(tokens) < 2 {
		tokens = alphabet
	}
	runes := make([]rune, n)
	for i := range runes {
		runes[i] = tokens[rand.Intn(len(tokens))]
	}
	return string(runes)
}
