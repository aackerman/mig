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
		cli.StringFlag{
			Name:  "env, e",
			Value: "development",
			Usage: "specify migration files location",
		},
		cli.StringFlag{
			Name:  "schema",
			Value: "db",
			Usage: "specify location for schema dump",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "db:create",
			Usage: "create the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				conf := lib.GetConfig(c.String("conf"), c.String("env"))
				db := lib.Connect(conf)
				lib.Create(db, conf)
			},
		},
		{
			Name:  "db:migrate",
			Usage: "run outstanding migrations",
			Flags: flags,
			Action: func(c *cli.Context) {
				conf := lib.GetConfig(c.String("conf"), c.String("env"))
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
				conf := lib.GetConfig(c.String("conf"), c.String("env"))
				lib.StructureDump(c.String("schema"), conf)
			},
		},
		{
			Name:  "db:version",
			Usage: "print the last migrated version",
			Flags: flags,
			Action: func(c *cli.Context) {
				conf := lib.GetConfig(c.String("conf"), c.String("env"))
				db := lib.Connect(conf)
				log.Println(lib.GetCurrentVersion(db))
			},
		},
		{
			Name:  "db:drop",
			Usage: "drop the db",
			Flags: flags,
			Action: func(c *cli.Context) {
				conf := lib.GetConfig(c.String("conf"), c.String("env"))
				lib.Drop(conf)
			},
		},
		{
			Name:  "db:test:prepare",
			Usage: "run migrations against test database",
			Flags: flags,
			Action: func(c *cli.Context) {
				conf := lib.GetConfig(c.String("conf"), "test")
				db := lib.Connect(conf)
				lib.Drop(conf)
				lib.Create(db, conf)
				lib.Migrate(c.String("migrations"), db)
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
