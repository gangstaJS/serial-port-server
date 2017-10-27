package main


import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync/atomic"
)

var (
	atomicDB atomic.Value
)

func initDb() {
	var db *sql.DB
	var err error

	db, err = sql.Open("sqlite3", "./serial.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	setDB(db)

	sqlStmt := `CREATE TABLE IF NOT EXISTS 'log_objects'
		(
		'uid' INTEGER PRIMARY KEY AUTOINCREMENT,
		'name' VARCHAR(128) NULL,
		'state' BOOL NULL,
		'value' INTEGER,
		'automated' BOOL NULL,
		'timestamp' DATETIME NULL DEFAULT CURRENT_TIMESTAMP
	)`

	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func setItem(serialData map[string]interface{}) {
	stmt, err := getDB().Prepare("INSERT INTO log_objects(name, state, value, automated) values(?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(serialData["name"], serialData["state"], serialData["value"], serialData["automated"])

	defer stmt.Close()
}

func getAll() {

}

func setDB(db *sql.DB) {
	atomicDB.Store(db)
}

func getDB() *sql.DB {
	return atomicDB.Load().(*sql.DB)
}
