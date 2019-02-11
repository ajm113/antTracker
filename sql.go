package main

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseDriverType int

const (
	SQLITE3 DatabaseDriverType = 0
	MYSQL DatabaseDriverType = 1
	PG DatabaseDriverType = 2
)

func databaseDriverTypeStringToInt(t string) (ct CacheDriverType, err error) {
	switch (t) {
	case "sqlite3":
		ct = SQLITE3
	case "mysql", "mariadb":
		ct = MYSQL
	case "pg", "postgres":
		ct = PG
	default:
		err = errors.New("Database driver " + t + " is unsupported!")
	}

	return
}

func databaseDriverTypeToString(t DatabaseDriverType) (ct string, err error) {
	switch (t) {
	case SQLITE3:
		ct = "sqlite3"
	case MYSQL:
		ct = "mysql"
	case PG:
		ct = "postgres"
	default:
		err = errors.New("Database driver " + t + " is unsupported!")
	}

	return
}

func createConnectionStringByDatabaseDriverType(t DatabaseDriverType, host, port, username, password, database string) string {
	switch (t) {
	case SQLITE3:
		return host
	case MYSQL:
		return username +":" + password + "@" + host + ":" + port + "/" + database
	case PG:
		return "postgres://" + username + ":" + password + "@" +
			host + ":" + port + " /" + database + "?sslmode=verify-full"
	}
}

func connectToDatabaseServer(dbType, host, port, username, password, database string) (db *sql.DB, err error) {

	t, err = databaseDriverTypeStringToInt(dbType)

	if err != nil {
		return
	}

	db, err = sql.Open(
		databaseDriverTypeToString(t),
		createConnectionStringByDatabaseDriverType(t, host, port, username, password, database),
	)
	return
}
