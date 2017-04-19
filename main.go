package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/Jeffail/gabs"
	"github.com/urfave/cli"
)

// Objects?
// Arrays?
// Quoted keys?

// Version of the CLI injected by the build.
var Version string

// Revision of the CLI injected by the build.
var Revision string

// PathValuePair contains a JSON path and its associated value
type PathValuePair struct {
	Type  reflect.Kind
	Path  string
	Value interface{}
}

// ParseError represents an error parsing input into a PathValuePair
type ParseError struct {
	input string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Cannot parse input into JSON path and value - %s", e.input)
}

func main() {
	app := cli.NewApp()
	app.Name = "jg"
	app.Version = fmt.Sprintf("%s - Revision: %s", Version, Revision)
	app.Usage = "Generate JSON from your shell"
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "pretty, p"},
	}
	app.Action = generate

	app.Run(os.Args)
}

func generate(c *cli.Context) error {
	input, err := parse(c.Args())
	if err != nil {
		return err
	}

	json := gabs.New()
	for _, i := range input {
		json.SetP(i.Value, i.Path)
	}

	var output string
	if c.Bool("pretty") {
		output = json.StringIndent("", "  ")
	} else {
		output = json.String()
	}

	fmt.Println(output)

	return nil
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
