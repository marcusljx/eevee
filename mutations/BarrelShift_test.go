package mutations

import (
	"math/rand"
	"testing"
	"time"

	"fmt"

	"github.com/marcusljx/eevee/entities"
	"github.com/stretchr/testify/assert"
)

const (
	msg = "HELLOWORLD"
)

func TestBarrelShiftMutation_Do(t *testing.T) {
	rand.Seed(time.Now().Unix())
	b := entities.NewSimpleTextEntity(msg)

	c := NewBarrelShiftMutation(1.0, 1, 3, 5, 7)
	err := c.Do(b)
	assert.NoError(t, err)
	assert.NotEqual(t, msg, fmt.Sprint(b))
}
