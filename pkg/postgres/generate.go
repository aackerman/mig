package postgres

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"bitbucket.org/pkg/inflect"
)

func Generate(mpath string, basename string) {
	if basename == "" {
		log.Fatalln("filename is empty")
	}
	basename = inflect.Underscore(basename)
	exefile, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	filename := path.Join(exefile, mpath, NextFilename(basename))
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func NextFilename(basename string) string {
	return CreateVersion() + "_" + basename + ".sql"
}

func CreateVersion() string {
	now := time.Now()
	parts := []int{
		now.Year(),
		int(now.Month()),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	}
	stamp := ""
	for _, part := range parts {
		stamp += fmt.Sprintf("%02d", part)
	}

	return stamp
}
