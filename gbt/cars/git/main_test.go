package git

import (
    "testing"
)

func TestInit(t *testing.T) {
    car := Car{}

    car.Display = true
    car.Init()

    if car.Wrap != false {
        t.Errorf("Expected %s = %t, found %t.", "Wrap", false, car.Wrap)
    }
}
