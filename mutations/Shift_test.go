package mutations

import (
	"math/rand"
	"testing"
	"time"

	"fmt"

	"github.com/marcusljx/eevee/entities"
	"github.com/stretchr/testify/assert"
)

var (
	UpperAlphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func TestShiftMutation_Do(t *testing.T) {
	rand.Seed(time.Now().Unix())
	b := entities.NewSimpleTextEntity(msg)

	c := NewShiftMutation(1.0, UpperAlphabet...)
	err := c.Do(b)
	assert.NoError(t, err)
	assert.NotEqual(t, msg, fmt.Sprint(b))
}
