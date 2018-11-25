package command

import "github.com/sepuka/gowatcher/services"

// Do nothing ResultHandler, send result to chan only, without parse
type DummyResultHandler struct {
	c chan Result
}

func NewDummyResultHandler() ResultHandler {
	c := services.Container.Get(services.TransportChan).(chan Result)

	return DummyResultHandler{c}
}

func (handler DummyResultHandler) Handle(result Result) {
	handler.c <- result
}
