package getenv

import (
	"os"
)

// Exists : Verifies if the given environment variable is set
func Exists(s string) bool {
	_, exists := os.LookupEnv(s)
	return exists
}

// ReadEnvVar : Returns value of given environment variable
func ReadEnvVar(s string) string {
	return os.Getenv(s)
}
