package main

import (
	"log"
	"os"

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
	entityLength  = 8
)

type ClemenceNameExperiment struct {
	interfaces.BaseExperiment
}

func initialisePopulation(n int) (result []interfaces.SolutionEntity) {
	result = make([]interfaces.SolutionEntity, n)
	for i := 0; i < n; i++ {
		result[i] = entities.NewClemenceNameEntity(rand.String(entityLength, upperAlphabet...))
	}
	return
}

func main() {
	x := new(ClemenceNameExperiment)
	x.Crossover = crossovers.NewSimpleCrossover(0.5)
	x.Mutation = mutations.NewShiftMutation(0.2, upperAlphabet...)
	x.Selector = selectors.NewSimpleSelection(0.5)
	x.CycleLimit = 1000
	x.ScoreThreshold = 8
	x.InitialPopulation = initialisePopulation(20)

	err := experiments.Run(x)
	if err != nil {
		log.Printf("experiment returned error: %v", err)
		os.Exit(1)
	}
}
