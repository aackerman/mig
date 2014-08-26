package postgres

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func Migrate(mpath string, db *sql.DB) {
	exefile, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	mpath = path.Join(exefile, mpath)
	pending := GetPendingMigrations(mpath, db)
	for version, filename := range pending {
		contents, err := ioutil.ReadFile(path.Join(mpath, filename))
		if err != nil {
			log.Fatalln("Readfile", filename, err)
		}
		_, err = db.Exec(string(contents))
		if err != nil {
			log.Fatalln("Exec:", filename, err)
		}
		_, err = db.Exec(`insert into schema_migrations (version) values ($1)`, version)
		if err != nil {
			log.Fatalln("Insert version", err)
		}
	}
}

func CheckForOutstandingMigrations(mpath string, db *sql.DB) {
	pending := GetPendingMigrations(mpath, db)
	if len(pending) > 0 {
		log.Println("--- Pending Migrations ---")
		for _, filename := range pending {
			log.Println(filename)
		}
		log.Fatalln("--- Pending Migrations ---")
	}
}

func GetCurrentVersion(db *sql.DB) string {
	row := db.QueryRow(`select * from schema_migrations order by version desc limit 1`)
	var version string
	if err := row.Scan(&version); err != nil {
		panic(err)
	}
	return version
}

func GetDbMigrations(db *sql.DB) []string {
	rows, err := db.Query(`select * from schema_migrations`)
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
	if rows.Err() != nil {
		panic(rows.Err())
	}
	return versions
}

func GetMigrationVersionFilenameMap(mpath string) map[string]string {
	infos, err := ioutil.ReadDir(mpath)
	if err != nil {
		log.Fatal(err)
	}
	_map := make(map[string]string)
	regex := regexp.MustCompile(`^(\d+?)_.*`)
	for _, info := range infos {
		_map[regex.FindStringSubmatch(info.Name())[1]] = info.Name()
	}
	return _map
}

func GetPendingMigrations(mpath string, db *sql.DB) map[string]string {
	pending := GetMigrationVersionFilenameMap(mpath)
	migrated := GetDbMigrations(db)
	for _, version := range migrated {
		delete(pending, version)
	}
	return pending
}
