package utils

import (
	log "github.com/sirupsen/logrus"
	"io"
	"runtime"
	"sync"
)

type LogHook struct {
	lock    *sync.Mutex
	levels  []log.Level
	format  log.Formatter
	writers []io.Writer
	LogChan chan string
}

func (l *LogHook) Levels() []log.Level {
	if len(l.levels) == 0 {
		return log.AllLevels
	} else {
		return l.levels
	}
}

func (l *LogHook) Fire(entry *log.Entry) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	data, err := l.format.Format(entry)

	if runtime.GOOS == "windows" {
		go func() {
			l.LogChan <- string(data)
		}()
	}

	if err != nil {
		return err
	}
	for _, writer := range l.writers {
		_, err := writer.Write(data)
		if err != nil {
			continue
		}
	}

	return err
}

func (l *LogHook) SetFormat(format log.Formatter) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.format == nil {
		// 用默认的
		l.format = &log.TextFormatter{DisableColors: true}
	} else {
		switch f := l.format.(type) {
		case *log.TextFormatter:
			textFormatter := f
			textFormatter.DisableColors = true
		default:
			// todo
		}
	}
	log.SetFormatter(format)
	l.format = format
}

func (l *LogHook) AddWriter(writer ...io.Writer) {
	l.writers = append(l.writers, writer...)
}

func (l *LogHook) AddLevel(level ...log.Level) {
	l.levels = append(l.levels, level...)
}

func NewLogHook(formatter log.Formatter, levels []log.Level, writers ...io.Writer) *LogHook {
	hook := &LogHook{
		lock:    new(sync.Mutex),
		LogChan: make(chan string, 10),
	}
	hook.AddWriter(writers...)
	hook.AddLevel(levels...)
	hook.SetFormat(formatter)
	return hook
}

// GetLogLevel 获取日志等级
//
// 可能的值有
//
// "trace","debug","info","warn","warn","error"
func GetLogLevel(level string) []log.Level {
	switch level {
	case "trace":
		return []log.Level{
			log.TraceLevel, log.DebugLevel,
			log.InfoLevel, log.WarnLevel, log.ErrorLevel,
			log.FatalLevel, log.PanicLevel,
		}
	case "debug":
		return []log.Level{
			log.DebugLevel, log.InfoLevel,
			log.WarnLevel, log.ErrorLevel,
			log.FatalLevel, log.PanicLevel,
		}
	case "info":
		return []log.Level{
			log.InfoLevel, log.WarnLevel,
			log.ErrorLevel, log.FatalLevel, log.PanicLevel,
		}
	case "warn":
		return []log.Level{
			log.WarnLevel, log.ErrorLevel,
			log.FatalLevel, log.PanicLevel,
		}
	case "error":
		return []log.Level{
			log.ErrorLevel, log.FatalLevel,
			log.PanicLevel,
		}
	default:
		return []log.Level{
			log.InfoLevel, log.WarnLevel,
			log.ErrorLevel, log.FatalLevel, log.PanicLevel,
		}
	}
}
