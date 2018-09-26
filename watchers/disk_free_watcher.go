package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"time"
)

const (
	diskFreeCommand = "df"
	DfLoopInterval  = time.Hour * 6
)

func DiskFree(c chan<- command.Result, config config.WatcherConfig) {
	cmd := command.NewCmd(diskFreeCommand, []string{"-hl", "--type=ext4", "--type=ext2", "--type=vfat"})
	resultHandler := command.NewDfFormatResultHandler(c)
	command.RunCmdLoop(cmd, config.GetLoop(), resultHandler)
}
