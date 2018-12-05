package job

var (
	logicRun = make(map[string]logiclayer)
)

type logiclayer interface {
	Validate() bool
	Run(ChanInputData) (interface{}, error)
	Done(*OutputData) bool
}

// RegisterLogic - Register Logic Inside Scheduler
func registerLogic(nmLogic string, logic logiclayer) {
	if _, ok := logicRun[nmLogic]; ok {
		msg := "failed Registered Logic " + nmLogic + "Already Registered"
		panic(msg)
		return
	}

	logicRun[nmLogic] = logic
}
