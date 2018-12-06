package job

import (
	"errors"
)

var (
	logicRun = make(map[string]logiclayer)
)

type logiclayer interface {
	Validate() (tasks map[int]interface{}, state bool)
	Run(ChanInputData) (interface{}, error)
	Done(*OutputData) bool
}

var Action action

type action struct {
	faset map[string]wrapperActionChannel
}
type wrapperActionChannel struct {
	input  chan interface{}
	output chan correlated
}

func (a action) Stop(nmRoutine string) (resp interface{}, err error) {
	mappingStatusTasks[nmRoutine] = stop
	tearDown(a.faset[nmRoutine].input, a.faset[nmRoutine].output)

	return nil, errors.New(status.String(stop))
}

// RegisterLogic - Register Logic Inside Scheduler
func registerLogic(
	nmRoutine string,
	logic logiclayer,
	input chan interface{},
	output chan correlated,
) (err error) {
	if _, ok := logicRun[nmRoutine]; ok {
		msg := "failed Registered Logic " + nmRoutine + "Already Registered"
		err = errors.New(msg)
		return
	}
	logicRun[nmRoutine] = logic

	actionMap := make(map[string]wrapperActionChannel)
	wraperChan := wrapperActionChannel{input: input, output: output}
	actionMap[nmRoutine] = wraperChan
	Action.faset = actionMap

	return
}
