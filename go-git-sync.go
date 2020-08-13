package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func SyncFunction(c *cli.Context) error {
	fmt.Println("Will sync from", c.String("source"))
	return nil
}

func main() {

	app := &cli.App{
		Name:                 "go-git-sync",
		Usage:                "Simple tool to sync content from Git on change",
		EnableBashCompletion: true,
		Version:              "v0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "source",
				Aliases:  []string{"s"},
				Usage:    "Source Git URL to sync from",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "script",
				Aliases:  []string{"r"},
				Usage:    "Script to run on change",
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "sync",
				Aliases: []string{"s"},
				Usage:   "Perform a sync",
				Action:  SyncFunction,
			},
		},
		// Action: func(c *cli.Context) error {
		// 	if sourceUrl == "" {
		// 		return cli.Exit("Source URL is required", 1)
		// 	}
		// 	return nil
		// },
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
