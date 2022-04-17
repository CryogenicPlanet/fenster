package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:                 "fenster",
		EnableBashCompletion: true,
		Description:          "Better fullscreen spaces",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Starts fenster service",
				Action: func(c *cli.Context) error {

					return StartServer()
				},
			},
			{
				Name:  "prepare",
				Usage: "Collects data about spaces before the full screen",
				Action: func(c *cli.Context) error {
					fmt.Println("Not implement yet")
					return nil
				},
			},
			{
				Name:        "space",
				Description: "Handling spaces",
				Subcommands: []*cli.Command{
					{
						Name:  "next",
						Usage: "Move current app to next avalibale window",
						Action: func(c *cli.Context) error {
							return HandleRight()
						},
					},
					{
						Name:  "previous",
						Usage: "Move current app to previous avalibale window",
						Action: func(c *cli.Context) error {
							return HandleLeft()
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
