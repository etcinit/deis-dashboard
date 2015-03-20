package util

import "os"

// Getopt gets an environment variable or a default value
func Getopt(name, dfault string) string {
	value := os.Getenv(name)

	if value == "" {
		value = dfault
	}

	return value
}

// Hasenv checks if an environment variable is defined
func Hasenv(names []string) bool {
	for _, opt := range names {
		if os.Getenv(opt) == "" {
			return false
		}
	}

	return true
}
