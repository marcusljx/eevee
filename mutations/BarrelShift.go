package mutations

import (
	"math/rand"

	"log"

	"github.com/marcusljx/eevee/interfaces"
)

// BarrelShift is a probabilistically-deterministic mutation that barrel-shifts a SolutionEntity's runes by a random amount.
// With BarrelShift, the SolutionEntity representing "ABCDE" will become one of any of these limited possibilities:
//  BCDEA
//  CDEAB
//  DEABC
//  EABCD
//  ABCDE
//
// BarrelShift implements the Mutation interface
type BarrelShift struct {
	probability float64
	rollValues  []int
}

// NewBarrelShiftMutation returns a new BarrelShift object
func NewBarrelShiftMutation(probability float64, rollValues ...int) *BarrelShift {
	b := &BarrelShift{rollValues: rollValues}
	b.Probability(probability)
	return b
}

// Do performs the operation
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

// Probability specifies the probability that the operation will be performed
func (b *BarrelShift) Probability(probability float64) interfaces.SingleChangeOperation {
	if probability > 1 || probability < 0 {
		log.Print(interfaces.InvalidProbabilityError)
	}
	b.probability = probability
	return b
}
