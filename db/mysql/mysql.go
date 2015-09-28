package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stevepartridge/go/log"
)

type Database struct {
	Id         string
	Host       string
	Port       string
	Name       string
	User       string
	Pass       string
	Connection *sql.DB
}

type Mysql struct{}

var databases []Database

func Create() *Mysql {
	return &Mysql{}
}

func (m *Mysql) Get(id string) *sql.DB {
	for _, mydb := range databases {
		if mydb.Id == id {
			mydb.Connection = m.connect(mydb)
			return mydb.Connection
		}
	}
	log.Fatal("db.mysql.ERROR", id, "not found")

	return nil
}

func (m *Mysql) connect(database Database) *sql.DB {
	if database.Connection == nil {

		var host = database.Host

		if database.Port != "" {
			host = host + ":" + database.Port
		}

		connection := fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			database.User,
			database.Pass,
			host,
			database.Name,
		)

		conn, err := sql.Open("mysql", connection)
		if err == nil {
			database.Connection = conn
			return database.Connection
		} else {
			log.IfError(err)
			return nil
		}
	} else {
		return database.Connection
	}
}

func (m *Mysql) Add(database Database) {
	databases = append(databases, database)
	log.Info("db.mysql.Add", database.Id)
}
