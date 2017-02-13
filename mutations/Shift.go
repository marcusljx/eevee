package mutations

import (
	"errors"
	"math/rand"

	"github.com/marcusljx/eevee/interfaces"
)

var (
	InvalidNumOfTokensError = errors.New("Invalid number of tokens.")
)

type ShiftMutation struct {
	probability float64
	tokens      []rune
}

func NewShiftMutation(probability float64, tokens ...rune) *ShiftMutation {
	s := &ShiftMutation{}
	s.Probability(probability)
	s.AddTokens(tokens...)
	return s
}

func (s *ShiftMutation) Do(entity interfaces.SolutionEntity) error {
	if len(s.tokens) < 2 {
		return InvalidNumOfTokensError
	}

	rArr := entity.RuneArray()
	for i, _ := range rArr {
		if rand.Float64() < s.probability {
			rArr[i] = s.tokens[rand.Intn(len(s.tokens))] // assign it a new random value from s.tokens
		}
	}
	return nil
}

func (s *ShiftMutation) Probability(p float64) error {
	if p > 1 || p < 0 {
		return interfaces.InvalidProbabilityError
	}
	s.probability = p
	return nil
}

func (s *ShiftMutation) AddTokens(tokens ...rune) error {
	if len(tokens) < 2 {
		return InvalidNumOfTokensError
	}
	s.tokens = tokens
	return nil
}
