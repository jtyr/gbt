package dir

import (
    "os"
    "testing"
)

func TestInit(t *testing.T) {
    tests := []struct {
        pwd string
        expectedOutput string
        depth string
        nonCurLen string
    }{
        {
            pwd: "/",
            expectedOutput: "/",
        },
        {
            pwd: "//",
            expectedOutput: "//",
        },
        {
            pwd: os.Getenv("HOME"),
            expectedOutput: "~",
        },
        {
            pwd: "/usr",
            expectedOutput: "/usr",
            depth: "999",
        },
        {
            pwd: "/usr/share/ssl",
            expectedOutput: "share/ssl",
            depth: "2",
        },
        {
            pwd: "/usr/share/ssl",
            expectedOutput: "s/ssl",
            depth: "2",
            nonCurLen: "1",
        },
    }

    for i, test := range tests {
        car := Car{}

        os.Setenv("PWD", test.pwd)

        if len(test.depth) > 0 {
            os.Setenv("GBT_CAR_DIR_DEPTH", test.depth)
        }

        if len(test.nonCurLen) > 0 {
            os.Setenv("GBT_CAR_DIR_NONCURLEN", test.nonCurLen)
        }

        car.Init()

        if car.Model["Dir"].Text != test.expectedOutput {
            t.Errorf("Test [%d]: Expected '%s', found '%s'.", i, test.expectedOutput, car.Model["Dir"].Text)
        }
    }
}
