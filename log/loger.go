/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	logger  zerolog.Logger
	verbose bool
}

// NewLogger initializes and returns a new Logger instance
func NewLogger(verbose bool) *Logger {
	// Set up zerolog to output human-readable logs
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	return &Logger{
		logger:  logger,
		verbose: verbose,
	}
}

// Info logs an informational message
func (l *Logger) Info(message string, fields map[string]interface{}) {
	event := l.logger.Info()
	for k, v := range fields {
		event.Interface(k, v)
	}
	event.Msg(message)
}

// Error logs an error message
func (l *Logger) Error(message string, fields map[string]interface{}) {
	event := l.logger.Error()
	for k, v := range fields {
		event.Interface(k, v)
	}
	event.Msg(message)
}

// Debug logs a debug message
func (l *Logger) Debug(message string, fields map[string]interface{}) {
	event := l.logger.Debug()
	for k, v := range fields {
		event.Interface(k, v)
	}
	event.Msg(message)
}

// Verbose logs a message only if verbose mode is enabled
func (l *Logger) Verbose(message string, fields map[string]interface{}) {
	if l.verbose {
		event := l.logger.Debug() // You might want to log verbose messages at the debug level
		for k, v := range fields {
			event.Interface(k, v)
		}
		event.Msg(message)
	}
}
