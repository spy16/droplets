// Package logger provides logging functions. The loggers implemented in this
// package will have the API defined by the Logger interface. The interface
// is defined here (instead of where it is being used which is the right place),
// is because Logger interface is a common thing that gets used across the code
// base while being fairly constant in terms of its API.
package logger

// Logger implementation is responsible for providing structured and levled
// logging functions.
type Logger interface {
	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})

	// WithFields should return a logger which is annotated with the given
	// fields. These fields should be added to every logging call on the
	// returned logger.
	WithFields(m map[string]interface{}) Logger
}
