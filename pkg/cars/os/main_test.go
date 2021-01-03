package os

import (
    "io/ioutil"
    "log"
    "os"
    "testing"

    "github.com/jtyr/gbt/pkg/core/utils"
)

func TestInitDefault(t *testing.T) {
    utils.ResetEnv()

    tests := []struct {
        goos string
        name string
        osReleaseFile string
        expectedOutput string
    }{
        {
            goos: "linux",
            name: "linux",
            osReleaseFile: "/proc/1/environ",
            expectedOutput: "\uf17c",
        },
        {
            goos: "linux",
            name: "unknown",
            expectedOutput: "\uf17c",
        },
        {
            goos: "linux",
            name: "arch",
            expectedOutput: "\uf303",
        },
        {
            goos: "unknown",
            name: "unknown",
            osReleaseFile: "/etc/os-release.unknown",
            expectedOutput: "?",
        },
    }

    for i, test := range tests {
        osName = ""
        goos = test.goos

        if test.osReleaseFile == "" {
            content := []byte("ID=" + test.name)
            tmpfile, err := ioutil.TempFile("", "test")

            if err != nil {
                log.Fatal(err)
            }

            osReleaseFile = tmpfile.Name()

            defer os.Remove(tmpfile.Name())

            if _, err := tmpfile.Write(content); err != nil {
                log.Fatal(err)
            }

            if err := tmpfile.Close(); err != nil {
                log.Fatal(err)
            }
        } else {
            osReleaseFile = test.osReleaseFile
        }

        car := Car{}

        car.Init()

        if car.Model["Symbol"].Text != test.expectedOutput {
            t.Errorf("Test [%d]: Expected '%s', found '%s'.", i, test.expectedOutput, car.Model["Symbol"].Text)
        }
    }
}
