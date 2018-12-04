package job

var (
	logicRun = make(map[string]logiclayer)
)

// ChanInputData - Channel Receiver Data
type ChanInputData struct {
	State chan int
	Data  interface{}
}
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
