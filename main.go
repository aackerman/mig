package main

import (
	"log"
	"mig/lib"
	"os"
	"path"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "mig"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: path.Join("config", "database.toml"),
			Usage: "specify database configuration file",
		},
		cli.StringFlag{
			Name:  "migrations, m",
			Value: path.Join("db", "migrate"),
			Usage: "specify migration files location",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "db:create",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				conf := lib.GetConfig(c.String("conf"), "development")
				db := lib.Connect(conf)
				lib.Create(db, conf)
			},
		},
		{
			Name:  "db:migrate",
			Usage: "run outstanding migrations",
			Flags: flags,
			Action: func(c *cli.Context) {
				conf := lib.GetConfig(c.String("conf"), "development")
				db := lib.Connect(conf)
				lib.Migrate(c.String("migrations"), db)
			},
		},
		{
			Name:  "db:rollback",
			Usage: "rollback last migration",
			Flags: flags,
			Action: func(c *cli.Context) {
				log.Fatalln("Not implemented")
			},
		},
		{
			Name:  "db:structure:dump",
			Usage: "dump the structure of the database",
			Flags: flags,
			Action: func(c *cli.Context) {
				log.Fatalln("Not implemented")
			},
		},
		{
			Name:  "db:version",
			Usage: "print the last migrated version",
			Flags: flags,
			Action: func(c *cli.Context) {
				log.Fatalln("Not implemented")
			},
		},
		{
			Name:  "db:drop",
			Usage: "drop the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				conf := lib.GetConfig(c.String("conf"), "development")
				lib.Drop(conf)
			},
		},
		{
			Name:  "db:test:prepare",
			Usage: "run migrations against test database",
			Flags: flags,
			Action: func(c *cli.Context) {
				log.Fatalln("Not implemented")
			},
		},
		{
			Name:      "generate",
			ShortName: "g",
			Usage:     "generate a new migration",
			Flags:     flags,
			Action: func(c *cli.Context) {
				lib.Generate(c.String("migrations"), c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
