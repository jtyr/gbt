package ttime

import (
    "os"
    "testing"
    "time"

    ccar "github.com/jtyr/gbt/pkg/core/car"
)

func TestInit(t *testing.T) {
    fakedDate := time.Date(2018, time.January, 6, 23, 57, 41, 0, time.UTC)
    tnow = func() time.Time {
        return fakedDate
    }

    ccar.Shell = "plain"

    tests := []struct {
        format string
        expectedOutput string
    }{
        {
            format: "{{ DateTime }}",
            expectedOutput: "\x1b[48;5;12m\x1b[38;5;7m\x1b[48;5;12m\x1b[38;5;7m\x1b[48;5;12m\x1b[38;5;7mSat 06 Jan\x1b[48;5;12m\x1b[38;5;7m \x1b[48;5;12m\x1b[38;5;11m23:57:41\x1b[48;5;12m\x1b[38;5;7m\x1b[48;5;12m\x1b[38;5;7m",
        },
        {
            format: "{{ Date }}",
            expectedOutput: "\x1b[48;5;12m\x1b[38;5;7m\x1b[48;5;12m\x1b[38;5;7mSat 06 Jan\x1b[48;5;12m\x1b[38;5;7m",
        },
        {
            format: "{{ Time }}",
            expectedOutput: "\x1b[48;5;12m\x1b[38;5;7m\x1b[48;5;12m\x1b[38;5;11m23:57:41\x1b[48;5;12m\x1b[38;5;7m",
        },
    }

    for i, test := range tests {
        os.Setenv("GBT_CAR_TIME_FORMAT", test.format)

        car := Car{}
        car.Init()
        output := car.Format()

        if output != test.expectedOutput {
            t.Errorf("Test [%d]: Expected '%x', found '%x'.", i, test.expectedOutput, output)
        }
    }
}
