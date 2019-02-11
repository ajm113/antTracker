package main

import "os"

func Getenv(key, fallbackValue string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallbackValue;
	}

	return value
}
