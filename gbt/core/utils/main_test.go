package utils

import (
    "fmt"
    "os"
    "testing"
)

func getTestPrefix(testN int) string {
    return fmt.Sprintf("Test [%d]: ", testN)
}

func TestGetEnv(t *testing.T) {
    tests := []struct {
        name string
        fallback string
        expectedVal string
        set bool
    }{
        { name: "XXX", fallback: "", expectedVal: "", set: false, },
        { name: "YYY", fallback: "", expectedVal: "test", set: true, },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)

        if test.set {
            os.Setenv(test.name, test.expectedVal)
        }

        val := GetEnv(test.name, test.fallback)

        if val != test.expectedVal {
            t.Errorf("%sExpected '%s', found '%s'.", testPrefix, test.expectedVal, val)
        }

        if test.set {
            os.Unsetenv(test.name)
        }
    }
}

func TestGetEnvBool(t *testing.T) {
    tests := []struct {
        name string
        fallback bool
        expectedVal bool
        set bool
        setVal string
    }{
        { name: "XXX", fallback: false, expectedVal: false, set: false, setVal: "" },
        { name: "YYY", fallback: true, expectedVal: true, set: true, setVal: "1" },
        { name: "YYY", fallback: true, expectedVal: true, set: true, setVal: "true" },
        { name: "YYY", fallback: true, expectedVal: true, set: true, setVal: "True" },
        { name: "YYY", fallback: true, expectedVal: true, set: true, setVal: "TRUE" },
        { name: "YYY", fallback: true, expectedVal: true, set: true, setVal: "yes" },
        { name: "YYY", fallback: true, expectedVal: true, set: true, setVal: "Yes" },
        { name: "YYY", fallback: true, expectedVal: true, set: true, setVal: "YES" },
        { name: "YYY", fallback: true, expectedVal: false, set: true, setVal: "0" },
        { name: "YYY", fallback: true, expectedVal: false, set: true, setVal: "false" },
        { name: "YYY", fallback: true, expectedVal: false, set: true, setVal: "False" },
        { name: "YYY", fallback: true, expectedVal: false, set: true, setVal: "FALSE" },
        { name: "YYY", fallback: true, expectedVal: false, set: true, setVal: "no" },
        { name: "YYY", fallback: true, expectedVal: false, set: true, setVal: "No" },
        { name: "YYY", fallback: true, expectedVal: false, set: true, setVal: "NO" },
        { name: "YYY", fallback: true, expectedVal: false, set: true, setVal: "UNKNOWN" },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)

        if test.set {
            os.Setenv(test.name, test.setVal)
        }

        val := GetEnvBool(test.name, test.fallback)

        if val != test.expectedVal {
            t.Errorf("%sExpected '%s', found '%s'.", testPrefix, test.expectedVal, val)
        }

        if test.set {
            os.Unsetenv(test.name)
        }
    }
}

func TestRun(t *testing.T) {
    tests := []struct {
        cmd []string
        rc int
        stdout string
        stderr string
    }{
        { cmd: []string{"find", "/", "-maxdepth", "1", "-name", "tmp"}, rc: 0, stdout: "/tmp", stderr: "", },
        { cmd: []string{"cd", "/"}, rc: 1, stdout: "", stderr: "exec: \"cd\": executable file not found in $PATH", },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)
        rc, stdout, stderr := Run(test.cmd)

        if test.rc != rc || test.stdout != stdout || test.stderr != stderr {
            t.Errorf(
                "%sExpected (RC='%d'; STDOUT='%s'; STDERR='%s'), found (RC='%d'; STDOUT='%s'; STDERR='%s').",
                testPrefix,
                test.rc,
                test.stdout,
                test.stderr,
                rc,
                stdout,
                stderr)
        }
    }
}
