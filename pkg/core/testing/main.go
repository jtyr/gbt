package testing

import (
    "os"
    "strings"
)

var ResetEnvIgnore = []string{"PATH", "HOME"}

func getEnvKeys() []string {
    var keys []string

    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        keys = append(keys, pair[0])
    }

    return keys
}

func ResetEnv() {
    for _, k := range getEnvKeys() {
        ignore := false

        for _, i := range ResetEnvIgnore {
            if i == k {
                ignore = true

                break
            }
        }

        if ! ignore {
            os.Unsetenv(k)
        }
    }
}
