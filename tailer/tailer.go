package tailer

import (
	"glogagent/config"

	"github.com/hpcloud/tail"
)

func Init(c config.Collect) (*tail.Tail, error) {
	fileName := c.LogFilePath

	tailConf := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}
	return tail.TailFile(fileName, tailConf)
}
