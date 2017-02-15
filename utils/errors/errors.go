package errors

import "errors"

var (
	// EntitySizesDoNotMatchError is the error thrown when two SolutionEntity objects in the same Crossover operation do not have matching sizes
	EntitySizesDoNotMatchError = errors.New("Entity sizes do not match.")

	// OddNumberOfEntitiesWarning is a Crossover-level warning that is thrown when a pair-wise Crossover operation is attempted with an EntityArray of uneven size.
	OddNumberOfEntitiesWarning = errors.New("Expected even number of entities but EntityArray was of odd-number-size. The last element will crossover with the 1st element (ie. 1st element will be used twice)")

	// InvalidNumOfTokensError is the error for checking if a correct number of tokens are passed.
	InvalidNumOfTokensError = errors.New("Invalid number of tokens.")
)
