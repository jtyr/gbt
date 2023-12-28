package exectime

import (
	"os"
	"testing"
	"time"

	ct "github.com/jtyr/gbt/pkg/core/testing"
)

func TestInit(t *testing.T) {
	ct.ResetEnv()

	os.Setenv("GBT_CAR_EXECTIME_SECS", "1515278961.987654321")

	fakedDate := time.Date(2018, time.January, 6, 23, 57, 41, 123456789, time.UTC)
	tnow = func() time.Time {
		return fakedDate
	}

	tests := []struct {
		format         string
		precision      string
		expectedOutput string
	}{
		{
			format:         "Duration",
			precision:      "0",
			expectedOutput: "1h8m19s",
		},
		{
			format:         "Duration",
			precision:      "7",
			expectedOutput: "1h8m19s135ms802µs507ns",
		},
		{
			format:         "Seconds",
			precision:      "0",
			expectedOutput: "4099",
		},
		{
			format:         "Seconds",
			precision:      "4",
			expectedOutput: "4099.1358",
		},
		{
			format:         "Time",
			precision:      "0",
			expectedOutput: "01:08:19",
		},
		{
			format:         "Time",
			precision:      "4",
			expectedOutput: "01:08:19.1358",
		},
	}

	for i, test := range tests {
		os.Setenv("GBT_CAR_EXECTIME_PRECISION", test.precision)

		car := Car{}
		car.Init()

		if car.Model[test.format].Text != test.expectedOutput {
			t.Errorf("Test [%d]: Expected '%s', found '%s'.", i, test.expectedOutput, car.Model[test.format].Text)
		}
	}
}
