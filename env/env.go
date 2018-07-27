package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var vars = map[string]string{}

// Add Environment Variable to session
func Add(key, value string) {

	if os.Getenv(key) == "" {
		vars[key] = value
		fmt.Printf("Set Environment Variable (Default): %s = %s\n", key, value)
		return
	}

	vars[key] = os.Getenv(key)
	fmt.Printf("Set Environment Variable: %s = %s\n", key, vars[key])

}

// Get Environment Variable or grab default if set
func Get(key string) string {

	if vars[key] == "" {

		v := os.Getenv(key)

		if v != "" {
			vars[key] = v
			return vars[key]
		}

		return ""
	}

	return vars[key]
}

// Set Environment Variable
func Set(key, value string) error {

	vars[key] = value

	return os.Setenv(key, value)
}

// Get env var as boolean helper
func GetAsBool(key string) bool {

	val := strings.ToLower(strings.TrimSpace(Get(key)))

	switch val {
	case "0", "no", "false":
		return false
	default:
		if val != "" {
			return true
		}
		return false
	}

}

// Get env var as int helper
func GetAsInt(key string) int {

	i, _ := strconv.Atoi(strings.ToLower(strings.TrimSpace(Get(key))))

	return i

}
