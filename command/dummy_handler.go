package command

// Do nothing ResultHandler, send result to chan only, without parse
type DummyResultHandler struct {
	c chan<- Result
}

func NewDummyResultHandler(c chan<- Result) ResultHandler {
	return DummyResultHandler{c}
}

func (handler DummyResultHandler) Handle(result Result) {
	handler.c <- result
}
