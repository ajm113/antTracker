package main

import (
	"database/sql"
	"crypto/sha1"
	"io"
	"fmt"
)

type Client struct {
	HashID string
	IP string
	PeerID string
	Port string
	Expires int64
	InfoHash string
	UserAgent string
	Key string
	IsSeeder bool
}

func GetClientByHashId(db *sql.DB, h string) (c Client, err error) {
	rows, err := db.Query(`
	SELECT
		hashid,
		ip,
		peerid,
		port,
		expires,
		infohash,
		useragent,
		hashKey,
		isSeeder
	FROM clients
	WHERE hashid=?
	LIMIT 1
	`, h)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&c.HashID,
			&c.IP,
			&c.PeerID,
			&c.Port,
			&c.Expires,
			&c.InfoHash,
			&c.UserAgent,
			&c.Key,
			&c.IsSeeder,
		)
	}
	err = rows.Err()

	return
}

func InsertOrUpdateClient(db *sql.DB, c Client) (rc Client, err error) {

	return
}

func RemoveClient(db *sql.DB, c Client) (err error) {
	_, err = db.Exec(`
		DELETE FROM clients WHERE hashid=? LIMIT 1
	`, c.HashID)

	return
}

func GetClients(db *sql.DB) (c[] Client, err error) {
	return
}

func CreateHashIdIntoClient(c Client) (r Client) {
	h := sha1.New()
	io.WriteString(h, c.PeerID + c.InfoHash)
	c.HashID = fmt.Sprintf("% x", h.Sum(nil))
	r = c
	return
}
