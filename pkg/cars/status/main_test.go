package status

import (
    "testing"
)

func TestInit(t *testing.T) {
    tests := []struct {
        args string
        expectedDisplay bool
    }{
        {
            args: "0",
            expectedDisplay: false,
        },
        {
            args: "1",
            expectedDisplay: true,
        },
    }

    for i, test := range tests {
        car := Car{}

        car.SetParamStr("args", test.args)
        car.Init()

        if car.Display != test.expectedDisplay {
            t.Errorf("Test [%d]: Expected %t, found %t.", i, test.expectedDisplay, car.Display)
        }
    }
}
