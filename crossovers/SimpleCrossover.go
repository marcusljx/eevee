package crossovers

import (
	"errors"

	"github.com/marcusljx/eevee/interfaces"
)

var (
	EntitySizesDoNotMatchError = errors.New("Entity sizes do not match.")
)

type SimpleCrossover struct {
	probability float64
}

func NewSimpleCrossover() *SimpleCrossover {
	return &SimpleCrossover{
		probability: 0.5,
	}
}

func (s *SimpleCrossover) Probability(p float64) error {
	if p > 1 || p < 0 {
		return interfaces.InvalidProbabilityError
	}
	s.probability = p
	return nil
}

func (s *SimpleCrossover) Do(entityA interfaces.SolutionEntity, entityB interfaces.SolutionEntity) error {
	a := entityA.RuneArray()
	b := entityB.RuneArray()
	if len(a) != len(b) {
		return EntitySizesDoNotMatchError
	}

	midpt := len(a) / 2
	for i := midpt; i < len(a); i++ {
		a[i], b[i] = b[i], a[i]
	}

	entityA.ParseRuneArray(a)
	entityB.ParseRuneArray(b)

	return nil
}
