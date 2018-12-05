package job

import "errors"

var (
	logicRun = make(map[string]logiclayer)
)

type logiclayer interface {
	Validate() (tasks map[int]interface{}, state bool)
	Run(ChanInputData) (interface{}, error)
	Done(*OutputData) bool
}

// RegisterLogic - Register Logic Inside Scheduler
func registerLogic(nmLogic string, logic logiclayer) (err error) {
	if _, ok := logicRun[nmLogic]; ok {
		msg := "failed Registered Logic " + nmLogic + "Already Registered"
		err = errors.New(msg)
		return
	}

	logicRun[nmLogic] = logic
	return
}
