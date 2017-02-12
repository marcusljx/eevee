package experiments

import (
	"github.com/marcusljx/eevee/interfaces"
)

type HighStringExperiment struct {
	alphabet  interfaces.Alphabet
	crossover interfaces.Crossover
	mutation  interfaces.Mutation
	create    func(alphabet interfaces.Alphabet) interfaces.SolutionEntity

	workingPopulation []interfaces.SolutionEntity
	cycleLimit        int
	scoreThreshold    float64
}

func (h *HighStringExperiment) GenesisFunc(f func(alphabet interfaces.Alphabet) interfaces.SolutionEntity) interfaces.Experiment {
	h.create = f
	return h
}

func (h *HighStringExperiment) C(c interfaces.Crossover) interfaces.Experiment {
	h.crossover = c
	return h
}

func (h *HighStringExperiment) M(m interfaces.Mutation) interfaces.Experiment {
	h.mutation = m
	return h
}

func (h *HighStringExperiment) E(e interfaces.SolutionEntity) interfaces.Experiment {
	h.workingPopulation = append(h.workingPopulation, e)
	return h
}

func (h *HighStringExperiment) Es(es []interfaces.SolutionEntity) interfaces.Experiment {
	h.workingPopulation = append(h.workingPopulation, es...)
	return h
}

func (h *HighStringExperiment) CycleLimit(n int) interfaces.Experiment {
	h.cycleLimit = n
	return h
}

func (h *HighStringExperiment) ScoreThreshold(s float64) interfaces.Experiment {
	h.scoreThreshold = s
	return h
}

func (h *HighStringExperiment) Run() (solution interfaces.SolutionEntity, finalPopulation []interfaces.SolutionEntity) {

}