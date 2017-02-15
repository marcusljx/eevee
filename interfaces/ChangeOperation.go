package interfaces

import "errors"

var (
	// InvalidProbabilityError is the standard error for reporting that a probability error is invalid
	InvalidProbabilityError = errors.New("Invalid probability value. Expected range is [0.0, 1.0]")
)

// SingleChangeOperation is the interface type for an operation that modifies a single SolutionEntity
type SingleChangeOperation interface {
	// Do performs the operation
	Do(entity SolutionEntity) error

	// Probability sets the probability at which the operation will happen
	Probability(probability float64) SingleChangeOperation
}

// MultipleChangeOperation is the interface type for an operation that modifies an EntityArray(group of SolutionEntity objects)
type MultipleChangeOperation interface {
	// Do performs the operation
	Do(entities EntityArray) error

	// Probability sets the probability at which the operation will happen
	Probability(probability float64) MultipleChangeOperation
}
