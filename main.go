package main

import (
	"fmt"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/urfave/cli"
)

// Add support for null values.
// Add support for boolean values.
// Add support for number values.
// Add support for number e format.
// Objects?
// Arrays?
// Quoted keys?

// Version of the CLI injected by the build.
var Version string

// Revision of the CLI injected by the build.
var Revision string

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
