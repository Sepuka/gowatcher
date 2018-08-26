package command

type ResultHandler interface {
	Handle(result Result)
}
