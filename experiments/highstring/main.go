package main

import (
	"flag"

	"github.com/marcusljx/eevee/crossovers"
	"github.com/marcusljx/eevee/entities"
	"github.com/marcusljx/eevee/experiments"
	"github.com/marcusljx/eevee/interfaces"
	"github.com/marcusljx/eevee/mutations"
	"github.com/marcusljx/eevee/selectors"
	"github.com/marcusljx/eevee/utils/rand"
)

var (
	upperAlphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	entityLength  = flag.Int("entity-length", 10, "Specify the length of the solution entity.")
)

// HighStringExperiment demonstrates a simple experiment that aims to discover a string of the following properties:
//  The closer a character is to "Z", the higher its score is. "Z" is the "best" character.
//  A string (SolutionEntity) has no knowledge of itself.
//  Each SolutionEntity encounters operations that changes its structure.
//  Each generation's population is made up of the top SolutionEntities from the previous generation,
//  and the "new" 10 that are modified clones of the top half.
//  Of these, the top half of the total population will move on to the next generation
type HighStringExperiment struct {
	interfaces.BaseExperiment
}

func initialisePopulation(n int) (result []interfaces.SolutionEntity) {
	result = make([]interfaces.SolutionEntity, n)
	for i := 0; i < n; i++ {
		result[i] = entities.NewSimpleTextEntity(rand.String(*entityLength, upperAlphabet...))
	}
	return
}

func main() {
	x := new(HighStringExperiment)
	x.Crossover = crossovers.NewSimpleCrossover(0.5)
	x.Mutation = mutations.NewShiftMutation(0.2, upperAlphabet...)
	x.Selector = selectors.NewSimpleSelection(0.5)
	x.CycleLimit = 400
	x.ScoreThreshold = 900
	x.InitialPopulation = initialisePopulation(20)

	experiments.Run(x)
}
