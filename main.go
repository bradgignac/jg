package main

import (
	"fmt"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/urfave/cli"
)

// Generate null for nothing.
// Generate object with each key and null value.
// Add support for null values.
// Add support for string values.
// Add support for boolean values.
// Add support for number values.
// Add support for number e format.
// Objects?
// Arrays?
// Quoted keys?

var Version string
var Revision string

func main() {
	app := cli.NewApp()
	app.Name = "jg"
	app.Version = fmt.Sprintf("%s - Revision: %s", Version, Revision)
	app.Usage = "Generate JSON from your shell"
	app.Action = generate

	app.Run(os.Args)
}

func generate(c *cli.Context) error {
	input, err := parseAll(c.Args())
	if err != nil {
		return err
	}

	json := gabs.New()
	for _, i := range input {
		json.SetP(i.Value, i.Path)
	}

	fmt.Println(json.StringIndent("", "  "))

	return nil
}
