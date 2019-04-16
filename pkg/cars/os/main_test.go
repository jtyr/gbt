package os

import (
    "os"
    "testing"
)

func TestInitDefault(t *testing.T) {
    tests := []struct {
        name string
        osRelease string
        expectedOutput string
    }{
        {
            name: "linux",
            osRelease: "/proc/1/environ",
            expectedOutput: "\uf17c",
        },
        {
            name: "unknown",
            osRelease: "/etc/os-release",
            expectedOutput: "?",
        },
        {
            name: "unknown",
            osRelease: "/etc/os-release",
            expectedOutput: "?",
        },
        {
            name: "opensuse-leap",
            osRelease: "/etc/os-release",
            expectedOutput: "\uf314",
        },
        {
            name: "opensuse-tumbleweed",
            osRelease: "/etc/os-release",
            expectedOutput: "\uf314",
        },
    }

    for i, test := range tests {
        os.Setenv("GBT_CAR_OS_NAME", test.name)
        osReleaseFile = test.osRelease
        osName = ""

        car := Car{}

        car.Init()

        if car.Model["Symbol"].Text != test.expectedOutput {
            t.Errorf("Test [%d]: Expected '%s', found '%s'.", i, test.expectedOutput, car.Model["Symbol"].Text)
        }
    }
}
