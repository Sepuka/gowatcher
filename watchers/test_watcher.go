package watchers

import "github.com/sepuka/gowatcher/command"

const msg = "It's work"

func Test() command.Result {
	return command.NewResult("test", msg,nil)
}
