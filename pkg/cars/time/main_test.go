package ttime

import (
    "testing"
    "time"

    ccar "github.com/jtyr/gbt/pkg/core/car"
    ct "github.com/jtyr/gbt/pkg/core/testing"
)

func TestInit(t *testing.T) {
    ct.ResetEnv()

    fakedDate := time.Date(2018, time.January, 6, 23, 57, 41, 0, time.UTC)
    tnow = func() time.Time {
        return fakedDate
    }

    ccar.Shell = "plain"

    tests := []struct {
        field string
        expectedOutput string
    }{
        {
            field: "Date",
            expectedOutput: "Sat 06 Jan",
        },
        {
            field: "Time",
            expectedOutput: "23:57:41",
        },
    }

    for i, test := range tests {
        car := Car{}
        car.Init()

        val := car.Model[test.field].Text

        if val != test.expectedOutput {
            t.Errorf("Test [%d]: Expected '%s', found '%s'.", i, test.expectedOutput, val)
        }
    }
}
