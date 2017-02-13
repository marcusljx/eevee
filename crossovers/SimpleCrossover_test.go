package crossovers

import (
	"math/rand"
	"testing"
	"time"

	"github.com/marcusljx/eevee/entities"
)

func TestSimpleCrossover_Do(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a := entities.NewSimpleTextEntity("HELLO")
	b := entities.NewSimpleTextEntity("WORLD")

	c := NewSimpleCrossover()
	c.Do(a, b)

	t.Log(a)
	t.Log(b)
}
