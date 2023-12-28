package dir

import (
	"os"
	"testing"

	ct "github.com/jtyr/gbt/pkg/core/testing"
)

func TestInit(t *testing.T) {
	ct.ResetEnv()

	tests := []struct {
		pwd            string
		expectedOutput string
		sep            string
		depth          string
		nonCurLen      string
	}{
		{
			pwd:            "/",
			expectedOutput: "/",
		},
		{
			pwd:            "/bin",
			expectedOutput: "bin",
		},
		{
			pwd:            "C:\\",
			expectedOutput: "C:",
			sep:            "\\",
		},
		{
			pwd:            "C:\\tmp",
			expectedOutput: "tmp",
			sep:            "\\",
		},
		{
			pwd:            "//",
			expectedOutput: "//",
		},
		{
			pwd:            os.Getenv("HOME"),
			expectedOutput: "~",
		},
		{
			pwd:            "/usr",
			expectedOutput: "",
			depth:          "0",
		},
		{
			pwd:            "/usr",
			expectedOutput: "/usr",
			depth:          "999",
		},
		{
			pwd:            "C:\\Windows\\system32",
			expectedOutput: "C:\\Windows\\system32",
			sep:            "\\",
			depth:          "999",
		},
		{
			pwd:            "/usr/share/ssl",
			expectedOutput: "share/ssl",
			depth:          "2",
		},
		{
			pwd:            "/usr/share/ssl",
			expectedOutput: "s/ssl",
			depth:          "2",
			nonCurLen:      "1",
		},
		{
			pwd:            "C:\\Windows\\system32",
			expectedOutput: "W\\system32",
			sep:            "\\",
			depth:          "2",
			nonCurLen:      "1",
		},
		{
			pwd:            "C:\\Windows\\system32",
			expectedOutput: "C:\\W\\system32",
			sep:            "\\",
			depth:          "999",
			nonCurLen:      "1",
		},
	}

	for i, test := range tests {
		car := Car{}

		os.Setenv("PWD", test.pwd)

		if test.depth != "" {
			os.Setenv("GBT_CAR_DIR_DEPTH", test.depth)
		} else {
			os.Unsetenv("GBT_CAR_DIR_DEPTH")
		}

		if test.nonCurLen != "" {
			os.Setenv("GBT_CAR_DIR_NONCURLEN", test.nonCurLen)
		} else {
			os.Unsetenv("GBT_CAR_DIR_NONCURLEN")
		}

		if test.sep != "" {
			osSep = test.sep
		} else {
			osSep = string(os.PathSeparator)
		}

		car.Init()

		if car.Model["Dir"].Text != test.expectedOutput {
			t.Errorf("Test [%d]: Expected '%s', found '%s'.", i, test.expectedOutput, car.Model["Dir"].Text)
		}
	}
}
