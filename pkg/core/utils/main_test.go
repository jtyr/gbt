package utils

import (
    "fmt"
    "os"
    "testing"
)

func getTestPrefix(testN int) string {
    return fmt.Sprintf("Test [%d]: ", testN)
}

func TestIsTrue(t *testing.T) {
    tests := []struct {
        input string
        expectedOutput bool
    }{
        { input: "true", expectedOutput: true  },
        { input: "True", expectedOutput: true  },
        { input: "TRUE", expectedOutput: true  },
        { input: "yes",  expectedOutput: true  },
        { input: "Yes",  expectedOutput: true  },
        { input: "YES",  expectedOutput: true  },
        { input: "1",    expectedOutput: true  },
        { input: "No",   expectedOutput: false },
        { input: "0",    expectedOutput: false },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)

        output := IsTrue(test.input)

        if output != test.expectedOutput {
            t.Errorf("%sExpected '%t', found '%t'.", testPrefix, test.expectedOutput, output)
        }
    }
}

func TestGetEnv(t *testing.T) {
    tests := []struct {
        name string
        fallback string
        expectedVal string
        set bool
    }{
        { name: "XXX", fallback: "", expectedVal: "",     set: false, },
        { name: "YYY", fallback: "", expectedVal: "test", set: true,  },
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
        { name: "XXX", fallback: false, expectedVal: false, set: false, setVal: ""       },
        { name: "YYY", fallback: true,  expectedVal: true,  set: true, setVal: "1"       },
        { name: "YYY", fallback: true,  expectedVal: true,  set: true, setVal: "true"    },
        { name: "YYY", fallback: true,  expectedVal: true,  set: true, setVal: "True"    },
        { name: "YYY", fallback: true,  expectedVal: true,  set: true, setVal: "TRUE"    },
        { name: "YYY", fallback: true,  expectedVal: true,  set: true, setVal: "yes"     },
        { name: "YYY", fallback: true,  expectedVal: true,  set: true, setVal: "Yes"     },
        { name: "YYY", fallback: true,  expectedVal: true,  set: true, setVal: "YES"     },
        { name: "YYY", fallback: true,  expectedVal: false, set: true, setVal: "0"       },
        { name: "YYY", fallback: true,  expectedVal: false, set: true, setVal: "false"   },
        { name: "YYY", fallback: true,  expectedVal: false, set: true, setVal: "False"   },
        { name: "YYY", fallback: true,  expectedVal: false, set: true, setVal: "FALSE"   },
        { name: "YYY", fallback: true,  expectedVal: false, set: true, setVal: "no"      },
        { name: "YYY", fallback: true,  expectedVal: false, set: true, setVal: "No"      },
        { name: "YYY", fallback: true,  expectedVal: false, set: true, setVal: "NO"      },
        { name: "YYY", fallback: true,  expectedVal: false, set: true, setVal: "UNKNOWN" },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)

        if test.set {
            os.Setenv(test.name, test.setVal)
        }

        val := GetEnvBool(test.name, test.fallback)

        if val != test.expectedVal {
            t.Errorf("%sExpected '%t', found '%t'.", testPrefix, test.expectedVal, val)
        }

        if test.set {
            os.Unsetenv(test.name)
        }
    }
}

func TestGetEnvInt(t *testing.T) {
    tests := []struct {
        name string
        fallback int
        expectedVal int
        set bool
        setVal string
    }{
        { name: "XXX", fallback: 1, expectedVal: 1, set: false, setVal: ""  },
        { name: "YYY", fallback: 1, expectedVal: 2, set: true,  setVal: "2" },
        { name: "YYY", fallback: 1, expectedVal: 1, set: true,  setVal: "t" },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)

        if test.set {
            os.Setenv(test.name, test.setVal)
        }

        val := GetEnvInt(test.name, test.fallback)

        if val != test.expectedVal {
            t.Errorf("%sExpected '%d', found '%d'.", testPrefix, test.expectedVal, val)
        }

        if test.set {
            os.Unsetenv(test.name)
        }
    }
}

func TestGetEnvFloat(t *testing.T) {
    tests := []struct {
        name string
        fallback float64
        expectedVal float64
        set bool
        setVal string
    }{
        { name: "XXX", fallback: 1.0, expectedVal: 1.0, set: false, setVal: ""    },
        { name: "YYY", fallback: 1.0, expectedVal: 2.0, set: true,  setVal: "2.0" },
        { name: "YYY", fallback: 1.0, expectedVal: 1.0, set: true,  setVal: "x.y" },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)

        if test.set {
            os.Setenv(test.name, test.setVal)
        }

        val := GetEnvFloat(test.name, test.fallback)

        if val != test.expectedVal {
            t.Errorf("%sExpected '%f', found '%f'.", testPrefix, test.expectedVal, val)
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
        { cmd: []string{"curl", "-s", "-S", "http://localhost:12345"},  rc: 7, stdout: "",     stderr: "curl: (7) Failed to connect to localhost port 12345: Connection refused", },
        { cmd: []string{"unknown_command"},                             rc: 1, stdout: "",     stderr: "exec: \"unknown_command\": executable file not found in $PATH", },
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
