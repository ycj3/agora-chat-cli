/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package util

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/thediveo/enumflag/v2"
)

type OutputFormat enumflag.Flag

const (
	OutputFormatTable OutputFormat = iota
	OutputFormatJSON
)

func OutputFormatFromString(value string) (OutputFormat, error) {
	switch value {
	case "json":
		return OutputFormatJSON, nil
	case "text":
		return OutputFormatTable, nil
	default:
		return OutputFormatJSON, fmt.Errorf("invalid output format '%s'", value)
	}
}

// Print function that handles both JSON and table outputs.
func Print[T any](data []T, format OutputFormat, headers []any) error {
	switch format {
	case OutputFormatJSON:
		OutputJson(data)
	case OutputFormatTable:
		OutputTable(data, headers)
	}
	return nil
}

func OutputJson(data any) error {
	// Marshal the data to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %w", err)
	}
	fmt.Println(string(jsonData))
	return nil
}

func OutputTable[T any](data []T, headers []any) error {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(headers)
	for _, item := range data {
		row, ok := convertToRow(item)
		if ok {
			t.AppendRow(row)
		}
	}
	t.Render()
	return nil
}

func convertToRow(v any) ([]interface{}, bool) {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Struct:
		row := make([]interface{}, val.NumField())
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			row[i] = fmt.Sprintf("%v", field.Interface())
		}
		return row, true

	case reflect.Slice:
		// Check if it's a slice of strings
		if val.Type().Elem().Kind() == reflect.String {
			row := make([]interface{}, val.Len())
			for i := 0; i < val.Len(); i++ {
				row[i] = val.Index(i).Interface()
			}
			return row, true
		}
	}

	return nil, false
}
