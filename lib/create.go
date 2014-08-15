package lib

import (
	"database/sql"
	"log"
	"os/exec"

	"strings"
)

func Create(db *sql.DB, conf DatabaseConfig) {
	CreateDatabase(conf)
	CreateSchemaMigrations(db)
}

func CreateSchemaMigrations(db *sql.DB) {
	if _, err := db.Exec(`
    create table schema_migrations (
      version varchar(255) not null unique
    );
  `); err != nil && !strings.Contains(err.Error(), "already exists") {
		log.Fatalln(err)
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
