package logger

import (
	"fmt"
	"github.com/go-pkgz/lgr"
	"os"
)

// Logger structure for storing initialized logger
type Logger struct {
	logger *lgr.Logger
}

// Logf - function for printing to a file and to stdout
func (l *Logger) Logf(format string, args ...interface{}) {
	l.logger.Logf(format, args...)
	lgr.Printf(format, args...)
}

// Init - function for getting a handle and creating a folder with a file
func Init(folder, fileLog string) *Logger {
	err := os.MkdirAll(folder, 0755)
	if err != nil || os.IsExist(err) {
		panic("can't create log dir. no configured logging to files")
	}
	projectLog, err := os.OpenFile(fmt.Sprintf("%s/%s", folder, fileLog), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}

	l := lgr.New(lgr.Msec, lgr.Debug, lgr.CallerFile, lgr.CallerFunc, lgr.Format(lgr.FullDebug), lgr.Out(projectLog),
		lgr.Err(projectLog))

	return &Logger{
		logger: l,
	}
}
