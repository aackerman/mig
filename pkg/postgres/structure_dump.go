package postgres

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func StructureDump(dumppath string, conf DatabaseConfig) {
	cmd := exec.Command("pg_dump", "-s", conf.Database)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	exefile, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(filepath.Join(exefile, dumppath, "schema.sql"))
	f.Write(output)
	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
