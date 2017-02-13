package interfaces

import "fmt"

type SolutionEntity interface {
	fmt.Stringer
	Score() float64
	RuneArray() []rune
	Parse(r []rune)
}
