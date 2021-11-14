package utils

import (
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat       = "[%lvl%]: %time% - %msg%"
	defaultTimestampFormat = time.RFC3339
)

type LogFormat struct {
	TimeStampFormat string `json:"time_stamp_format"`

	LogContent string `json:"log_content"`

	LogTruncate bool `json:"log_truncate"`
}

// Format
/**
 * @Description:
 * @receiver f
 * @param entry
 * @return []byte
 * @return error
 * example
 */
func (f *LogFormat) Format(entry *log.Entry) ([]byte, error) {
	output := f.LogContent
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimeStampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)

	_, s, _, o := runtime.Caller(8)
	if o {
		files := regexp.MustCompile(`plugin(.*?)/`).FindStringSubmatch(s)
		if len(files) < 1 {
			output = strings.Replace(output, "%file%", "leafBot", 1)
		} else {
			output = strings.Replace(output, "%file%", strings.TrimLeft(files[1], "_"), 1)
		}
	}

	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	if len(output) > 500 && f.LogTruncate {
		output = output[0:500] + "...\n"
	}

	return []byte(output), nil
}
