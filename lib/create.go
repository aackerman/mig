package lib

import (
	"log"
	"os/exec"

	"strings"
)

func Create() {
	CreateDatabase(conf)
	CreateSchemaMigrationsTable()
}

func CreateSchemaMigrations() {
	if _, err := Get().Exec(`
    create table schema_migrations (
      version varchar(255) not null unique
    );
  `); err != nil {
		log.Println("CreateSchemaMigrationsTable", err)
	}
}

func CreateDatabase(conf DatabaseConfig) {
	cmd := exec.Command("createdb", conf.Database)
	if output, err := cmd.CombinedOutput(); err != nil {
		if !strings.Contains(string(output), "already exists") {
			log.Fatalln(string(output))
		}
	}
}
