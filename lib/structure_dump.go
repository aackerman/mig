package lib

import (
	"log"
	"os"
	"os/exec"
)

func StructureDump(conf DatabaseConfig) {
	cmd := exec.Command("pg_dump", "-s", conf.Database)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	f, err := os.Create("schema.sql")
	f.Write(output)
	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
