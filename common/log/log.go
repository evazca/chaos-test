package log

import (
	loggo "github.com/jeanphorn/log4go"
)

var commonLogger  *loggo.Filter

func Log()  {
	loggo.LoadConfiguration("./log.json")
	commonLogger = loggo.LOGGER("common")
}

func CommonLogger() *loggo.Filter{
	if CommonLogger == nil{
		Log()
	}
	return commonLogger
}