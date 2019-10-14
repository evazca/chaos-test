package log

import (
	loggo "github.com/jeanphorn/log4go"
)

var CommonLogger  *loggo.Filter

func Log()  {
	loggo.LoadConfiguration("./log.json")
	CommonLogger = loggo.LOGGER("common")
}