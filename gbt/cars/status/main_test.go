package status

import (
    "testing"
)

func TestInit(t *testing.T) {
    car := Car{}

    car.Init()

    if car.Wrap != false {
        t.Errorf("Expected %s = %t, found %t.", "Wrap", false, car.Wrap)
    }
}
