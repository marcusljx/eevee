package crossovers

import (
	"log"

	"github.com/marcusljx/eevee/interfaces"
	"github.com/marcusljx/eevee/utils/errors"
)

// SimpleCrossover is a deterministic crossover operation that swaps two halves of two SolutionEntity objects.
// Using SimpleCrossover, two SolutionEntities representing ABCDEF and UVWXYZ will become ABCXYZ and UVWDEF respectively.
type SimpleCrossover struct {
	probability float64
}

// NewSimpleCrossover returns a new SimpleCrossover object
func NewSimpleCrossover(probability float64) *SimpleCrossover {
	return &SimpleCrossover{
		probability: probability,
	}
}

// Probability sets the probability at which this operation will occur
func (s *SimpleCrossover) Probability(p float64) interfaces.MultipleChangeOperation {
	if p > 1 || p < 0 {
		log.Print(interfaces.InvalidProbabilityError)
	}
	s.probability = p
	return s
}

// Do performs the Crossover operation on each iterative pair of an EntityArray.
// If the EntityArray has a non-even number of elements, the first element is re-used for the crossover with the last element.
func (s *SimpleCrossover) Do(entities interfaces.EntityArray) error {
	if len(entities)%2 != 0 {
		log.Print(errors.OddNumberOfEntitiesWarning)
	}
	for i := 0; i < len(entities); i += 2 {
		crossoverOperation(entities[i], entities[(i+1)%len(entities)]) // crossover each with next (last element is cyclic if container has odd number of elements)
	}
	return nil
}

// crossoverOperation performs the crossover method between
func crossoverOperation(entityA, entityB interfaces.SolutionEntity) error {
	arr1 := entityA.RuneArray()
	arr2 := entityB.RuneArray()
	if len(arr1) != len(arr2) {
		return errors.EntitySizesDoNotMatchError
	}

	midpt := len(arr1) / 2
	for i := midpt; i < len(arr1); i++ {
		arr1[i], arr2[i] = arr2[i], arr1[i]
	}

	entityA.Parse(arr1)
	entityB.Parse(arr2)

	return nil
}
