package mutations

import "github.com/marcusljx/eevee/interfaces"

type BarrelShift struct {
	probabilty float64
	roll       int
}

func NewBarrelShiftMutation(rollValue int) *BarrelShift {
	return &BarrelShift{
		roll: rollValue,
	}
}

func (b *BarrelShift) Do(entity interfaces.SolutionEntity) error {
	data := entity.Data()
	splitPoint := (len(data) - b.roll) % len(data)

	front, back := data[:splitPoint], data[splitPoint:]
	entity.Parse(append(back, front...))
}

func (b *BarrelShift) Probability(p float64) {
	b.probabilty = p
}
