package mutations

import (
	"math/rand"

	"github.com/marcusljx/eevee/interfaces"
)

type BarrelShift struct {
	probability float64
	rollValues  []int
}

func NewBarrelShiftMutation(probability float64, rollValues ...int) *BarrelShift {
	b := &BarrelShift{rollValues: rollValues}
	b.Probability(probability)
	return b
}

func (b *BarrelShift) Do(entity interfaces.SolutionEntity) error {
	if rand.Float64() < b.probability {
		// Pick rollValue at random
		roll := b.rollValues[rand.Intn(len(b.rollValues))]

		data := entity.RuneArray()
		cutPoint := len(data) - (roll % len(data))
		entity.Parse(append(data[cutPoint:], data[:cutPoint]...))
	}
	return nil
}

func (b *BarrelShift) Probability(p float64) error {
	if p > 1 || p < 0 {
		return interfaces.InvalidProbabilityError
	}
	b.probability = p
	return nil
}
