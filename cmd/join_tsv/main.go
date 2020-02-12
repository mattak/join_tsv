package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"github.com/mattak/join_tsv/pkg/join_tsv"
	"os"
	"strconv"
	"strings"
)

func extractKeys(c *cli.Context) []int {
	argsLength := c.Args().Len()
	keysString := c.String("keys")
	keysStrings := strings.Split(keysString, ",")
	keys := make([]int, argsLength)

	if len(keysString) < 1 {
		for i := 0; i < argsLength; i++ {
			keys[i] = 0
		}
		return keys
	}

	if len(keysStrings) != argsLength {
		panic("not matched keys length with input tsv length")
	}

	for i := 0; i < argsLength; i++ {
		k, err := strconv.ParseInt(keysStrings[i], 10, 64)
		if err != nil {
			panic(err)
		}

		keys[i] = int(k) - 1
	}

	return keys
}

func run(c *cli.Context) bool {
	keys := extractKeys(c)
	files := c.Args().Slice()

	result := join_tsv.JoinTablesByFile(keys, files)
	join_tsv.PrintTableTsv(result)

	return true
}

func main() {
	app := cli.NewApp()
	app.Name = "join_tsv"
	app.Usage = "join tsv multiple data with specified column"
	app.Version = "0.0.1"
	app.ArgsUsage = "[reference_tsv] [join_tsv]+"
	app.Commands = nil
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "keys",
			Aliases: []string{"k"},
			Usage:   "keys to join tsv data. start from 1. e.g. -k 1,1,5",
		},
	}
	app.Action = func(c *cli.Context) error {
		if run(c) {
			return nil
		}

		return errors.New("unknown errors happened")
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
