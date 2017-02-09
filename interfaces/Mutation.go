package interfaces

type Mutation interface {
	Do(entity SolutionEntity) Mutation
}
