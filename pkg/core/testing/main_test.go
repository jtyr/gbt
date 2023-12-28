package testing

import (
	"os"
	"testing"
)

type env struct {
	name  string
	value string
}

func TestInit(t *testing.T) {
	tests := []struct {
		set      []env
		ignore   []string
		expected []string
	}{
		{
			expected: []string{"PATH", "HOME"},
		},
		{
			set: []env{
				{
					name:  "GBT_ENV_TEST",
					value: "test",
				},
			},
			expected: []string{"PATH", "HOME"},
		},
		{
			set: []env{
				{
					name:  "GBT_ENV_TEST",
					value: "test",
				},
			},
			ignore:   []string{"PATH", "GBT_ENV_TEST", "HOME"},
			expected: []string{"PATH", "GBT_ENV_TEST", "HOME"},
		},
	}

	for i, test := range tests {
		if len(test.set) > 0 {
			for _, e := range test.set {
				os.Setenv(e.name, e.value)
			}
		}

		if len(test.ignore) > 0 {
			ResetEnvIgnore = test.ignore
		}

		ResetEnv()
		keys := getEnvKeys()

		if len(keys) != len(test.expected) {
			t.Errorf("Test [%d]: Expected '%v', found '%v'.", i, test.expected, keys)
		}
	}
}
