package postgres

import (
	"log"
	"os/exec"
	"strings"
)

func Drop(conf DatabaseConfig) {
	DropDatabase(conf)
}

func DropDatabase(conf DatabaseConfig) {
	cmd := exec.Command("dropdb", conf.Database)
	if output, err := cmd.CombinedOutput(); err != nil {
		if !strings.Contains(string(output), "does not exist") {
			log.Fatalln(string(output), err)
		}
	}
}
