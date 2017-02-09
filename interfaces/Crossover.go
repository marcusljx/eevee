package interfaces

type Crossover interface {
	AddExtractor(extractFunc func(entity SolutionEntity) interface{}) Crossover
	Do([]SolutionEntity) Crossover
}
