package custom

import (
    "os"
    "testing"
)

func TestInit(t *testing.T) {
    os.Setenv("GBT_CAR_CUSTOM_TEXT_CMD", "ls /tmp")
    os.Setenv("GBT_CAR_CUSTOM_DISPLAY_CMD", "ls /tmp")

    car := Car{}

    car.SetParamStr("name", "")
    car.Init()

    if car.Wrap != false {
        t.Errorf("Expected %s = %t, found %t.", "Wrap", false, car.Wrap)
    }
}
