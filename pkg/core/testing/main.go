package testing

import (
	"os"
	"strings"
)

// ResetEnvIgnore holds list of env vars that won't be reset.
var ResetEnvIgnore = []string{"PATH", "HOME"}

func getEnvKeys() []string {
	var keys []string

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		keys = append(keys, pair[0])
	}

	return keys
}

// ResetEnv resets env vars.
func ResetEnv() {
	for _, k := range getEnvKeys() {
		ignore := false

		for _, i := range ResetEnvIgnore {
			if i == k {
				ignore = true

				break
			}
		}

		if !ignore {
			os.Unsetenv(k)
		}
	}
}
