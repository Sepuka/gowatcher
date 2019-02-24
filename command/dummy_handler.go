package command

// Do nothing ResultHandler, send result to chan only, without parse
type dummyResultHandler struct {
	c chan<-Result
}

func NewDummyResultHandler(transportChan chan<-Result) ResultHandler {
	return &dummyResultHandler{c: transportChan}
}

func (handler dummyResultHandler) Handle(result Result) {
	handler.c <- result
}
