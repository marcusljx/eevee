package entities

import "github.com/marcusljx/eevee/interfaces"

// SimpleTextEntity represents a string of data.
// It implements the SolutionEntity interface.
type SimpleTextEntity struct {
	data []rune
}

// NewSimpleTextEntity returns a new SimpleTextEntity object
func NewSimpleTextEntity(s string) *SimpleTextEntity {
	return &SimpleTextEntity{
		data: []rune(s),
	}
}

// RuneArray returns the internal []rune representation of the data.
func (s *SimpleTextEntity) RuneArray() []rune {
	return s.data
}

// Parse sets the internal []rune representation of the data
func (s *SimpleTextEntity) Parse(r []rune) {
	s.data = r
}

// String satisfies the fmt.Stringer interface
func (s *SimpleTextEntity) String() string {
	return string(s.data)
}

// Score returns the score of the the SolutionEntity
func (s *SimpleTextEntity) Score() (sum float64) {
	for _, r := range s.data {
		sum += float64(r)
	}
	return
}

// Clone returns a new SolutionEntity object with the same internal data structure and values
func (s *SimpleTextEntity) Clone() interfaces.SolutionEntity {
	result := &SimpleTextEntity{
		data: make([]rune, len(s.data)),
	}
	copy(result.data, s.data)
	return result
}
