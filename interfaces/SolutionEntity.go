package interfaces

import "fmt"

type SolutionEntity interface {
	fmt.Stringer

	Parse(string) SolutionEntity
	Copy(SolutionEntity) SolutionEntity
	Score() float64
	Validate() error
}