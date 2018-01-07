package sign

import (
    "os/user"
    "testing"
)

func TestInitUser(t *testing.T) {
    curUser, _ := user.Current()

    tests := []struct {
        uid string
        expectedOutput string
    }{
        {
            uid: "12345",
            expectedOutput: "{{ User }}",
        },
        {
            uid: curUser.Uid,
            expectedOutput: "{{ Admin }}",
        },
    }

    for i, test := range tests {
        car := Car{}

        adminUID = test.uid

        car.Init()

        if car.Model["Symbol"].Text != test.expectedOutput {
            t.Errorf("Test [%d]: Expected '%s', found '%s'.", i, test.expectedOutput, car.Model["Symbol"].Text)
        }
    }
}
