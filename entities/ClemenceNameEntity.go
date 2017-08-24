package entities

import "github.com/marcusljx/eevee/interfaces"

// ClemenceNameEntity represents a string of data.
// It implements the SolutionEntity interface.
type ClemenceNameEntity struct {
	data []rune
}

// NewClemenceNameEntity returns a new ClemenceNameEntity object
func NewClemenceNameEntity(s string) *ClemenceNameEntity {
	return &ClemenceNameEntity{
		data: []rune(s),
	}
}

// RuneArray returns the internal []rune representation of the data.
func (s *ClemenceNameEntity) RuneArray() []rune {
	return s.data
}

// Parse sets the internal []rune representation of the data
func (s *ClemenceNameEntity) Parse(r []rune) {
	s.data = r
}

// String satisfies the fmt.Stringer interface
func (s *ClemenceNameEntity) String() string {
	return string(s.data)
}

// Score returns the score of the the SolutionEntity
func (s *ClemenceNameEntity) Score() (sum float64) {
	if s.data[0] == 'C' {
		sum++
	}
	if s.data[1] == 'L' {
		sum++
	}
	if s.data[2] == 'E' {
		sum++
	}
	if s.data[3] == 'M' {
		sum++
	}
	if s.data[4] == 'E' {
		sum++
	}
	if s.data[5] == 'N' {
		sum++
	}
	if s.data[6] == 'C' {
		sum++
	}
	if s.data[7] == 'E' {
		sum++
	}
	return
}

// Clone returns a new SolutionEntity object with the same internal data structure and values
func (s *ClemenceNameEntity) Clone() interfaces.SolutionEntity {
	result := &ClemenceNameEntity{
		data: make([]rune, len(s.data)),
	}
	copy(result.data, s.data)
	return result
}
