package rand

import "math/rand"

var (
	lowercase = []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numerals  = []rune("0123456789")

	alphabet    = append(lowercase, uppercase...)
	alphanumber = append(alphabet, numerals...)
)

func String(n int, tokens ...rune) string {
	if len(tokens) < 2 {
		tokens = alphabet
	}
	runes := make([]rune, n)
	for i, _ := range runes {
		runes[i] = tokens[rand.Intn(len(tokens))]
	}
	return string(runes)
}
