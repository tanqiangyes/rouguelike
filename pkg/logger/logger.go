package logger

import (
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/tanqiangyes/rouguelike/internal/config"
)

var (
	logger           = &logrus.Logger{}
	appLogsV1Entries = map[string]*logrus.Entry{}
	appLogsLock      sync.Mutex
)

func init() {
	logger.WithField("service", config.SERVICE).WithField("env", config.GetEnv())

	if v := os.Getenv("LOG_LEVEL"); v != "" {
		if lvl, err := logrus.ParseLevel(v); err == nil {
			logger.SetLevel(lvl)
		}
	}
}

// SetLevel 设置日志级别
func SetLevel(lvl logrus.Level) {
	logger.SetLevel(lvl)
}

// NewEntry 默认日志
func NewEntry() *logrus.Entry {
	return newAPPLogsV1Entry("")
}

// NewMainEntry 启动日志
func NewMainEntry() *logrus.Entry {
	return newAPPLogsV1Entry("main")
}

func newAPPLogsV1Entry(channel string) *logrus.Entry {
	appLogsLock.Lock()
	defer appLogsLock.Unlock()

	if e, ok := appLogsV1Entries[channel]; ok {
		return e
	}
	e := logrus.NewEntry(logger)
	if channel != "" {
		e = e.WithField("channel", channel)
	}
	if serverName := config.ServerName(); serverName != "" {
		e = e.WithField("serverName", serverName)
	}
	appLogsV1Entries[channel] = e
	return e
}

// PrintRuntime 打印执行时间日志对象
func PrintRuntime(channel string) func(message string, args ...interface{}) {
	st := time.Now()
	l := newAPPLogsV1Entry(channel).WithField("startTime", st)
	return func(message string, args ...interface{}) {
		et := time.Now()
		l = l.WithField("endTime", et).WithField("time", et.Sub(st).Milliseconds())
		if len(args) > 0 {
			l = l.WithField("args", args)
		}
		l.Debug(message)
	}
}
