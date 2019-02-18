package main

import "os"


func setupDatabase () (err error) {
	db, err := connectToDatabaseServer(ConnectionOptions{
		Driver: os.Getenv("DB_DRIVER"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	})

	if err != nil {
		return
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS clients (
			hashid VARCHAR(64) NOT NULL,
			ip VARCHAR(16) NOT NULL,
			peerid VARCHAR(64) NOT NULL,
			port INT,
			expires INT,
			infohash VARCHAR(64),
			useragent VARCHAR(100),
			hashKey VARCHAR(64),
			isSeeder TINYINT(1)
			PRIMARY KEY (hashid)
		) DEFAULT CHARSET=utf8;
	`)

	return
}
