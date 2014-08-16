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
			Usage: "run outstanding migrations",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:rollback",
			Usage: "rollback last migration",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:structure:dump",
			Usage: "dump the structure of the database",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:version",
			Usage: "print the last migrated version",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:drop",
			Usage: "drop the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "db:test:prepare",
			Usage: "run migrations against test database",
			Flags: flags,
			Action: func(c *cli.Context) {
				fmt.Printf("%#v", c.Args())
			},
		},
		{
			Name:  "generate",
			Usage: "generate a new migration",
			Flags: flags,
			Action: func(c *cli.Context) {
				lib.Generate(c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
