package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stevepartridge/go/log"
)

type Database struct {
	Id         string
	Host       string
	Port       string
	Name       string
	User       string
	Pass       string
	SSLMode    string
	Connection *sql.DB
}

type Postgres struct{}

var databases []Database

func Create() *Postgres {
	return &Postgres{}
}

func (pg *Postgres) Get(id string) *sql.DB {
	for _, pgdb := range databases {
		if pgdb.Id == id {
			pgdb.Connection = pg.connect(pgdb)
			return pgdb.Connection
		}
	}
	log.Fatal("Database id (", id, ") not found")
	return nil
}

func (pg *Postgres) connect(database Database) *sql.DB {
	if database.Connection == nil {

		connection := fmt.Sprintf(
			"host=%s dbname=%s user=%s password='%s' port=%s sslmode=%s",
			database.Host,
			database.Name,
			database.User,
			database.Pass,
			database.Port,
			database.SSLMode,
		)
		conn, err := sql.Open("postgres", connection)
		if err == nil {
			database.Connection = conn
			return database.Connection
		} else {
			panic(err)
		}
	} else {
		return database.Connection
	}
}

func (pg *Postgres) Add(database Database) {
	databases = append(databases, database)
}
