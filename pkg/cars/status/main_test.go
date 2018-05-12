package status

import (
    "testing"
)

func TestInit(t *testing.T) {
    tests := []struct {
        args string
        expectedMsg string
        expectedDisplay bool
    }{
        {
            args: "0",
            expectedMsg: "",
            expectedDisplay: false,
        },
        {
            args: "1",
            expectedMsg: "FAIL",
            expectedDisplay: true,
        },
        {
            args: "126",
            expectedMsg: "NOEXEC",
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

        if car.Model["Msg"].Text != test.expectedMsg {
            t.Errorf("Test [%d]: Expected error message %s, found %s.", i, test.expectedMsg, car.Model["Msg"].Text)
        }
    }
}
