package utils

import (
    "bytes"
    "os"
    "os/exec"
    "strconv"
    "strings"
    "syscall"
)

// IsTrue returns true if the string value is one of the true strings.
func IsTrue(value string) (ret bool) {
    ret = false
    trueValues := [7]string{
        "true",
        "True",
        "TRUE",
        "yes",
        "Yes",
        "YES",
        "1",
    }

    for _, v := range trueValues {
        if value == v {
            ret = true
            break
        }
    }

    return
}

// GetEnv returns the value of the environment variable or provided fallback.
// value if the environment variable is not defined.
func GetEnv(key string, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }

    return fallback
}

// GetEnvBool is the same like GetEnv but for boolean values.
func GetEnvBool(key string, fallback bool) bool {
    if value, ok := os.LookupEnv(key); ok {
        return IsTrue(value)
    }

    return fallback
}

// GetEnvInt is the same like GetEnv but for integer values.
func GetEnvInt(key string, fallback int) int {
    if value, ok := os.LookupEnv(key); ok {
        val, err := strconv.Atoi(value)

        if err != nil {
            return fallback
        }

        return val
    }

    return fallback
}

// GetEnvFloat is the same like GetEnv but for float values.
func GetEnvFloat(key string, fallback float64) float64 {
    if value, ok := os.LookupEnv(key); ok {
        val, err := strconv.ParseFloat(value, 64)

        if err != nil {
            return fallback
        }

        return val
    }

    return fallback
}

const defaultFailedCode = 1

// Run runs a command and returns the exit code, stdour and stderr output.
func Run(args []string) (rc int, stdout string, stderr string) {
    var outbuf, errbuf bytes.Buffer
    cmd := exec.Command(args[0], args[1:]...)
    cmd.Stdout = &outbuf
    cmd.Stderr = &errbuf

    err := cmd.Run()
    stdout = outbuf.String()
    stderr = errbuf.String()

    if err != nil {
        if exitError, ok := err.(*exec.ExitError); ok {
            ws := exitError.Sys().(syscall.WaitStatus)
            rc = ws.ExitStatus()
        } else {
            rc = defaultFailedCode

            if stderr == "" {
                stderr = err.Error()
            }
        }
    } else {
        ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
        rc = ws.ExitStatus()
    }

    stdout = strings.TrimSpace(stdout)
    stderr = strings.TrimSpace(stderr)

    return
}
