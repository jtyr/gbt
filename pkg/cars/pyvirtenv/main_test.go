package pyvirtenv

import (
    "os"
    "testing"

    ct "github.com/jtyr/gbt/pkg/core/testing"
)

func TestInit(t *testing.T) {
    ct.ResetEnv()

    tests := []struct {
        virtenv string
        expectedDisplay bool
    }{
        {
            virtenv: "",
            expectedDisplay: false,
        },
        {
            virtenv: "test",
            expectedDisplay: true,
        },
    }

    for i, test := range tests {
        os.Setenv("VIRTUAL_ENV", test.virtenv)

        car := Car{}

        car.Init()

        if car.Display != test.expectedDisplay {
            t.Errorf("Test [%d]: Expected %t, found %t.", i, test.expectedDisplay, car.Display)
        }
    }
}
