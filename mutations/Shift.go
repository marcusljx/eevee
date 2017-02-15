package mutations

import (
	"log"
	"math/rand"

	"github.com/marcusljx/eevee/interfaces"
	"github.com/marcusljx/eevee/utils/errors"
)

//)

// ShiftMutation is a non-deterministic index-shift mutation by a randomly chosen number
// Effectively it is simple a mutation where a rune is substituted for another rune, chosen at random.
// ShiftMutation implements the interfaces.Mutation interface
type ShiftMutation struct {
	probability float64
	tokens      []rune
}

// NewShiftMutation returns a new ShiftMutation object
func NewShiftMutation(probability float64, tokens ...rune) *ShiftMutation {
	s := &ShiftMutation{
		probability: probability,
		tokens:      tokens,
	}
	return s
}

// Probability specifies the probability where the mutation will happen
func (s *ShiftMutation) Probability(probability float64) interfaces.SingleChangeOperation {
	if probability > 1 || probability < 0 {
		log.Print(interfaces.InvalidProbabilityError)
	}
	s.probability = probability
	return s
}

// Tokens specify the set of runes that are used for the mutation
func (s *ShiftMutation) Tokens(tokens []rune) interfaces.Mutation {
	if len(tokens) < 2 {
		log.Print(errors.InvalidNumOfTokensError)
	}
	s.tokens = tokens
	return s
}

// Do performs the mutation operation.
func (s *ShiftMutation) Do(entity interfaces.SolutionEntity) error {
	if len(s.tokens) < 2 {
		return errors.InvalidNumOfTokensError
	}

	rArr := entity.RuneArray()
	for i, _ := range rArr {
		if rand.Float64() < s.probability {
			rArr[i] = s.tokens[rand.Intn(len(s.tokens))] // assign it a new random value from s.tokens
		}
	}
	entity.Parse(rArr)
	return nil
}
