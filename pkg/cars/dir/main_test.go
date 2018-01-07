package dir

import (
    "os"
    "testing"
)

func TestInit(t *testing.T) {
    os.Setenv("GBT_CAR_DIR_DEPTH", "2")

    tests := []struct {
        pwd string
        expectedOutput string
    }{
        {
            pwd: "/",
            expectedOutput: "/",
        },
        {
            pwd: os.Getenv("HOME"),
            expectedOutput: "~",
        },
        {
            pwd: "/usr",
            expectedOutput: "/usr",
        },
        {
            pwd: "/usr/share/ssl",
            expectedOutput: "share/ssl",
        },
    }

    for i, test := range tests {
        car := Car{}

        os.Setenv("PWD", test.pwd)
        car.Init()

        if car.Model["Dir"].Text != test.expectedOutput {
            t.Errorf("Test [%d]: Expected '%s', found '%s'.", i, test.expectedOutput, car.Model["Dir"].Text)
        }
    }
}
