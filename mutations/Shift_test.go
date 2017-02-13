package mutations

import (
	"math/rand"
	"testing"
	"time"

	"github.com/marcusljx/eevee/entities"
)

var (
	UpperAlphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func TestShiftMutation_Do(t *testing.T) {
	rand.Seed(time.Now().Unix())
	b := entities.NewSimpleTextEntity("WORLD")

	c := NewShiftMutation(0.5, UpperAlphabet...)
	c.Do(b)
	t.Log(b)
}
