package experiments

import (
	"log"

	"github.com/marcusljx/eevee/interfaces"
)

const (
	showPopulationScoresTableFormat = "|%3d| %-20s |%6.2f|\n"
)

var (
	logFunc = log.Printf
)

func showScores(cycle uint64, population []interfaces.SolutionEntity) {
	logFunc("\n[cycle %9d]\n", cycle)
	for i, e := range population {
		logFunc(showPopulationScoresTableFormat, i, e, e.Score())
	}
}

// Run performs the input experiment
func Run(experiment interfaces.Experiment) error {
	pop := experiment.GetInitialPopulation()
	var cycle uint64 = 0

	showScores(cycle, pop)

	for {
		// check cycle
		if cycle >= experiment.GetCycleLimit() {
			logFunc("Cycle limit [%d] reached. Stopping.", cycle)
			break
		}
		cycle++

		// check threshold
		if pop[len(pop)-1].Score() >= experiment.GetScoreThreshold() {
			logFunc("Score threshold [%6.3f] reached. Stopping.", experiment.GetScoreThreshold())
			break
		}

		// save original population
		clones := pop.Clone()

		// do mutation
		for _, e := range clones {
			err := experiment.GetMutation().Do(e)
			if err != nil {
				logFunc("ERROR while performing mutation: %v\n", err)
				return err
			}
		}

		// do crossover
		err := experiment.GetCrossover().Do(clones)
		if err != nil {
			logFunc("ERROR while performing crossover: %v\n", err)
			return err
		}

		// combine and select top
		combined := append(pop, clones...)
		pop = experiment.GetSelector().Do(combined...)

		// Show Population
		showScores(cycle, pop)
	}

	// Finish Experiment
	return nil
}
