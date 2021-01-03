package pyvirtenv

import (
    "os"
    "testing"

    "github.com/jtyr/gbt/pkg/core/utils"
)

func TestInit(t *testing.T) {
    utils.ResetEnv()

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
