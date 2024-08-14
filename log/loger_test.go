/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package log

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// captureOutput helps to capture the output of the logger for testing
func captureOutput(l *Logger, f func()) string {
	var buf bytes.Buffer
	writer := zerolog.ConsoleWriter{Out: &buf, NoColor: true}

	// Redirect logger output to the buffer
	l.logger = l.logger.Output(writer)

	// Call the function that generates log output
	f()

	// Return the captured output
	return buf.String()
}

func TestLogger_Info(t *testing.T) {
	logger := NewLogger(false)
	output := captureOutput(logger, func() {
		logger.Info("Info message", map[string]interface{}{
			"key": "value",
		})
	})

	assert.Contains(t, output, "Info message")
	assert.Contains(t, output, "key=value")
}

func TestLogger_Error(t *testing.T) {
	logger := NewLogger(false)
	output := captureOutput(logger, func() {
		logger.Error("Error message", map[string]interface{}{
			"error": "something went wrong",
		})
	})

	// Adjust the assertion to handle formatting
	assert.Contains(t, output, "Error message")
	assert.Contains(t, output, `error="something went wrong"`)
}

func TestLogger_Debug(t *testing.T) {
	logger := NewLogger(false)
	output := captureOutput(logger, func() {
		logger.Debug("Debug message", map[string]interface{}{
			"debug_key": "debug_value",
		})
	})

	assert.Contains(t, output, "Debug message")
	assert.Contains(t, output, "debug_key=debug_value")
}

func TestLogger_Verbose_Enabled(t *testing.T) {
	logger := NewLogger(true)
	output := captureOutput(logger, func() {
		logger.Verbose("Verbose message", map[string]interface{}{
			"verbose_key": "verbose_value",
		})
	})

	assert.Contains(t, output, "Verbose message")
	assert.Contains(t, output, "verbose_key=verbose_value")
}

func TestLogger_Verbose_Disabled(t *testing.T) {
	logger := NewLogger(false)
	output := captureOutput(logger, func() {
		logger.Verbose("Verbose message", map[string]interface{}{
			"verbose_key": "verbose_value",
		})
	})

	assert.NotContains(t, output, "Verbose message")
	assert.NotContains(t, output, "verbose_key=verbose_value")
}
