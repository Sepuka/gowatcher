package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

const (
	diskFreeCommand = "df"
	dfLoopInterval  = time.Hour * 6
)

func DiskFree(c chan<- command.Result) {
	cmd := command.NewCmd(diskFreeCommand, []string{"-hl", "--type=ext4", "--type=ext2", "--type=vfat"})
	resultHandler := command.NewDfFormatResultHandler(c)
	command.Runner(cmd, dfLoopInterval, resultHandler)
}
