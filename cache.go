package main

import (
	"strconv"

	"github.com/go-redis/redis"
	"github.com/bradfitz/gomemcache/memcache"
)

type CacheDriverType int

const (
	DUMMY CacheDriverType = 0
	REDIS CacheDriverType = 1
	MEMCACHED CacheDriverType = 2
)

func cacheTypeStringToInt(t string) CacheDriverType {
	switch (t) {
	case "redis":
		return REDIS
	case "memcached":
		return MEMCACHED
	default:
		return DUMMY
	}
}

func connectToCacheServer(cacheType, host, port, password, database string) (c interface{}, err error) {
	switch (cacheTypeStringToInt(cacheType)) {
	case DUMMY:
		return
	case REDIS:
		db, err := strconv.Atoi(database)

		c = redis.NewClient(&redis.Options{
			Addr: host + ":" + port,
			Password: password,
			DB: db,
		})
	case MEMCACHED:
		c = memcache.New(host + ":" + port)
	}

	return
}
