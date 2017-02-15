package interfaces

// Mutation represents a mutation operation
type Mutation interface {
	SingleChangeOperation

	// Tokens specifies the runes that make up the set of possible characters(runes) for mutation. For example, the tokens of a DNA sequence is specified as []rune("AGTC")
	Tokens(tokens []rune) Mutation
}
