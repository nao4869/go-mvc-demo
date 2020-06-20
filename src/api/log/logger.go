package log

import (
	"os"
	"strings"

	"../config"
	"github.com/Sirupsen/logrus"
)

var (
	// Log -
	Log *logrus.Logger
)

func init() {
	level, error := logrus.ParseLevel(config.LogLevel)
	if error != nil {
		level = logrus.DebugLevel
	}

	Log = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}

	if config.IsProduction() {
		// json for elastic search in production
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{}
	}
}

// Info -
func Info(message string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(message)
}

// Error -
func Error(message string, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Error(message)
}

// Debug -
func Debug(message string, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Error(message)
}

func parseFields(tags ...string) logrus.Fields {
	// allocation memory & creating the map with a size equals to length of tags
	result := make(logrus.Fields, len(tags))

	for _, tag := range tags {
		elements := strings.Split(tag, ":")
		result[strings.TrimSpace(elements[0])] = strings.TrimSpace(elements[1])
	}
	return result
}
