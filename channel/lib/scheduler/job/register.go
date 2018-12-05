package job

var (
	logicRun = make(map[string]logiclayer)
)

type logiclayer interface {
	Run(ChanInputData)
}

// RegisterLogic - Register Logic Inside Scheduler
func RegisterLogic(nmLogic string, logic logiclayer) {
	if _, ok := logicRun[nmLogic]; ok {
		msg := "failed Registered Logic " + nmLogic + "Already Registered"
		panic(msg)
		return
	}

	logicRun[nmLogic] = logic
}