package interfaces

import "fmt"

type Mutation interface {
	SingleChangeOperation
	Tokens(tokens []fmt.Stringer) error
}
