package solver_wrappers

import (
	"errors"
	"mservice/config"
	"mservice/models"
	"mservice/solvers"
)

func SelectWrapper(data [][]interface{}, taskName string) (models.ResultJson, error) {

	var solution models.ResultJson
	var err error

	switch taskName {
	case config.CycliclShift:
		solution, err = SolveForCyclicRotation(data)
	case config.Warrentries:
		solution, err = SolveForOthers(data, solvers.Warrentries, taskName)
	case config.AbscentElem:
		solution, err = SolveForOthers(data, solvers.AbscentElem, taskName)
	case config.SequenceCheck:
		solution, err = SolveForOthers(data, solvers.SequenceCheck, taskName)
	case "":
		return models.ResultJson{}, errors.New("no task name were specified")
	}

	if err != nil {
		return models.ResultJson{}, errors.New("failed during task solving")
	}

	return solution, nil
}
