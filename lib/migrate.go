package lib

import (
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"runtime"
	"strconv"
)

func Migrate() {
	_, currentfile, _, _ := runtime.Caller(1)
	pending := GetPendingMigrations()
	for version, filename := range pending {
		contents, err := ioutil.ReadFile(path.Join(path.Dir(currentfile), "migrations", filename))
		if err != nil {
			log.Fatalln("Readfile", filename, err)
		}
		_, err = Get().Exec(string(contents))
		if err != nil {
			log.Fatalln("Exec:", filename, err)
		}
		_, err = Get().Exec(`insert into schema_migrations (version) values ($1)`, version)
		if err != nil {
			log.Fatalln("Insert version", err)
		}
	}
}

func CheckForOutstandingMigrations() {
	pending := GetPendingMigrations()
	if len(pending) > 0 {
		log.Println("--- Pending Migrations ---")
		for _, filename := range pending {
			log.Println(filename)
		}
		log.Fatalln("--- Pending Migrations ---")
	}
}

func GetDbMigrations() []string {
	rows, err := Get().Query(`select * from schema_migrations`)
	if err != nil {
		log.Fatal("Get schema migrations ", err)
	}
	var versions []string
	var version string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&version)
		versions = append(versions, version)
	}
	return versions
}

func GetMigrationVersionFilenameMap() map[string]string {
	_, currentfile, _, _ := runtime.Caller(1)
	infos, err := ioutil.ReadDir(path.Join(path.Dir(currentfile), "migrations"))
	if err != nil {
		log.Fatal(err)
	}
	_map := make(map[string]string)
	regex := regexp.MustCompile(`^(\d+?)_.*`)
	for _, info := range infos {
		_map[regex.FindString(info.Name())] = info.Name()
	}
	return _map
}

func GetPendingMigrations() map[string]string {
	pending := GetMigrationVersionFilenameMap()
	migrated := GetDbMigrations()
	for version, _ := range migrated {
		delete(pending, strconv.Itoa(version))
	}
	return pending
}
