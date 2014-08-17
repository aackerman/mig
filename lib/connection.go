package lib

import (
	"database/sql"
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Username string
	Password string
	Database string
	Hostname string
}

func (d *DatabaseConfig) IsValid() bool {
	return (d.Username != "" &&
		d.Password != "" &&
		d.Database != "" &&
		d.Hostname != "")
}

func Connect(conf DatabaseConfig) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s", conf.Username, conf.Database, "disable"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetConfig(filepath string, env string) DatabaseConfig {
	if filepath == "" {
		filepath = "config/database.toml"
	}
	_, currentfile, _, _ := runtime.Caller(1) // get current file path
	abspath := path.Join(path.Dir(currentfile), filepath)
	tmpconf := map[string]DatabaseConfig{}
	if _, err := toml.DecodeFile(abspath, &tmpconf); err != nil {
		log.Fatal(err)
	}

	if conf, ok := tmpconf[env]; ok {
		return conf
	} else {
		panic(env + " configuration is not present in " + filepath)
	}

	return tmpconf[env]
}
