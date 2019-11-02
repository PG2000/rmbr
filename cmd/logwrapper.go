package cmd

import (
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

var (
	invalidArgMessage        = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage   = Event{2, "Invalid value for argument: %s: %v"}
	missingArgMessage        = Event{3, "Missing arg: %s"}
	unmarshalingErrorMessage = Event{4, "Error while unmarshal: %s: %v"}
	commandInfoMessage       = Event{5, "Running command: %s with args: %v"}
	commonError              = Event{6, "Error: %v"}
)

func (l *StandardLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

func (l *StandardLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}

func (l *StandardLogger) UnmarshalingErrorMessage(argumentName error) {
	l.Errorf(unmarshalingErrorMessage.message, argumentName)
}

func (l *StandardLogger) ErrorCommon(argumentName error) {
	l.Errorf(commonError.message, argumentName)
}

func (l *StandardLogger) InfoCommandMessage(commandName string, arguments ...string) {
	l.Infof(commandInfoMessage.message, commandName, arguments)
}
