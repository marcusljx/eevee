package interfaces

import "fmt"

type SolutionEntity interface {
	fmt.Stringer
	Score() float64
	AsRuneArray() []rune
	ParseRuneArray(rArr []rune)
}
