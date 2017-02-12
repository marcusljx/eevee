package interfaces

type Experiment interface {
	C(c Crossover) Experiment
	M(m Mutation) Experiment
	E(e SolutionEntity) Experiment
	Es(es []SolutionEntity) Experiment
	CycleLimit(n int) Experiment
	ScoreThreshold(s float64) Experiment

	Run() (solution SolutionEntity, finalPopulation []SolutionEntity)
}
