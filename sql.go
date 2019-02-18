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

type ConnectionOptions struct {
	Driver string
	Host string
	Port string
	Username string
	Password string
	Database string
}

func databaseDriverTypeStringToInt(t string) (ct DatabaseDriverType, err error) {
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
		err = errors.New("Database driver is unsupported!")
	}

	return
}

func createConnectionStringByConnectionOptions(co ConnectionOptions) (cstr string, err error) {
	t, err := databaseDriverTypeStringToInt(co.Driver)

	if err != nil {
		return
	}

	switch (t) {
	case SQLITE3:
		cstr = co.Host
	case MYSQL:
		cstr = co.Username + ":" + co.Password + "@" + co.Host + ":" + co.Port + "/" + co.Database
	case PG:
		cstr = "postgres://" + co.Username + ":" + co.Password + "@" +
			co.Host + ":" + co.Port + " /" + co.Database + "?sslmode=verify-full"
	}

	return
}

func connectToDatabaseServer(co ConnectionOptions) (db *sql.DB, err error) {

	t, err := databaseDriverTypeStringToInt(co.Driver)

	if err != nil {
		return
	}

	sqlDriverName, err := databaseDriverTypeToString(t)

	if err != nil {
		return
	}

	cstr, err := createConnectionStringByConnectionOptions(co)

	if err != nil {
		return
	}

	db, err = sql.Open(
		sqlDriverName,
		cstr,
	)
	return
}
