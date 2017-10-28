package dir

import (
    "os"
    "testing"
)

func TestInit(t *testing.T) {
    os.Setenv("GBT_CAR_DIR_DEPTH", "999")

    car := Car{}

    car.Init()

    if car.Wrap != false {
        t.Errorf("Expected %s = %x, found %x.", "Wrap", false, car.Wrap)
    }
}
