package logger

import (
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Log *logrus.Logger

func InitLogger() {
	Log = newLogger()

	// Set log level from env
	logLevelStr := viper.GetString("LOG_LEVEL")
	level, err := logrus.ParseLevel(strings.ToLower(logLevelStr))
	if err != nil {
		Log.Warnf("Invalid LOG_LEVEL: %s, fallback to info", logLevelStr)
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)
}

func newLogger() *logrus.Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)

	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339, // ISO8601 format, great for logs
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp", // Kibana/ES prefers this
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "caller",
		},
		PrettyPrint: false, // Compact logs for storage/transmission
	})

	l.SetReportCaller(true) // Adds file:line to every log entry

	return l
}

// TempLogger For internal use (like in InitConfig)
func TempLogger() *logrus.Entry {
	tmp := newLogger()
	return logrus.NewEntry(tmp)
}
