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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
