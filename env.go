package utils

import "os"

func Getenv(key string, def string) string {
	value, found := os.LookupEnv(key)

	if !found {
		value = def
	}

	return value

}
