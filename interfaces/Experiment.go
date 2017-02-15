package interfaces

// Experiment is the interface for an experiment
type Experiment interface {
	// GetInitialPopulation returns the InitialPopulation of the experiment
	GetInitialPopulation() EntityArray

	// GetCrossover returns the Crossover of the experiment
	GetCrossover() Crossover

	// GetMutation returns the Mutation of the experiment
	GetMutation() Mutation

	// GetSelector returns the Selector of the experiment
	GetSelector() Selector

	// GetCycleLimit returns the CycleLimit of the experiment
	GetCycleLimit() uint64

	// GetScoreThreshold returns the ScoreThreshold of the experiment
	GetScoreThreshold() float64
}

// BaseExperiment is a standard type for use with creating experiments.
// All experiments using eevee must embed BaseExperiment
type BaseExperiment struct {
	InitialPopulation EntityArray
	Crossover         Crossover
	Mutation          Mutation
	Selector          Selector
	CycleLimit        uint64
	ScoreThreshold    float64
}

// GetInitialPopulation returns the InitialPopulation of the experiment
func (b *BaseExperiment) GetInitialPopulation() EntityArray {
	return b.InitialPopulation
}

// GetCrossover returns the Crossover of the experiment
func (b *BaseExperiment) GetCrossover() Crossover {
	return b.Crossover
}

// GetMutation returns the Mutation of the experiment
func (b *BaseExperiment) GetMutation() Mutation {
	return b.Mutation
}

// GetSelector returns the Selector of the experiment
func (b *BaseExperiment) GetSelector() Selector {
	return b.Selector
}

// GetCycleLimit returns the CycleLimit of the experiment
func (b *BaseExperiment) GetCycleLimit() uint64 {
	return b.CycleLimit
}

// GetScoreThreshold returns the ScoreThreshold of the experiment
func (b *BaseExperiment) GetScoreThreshold() float64 {
	return b.ScoreThreshold
}
