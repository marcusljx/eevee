package crossovers

import (
	"math/rand"
	"testing"
	"time"

	"fmt"

	"github.com/marcusljx/eevee/entities"
	"github.com/marcusljx/eevee/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestSimpleCrossover_Do(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a := entities.NewSimpleTextEntity("ABCDEF")
	b := entities.NewSimpleTextEntity("UVWXYZ")

	c := NewSimpleCrossover(1.0)
	c.Do(interfaces.EntityArray{a, b})
	assert.Equal(t, "ABCXYZ", fmt.Sprint(a))
	assert.Equal(t, "UVWDEF", fmt.Sprint(b))
}
