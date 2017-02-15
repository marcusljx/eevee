package selectors

import (
	"sort"

	"github.com/marcusljx/eevee/interfaces"
)

// SimpleSelector is a basic selector.
// It selects the best r entities from a pool of N entities, where r is a ratio (float between 0 and 1)
type SimpleSelector struct {
	selectionRatio float64
}

// NewSimpleSelection returns a new SimpleSelector object
func NewSimpleSelection(ratio float64) *SimpleSelector {
	return &SimpleSelector{
		selectionRatio: ratio,
	}
}

// Do performs the selection operation
func (s *SimpleSelector) Do(entities ...interfaces.SolutionEntity) []interfaces.SolutionEntity {
	e := interfaces.EntityArray(entities)
	sort.Sort(e)

	trimmedPopulationSize := int(float64(len(e)) * s.selectionRatio)
	return []interfaces.SolutionEntity(e[len(e)-trimmedPopulationSize:])
}
