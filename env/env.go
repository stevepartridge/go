package env

import (
	"fmt"
	"os"
)

var vars = map[string]string{}

func Add(key, value string) {
	if os.Getenv(key) != "" {
		vars[key] = os.Getenv(key)
		fmt.Printf("Set Environment Variable: %s = %s\n", key, vars[key])
	} else {
		vars[key] = value
		fmt.Printf("Set Environment Variable (Default): %s = %s\n", key, value)
	}
}

func Get(key string) string {
	return vars[key]
}
