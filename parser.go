package main

import (
	"fmt"
	"regexp"
)

// PathValuePair contains a JSON path and its associated value.
type PathValuePair struct {
	Path  string
	Value string
}

// ParseError represents an error parsing input into a PathValuePair.
type ParseError struct {
	input string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Cannot parse input into JSON path and value - %s", e.input)
}

func parseAll(inputs []string) ([]*PathValuePair, error) {
	parsed := make([]*PathValuePair, len(inputs))

	for i, input := range inputs {
		input, err := parse(input)
		if err != nil {
			return nil, err
		}
		parsed[i] = input
	}

	return parsed, nil
}

func parse(input string) (*PathValuePair, error) {
	parser := regexp.MustCompile(`([^=]*)=([^=]*)`)
	matches := parser.FindStringSubmatch(input)

	if len(matches) != 3 {
		return nil, &ParseError{input: input}
	}

	return &PathValuePair{Path: matches[1], Value: matches[2]}, nil
}
