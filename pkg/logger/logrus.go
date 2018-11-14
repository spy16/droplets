package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// New returns a logger implemented using the logrus package.
func New(wr io.Writer, level string, format string) Logger {
	if wr == nil {
		wr = os.Stderr
	}

	lr := logrus.New()
	lr.SetOutput(wr)
	lr.SetFormatter(&logrus.TextFormatter{})
	if format == "json" {
		lr.SetFormatter(&logrus.JSONFormatter{})
	}

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.WarnLevel
		lr.Warnf("failed to parse log-level '%s', defaulting to 'warning'", level)
	}
	lr.SetLevel(lvl)

	return &logrusLogger{
		entry: logrus.NewEntry(lr),
	}
}

// logrusLogger provides functions for structured logging.
type logrusLogger struct {
	entry *logrus.Entry
}

func (ll *logrusLogger) Debugf(msg string, args ...interface{}) {
	ll.entry.Debugf(msg, args...)
}

func (ll *logrusLogger) Infof(msg string, args ...interface{}) {
	ll.entry.Infof(msg, args...)
}

func (ll *logrusLogger) Warnf(msg string, args ...interface{}) {
	ll.entry.Warnf(msg, args...)
}

func (ll *logrusLogger) Errorf(msg string, args ...interface{}) {
	ll.entry.Errorf(msg, args...)
}

func (ll *logrusLogger) WithFields(fields map[string]interface{}) Logger {
	annotatedEntry := ll.entry.WithFields(logrus.Fields(fields))
	return &logrusLogger{
		entry: annotatedEntry,
	}
}
