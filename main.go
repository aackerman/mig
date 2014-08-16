package main

import (
	"fmt"
	"mig/lib"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "mig"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: "",
			Usage: "specify database configuration file",
		},
		cli.StringFlag{
			Name:  "migrations, m",
			Value: "",
			Usage: "specify migration files location",
		},
	}

	app.Flags = flags

	app.Commands = []cli.Command{
		{
			Name:  "db:create",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				conf := lib.GetConfig("")
				db := lib.Connect(conf)
				lib.Create(db, conf)
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
