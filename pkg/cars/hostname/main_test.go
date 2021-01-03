package hostname

import (
    "os/user"
    "testing"

    "github.com/jtyr/gbt/pkg/core/utils"
)

func TestInit(t *testing.T) {
    utils.ResetEnv()

    curUser, _ := user.Current()

    tests := []struct {
        uid string
        expectedOutput string
    }{
        {
            uid: "12345",
            expectedOutput: "{{ User }}@{{ Host }}",
        },
        {
            uid: curUser.Uid,
            expectedOutput: "{{ Admin }}@{{ Host }}",
        },
    }

    for i, test := range tests {
        car := Car{}

        adminUID = test.uid

        car.Init()

        if car.Model["UserHost"].Text != test.expectedOutput {
            t.Errorf("Test [%d]: Expected '%s', found '%s'.", i, test.expectedOutput, car.Model["UserHost"].Text)
        }
    }
}
