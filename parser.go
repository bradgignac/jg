package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
)

// ParseError represents an error parsing input into a PathValuePair
type ParseError struct {
	input string
}

// PathValuePair contains a JSON path and its associated value
type PathValuePair struct {
	Type  reflect.Kind
	Path  string
	Value interface{}
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Cannot parse input into JSON path and value - %s", e.input)
}

func parse(inputs []string) ([]*PathValuePair, error) {
	parsed := make([]*PathValuePair, len(inputs))

	for i, input := range inputs {
		input, err := parseInput(input)
		if err != nil {
			return nil, err
		}
		parsed[i] = input
	}

	return parsed, nil
}

func parseInput(input string) (*PathValuePair, error) {
	parser := regexp.MustCompile(`([^=]*)=([^=]*)`)
	matches := parser.FindStringSubmatch(input)

	if len(matches) != 3 {
		return nil, &ParseError{input: input}
	}

	key := matches[1]
	kind, value := parseValue(matches[2])
	return &PathValuePair{kind, key, value}, nil
}

func parseValue(value string) (reflect.Kind, interface{}) {
	var numberVal float64
	var boolVal bool

	if value == "null" {
		return reflect.Invalid, nil
	}

	if err := json.Unmarshal([]byte(value), &numberVal); err == nil {
		return reflect.Float64, numberVal
	}

	if err := json.Unmarshal([]byte(value), &boolVal); err == nil {
		return reflect.Bool, boolVal
	}

	return reflect.String, value
}
