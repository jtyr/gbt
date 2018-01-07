package exectime

import (
    "os"
    "testing"
    "time"
)

func TestInit(t *testing.T) {
    os.Setenv("GBT_CAR_EXECTIME_SECS", "1515278961.234567890")

    fakedDate := time.Date(2018, time.January, 6, 23, 57, 41, 123456789, time.UTC)
    tnow = func() time.Time {
        return fakedDate
    }

    tests := []struct {
        precision string
        expectedOutput string
    }{
        {
            precision: "0",
            expectedOutput: "01:08:19",
        },
        {
            precision: "4",
            expectedOutput: "01:08:19.8889",
        },
    }

    for i, test := range tests {
        os.Setenv("GBT_CAR_EXECTIME_PRECISION", test.precision)

        car := Car{}
        car.Init()

        if car.Model["Time"].Text != test.expectedOutput {
            t.Errorf("Test [%d]: Expected '%s', found '%s'.", i, test.expectedOutput, car.Model["Time"].Text)
        }
    }
}
