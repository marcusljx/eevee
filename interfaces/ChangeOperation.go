package interfaces

import "errors"

var (
	InvalidProbabilityError = errors.New("Invalid probability value. Expected range is [0.0, 1.0]")
)

type ChangeOperation interface {
	Probability(probability float64) error
}

type SingleChangeOperation interface {
	Do(entity SolutionEntity) error
	ChangeOperation
}

type DupleChangeOperation interface {
	Do(entityA SolutionEntity, entityB SolutionEntity) error
	ChangeOperation
}

type MultipleChangeOperation interface {
	Do(entities []SolutionEntity) error
	ChangeOperation
}
