package git

import (
    "testing"
)

func compareRemote(display bool, ahead bool) bool {
    return true
}

func TestInit(t *testing.T) {
    car := Car{}

    car.Init()

    if car.Wrap != false {
        t.Errorf("Expected %s = %x, found %x.", "Wrap", false, car.Wrap)
    }
}
