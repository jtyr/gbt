package custom

import (
    "os"
    "testing"
)

func TestInit(t *testing.T) {
    tests := []struct {
        cmdText string
        cmdDisplay string
        expectedOutput string
        expectedDisplay bool
    }{
        {
            cmdText: "echo 70",
            cmdDisplay: "",
            expectedOutput: "70",
            expectedDisplay: true,
        },
        {
            cmdText: "echo 70",
            cmdDisplay: "echo YES",
            expectedOutput: "70",
            expectedDisplay: true,
        },
        {
            cmdText: "echo 70",
            cmdDisplay: "echo NO",
            expectedOutput: "70",
            expectedDisplay: false,
        },
    }

    for i, test := range tests {
        os.Setenv("GBT_CAR_CUSTOM_TEXT_CMD", test.cmdText)
        os.Setenv("GBT_CAR_CUSTOM_DISPLAY_CMD", test.cmdDisplay)

        car := Car{}
        car.SetParamStr("name", "")
        car.Init()

        if car.Model["Text"].Text != test.expectedOutput {
            t.Errorf("Test [%d]: Expected Text to be '%s', got '%s'.", i, test.expectedOutput, car.Model["Text"].Text)
        }

        if car.Display != test.expectedDisplay {
            t.Errorf("Test [%d]: Expected Display to be '%t', got '%t'.", i, test.expectedDisplay, car.Display)
        }
    }
}
