package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

type (
	Log struct {
		log    *logrus.Logger
		errLog *logrus.Logger
	}

	Params = map[string]interface{}
)

func (l *Log) Debug(msg string, params Params) { l.log.WithFields(params).Debug(msg) }
func (l *Log) Info(msg string, params Params)  { l.log.WithFields(params).Info(msg) }
func (l *Log) Warn(msg string, params Params)  { l.errLog.WithFields(params).Warn(msg) }
func (l *Log) Error(msg string, params Params) { l.errLog.WithFields(params).Error(msg) }

func New() *Log {
	l := &Log{
		log:    logrus.New(),
		errLog: logrus.New(),
	}

	for _, logger := range []*logrus.Logger{l.log, l.errLog} {
		logger.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: true,
		})
		logger.SetLevel(logrus.DebugLevel)
	}

	l.log.SetOutput(os.Stdout)
	l.errLog.SetOutput(os.Stderr)

	l.log.Info("logger initialized.")

	return l
}
