package main

import "database/sql"

type ConnectionOptions struct {
	Driver string
	Host string
	Port string
	Username string
	Password string
	Database string
}

type Client struct {
	DatabaseDriver DatabaseDriverType
	CacheDriver CacheDriverType
	DBClient *sql.DB
	CacheClient interface{}
	// Should ONLY be used in rare edge cases!
	UseCache bool
}

// Open Creates the cache and db instance.
func Open(db ConnectionOptions, co ConnectionOptions) (c *Client, err error) {

	// First lets setup the connection for cache.
	c.CacheClient, err = connectToCacheServer(
		co.Driver,
		co.Host,
		co.Port,
		co.Password,
		co.Database,
	)

	if err != nil {
		return
	}

	c.DBClient, err = connectToDatabaseServer(
		db.Driver,
		db.Host,
		db.Port,
		db.Username,
		db.Password,
		db.Database,
	)

	if err != nil {
		return
	}

	return
}

// Close Closes our cache and database instance.
func (c *Client) Close() {
	c.DBClient.Close()
	c.CacheClient.Close()
}

// Query Sends a query to cache first, and fallsback to db if connection fails
// or the key simply doesn't exist.
func (c *Client) Query() {

}

