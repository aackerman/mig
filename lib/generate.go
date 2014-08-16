package lib

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"bitbucket.org/pkg/inflect"
)

func Generate(basename string) {
	if basename == "" {
		log.Fatalln("filename is empty")
	}
	basename = inflect.Underscore(basename)
	_, currentfile, _, _ := runtime.Caller(1)
	filename := path.Join(currentfile, "..", "db", "migrate", NextFilename(basename))
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
