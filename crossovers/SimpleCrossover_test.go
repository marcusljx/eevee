package crossovers

import (
	"testing"

	"github.com/marcusljx/eevee/entities"
)

func TestSimpleCrossover_Do(t *testing.T) {
	a := entities.NewSimpleTextEntity("HELLO")
	b := entities.NewSimpleTextEntity("WORLD")

	c := NewSimpleCrossover()

	c.Do(a, b)

	t.Log(a)
	t.Log(b)
}
