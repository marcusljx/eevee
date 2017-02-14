package mutations

import (
	"math/rand"
	"testing"
	"time"

	"github.com/marcusljx/eevee/entities"
)

func TestBarrelShiftMutation_Do(t *testing.T) {
	rand.Seed(time.Now().Unix())
	b := entities.NewSimpleTextEntity("HELLOWORLD")

	c := NewBarrelShiftMutation(1.0, 1, 3, 5, 7)
	c.Do(b)
	t.Log(b)
}
