package interfaces

// Selector represents a selection operation
// This is most appropriately used at the end of a crossover/mutation iteration,
// to select the best from the combination of the previous generation's best and the newly changed generation.
type Selector interface {
	// Do performs the selection operation
	Do(entities ...SolutionEntity) []SolutionEntity
}
