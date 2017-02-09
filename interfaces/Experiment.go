package interfaces

type Experiment interface {
	C(c Crossover) Experiment
	M(m Mutation) Experiment
	E(e SolutionEntity) Experiment
	Es([]SolutionEntity) Experiment
	CycleLimit(n int) Experiment
	ScoreThreshold(score float64) Experiment
}
