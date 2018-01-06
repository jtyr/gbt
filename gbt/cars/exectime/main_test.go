package exectime

import (
    "os"
    "testing"
)

func TestInit(t *testing.T) {
    car := Car{}

    os.Setenv("GBT_CAR_EXECTIME_PRECISION", "3")

    car.Init()

    if car.Wrap != false {
        t.Errorf("Expected %s = %t, found %t.", "Wrap", false, car.Wrap)
    }
}
