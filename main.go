package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "migrate"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: "",
			Usage: "Database configuration file",
		},
	}

	app.Flags = flags

	app.Commands = []cli.Command{
		{
			Name:  "db:create",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:migrate",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:rollback",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:structure:dump",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:version",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:drop",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:setup",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:test:prepare",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
	}

	app.Run(os.Args)
}
