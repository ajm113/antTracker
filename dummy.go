package main

import "errors"

// Simply a dummy cache function to replicate calls to Redis/Memcached.
type DummyCache struct { }

func (d *DummyCache) Get(key string) (string, error) {
	return "", errors.New("Key unavailable!")
}

func (d *DummyCache) Set(key string, v interface{}) error {
	return nil
}
