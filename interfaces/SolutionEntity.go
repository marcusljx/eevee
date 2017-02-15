package interfaces

import "fmt"

// SolutionEntity represents a single solution. It also wraps the fmt.Stringer interface.
type SolutionEntity interface {
	fmt.Stringer
	Score() float64
	RuneArray() []rune
	Parse(r []rune)
	Clone() SolutionEntity
}

// EntityArray is a type alias for []SolutionEntity
type EntityArray []SolutionEntity

// Len is an interface method for sort.Sort(EntityArray)
func (e EntityArray) Len() int {
	return len(e)
}

// Less is an interface method for sort.Sort(EntityArray)
func (e EntityArray) Less(i, j int) bool {
	return e[i].Score() < e[j].Score()
}

// Swap is an interface method for sort.Sort(EntityArray)
func (e EntityArray) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Clone returns a clone of the EntityArray object. It is a new object with its own memory space.
func (e EntityArray) Clone() EntityArray {
	result := make(EntityArray, len(e))
	for i, entity := range e {
		result[i] = entity.Clone()
	}

	return result
}
