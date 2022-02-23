package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func Init(folder, fileLog string) {
	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll(folder, 0755)
	if err != nil || os.IsExist(err) {
		panic("can't create log dir. no configured logging to files")
	}

	projectLog, err := os.OpenFile(fmt.Sprintf("%s/%s", folder, fileLog), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}

	logrus.SetOutput(io.Discard) // из-за необходимости записи в 2 места одновременно

	log.AddHook(&writerHook{
		Writer:    []io.Writer{projectLog, os.Stdout},
		LogLevels: logrus.AllLevels,
	})
	log.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(log)
}
