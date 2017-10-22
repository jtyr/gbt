package custom

import (
    "github.com/jtyr/gbt/gbt/core/car"
    "github.com/jtyr/gbt/gbt/core/utils"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "yellow")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "default")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultTextBg := defaultRootBg
    defaultTextFg := defaultRootFg
    defaultTextFm := defaultRootFm

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_CUSTOM_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_CUSTOM_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_CUSTOM_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_CUSTOM_FORMAT", " {{ Text }} "),
        },
        "Text": {
            Bg: utils.GetEnv(
                "GBT_CAR_CUSTOM_DIR_BG", utils.GetEnv(
                    "GBT_CAR_CUSTOM_BG", defaultTextBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_CUSTOM_DIR_FG", utils.GetEnv(
                    "GBT_CAR_CUSTOM_FG", defaultTextFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_CUSTOM_DIR_FM", utils.GetEnv(
                    "GBT_CAR_CUSTOM_FM", defaultTextFm)),
            Text: utils.GetEnv("GBT_CAR_CUSTOM_DIR_TEXT", "?"),
        },
    }

    c.Display = utils.GetEnvBool("GBT_CAR_CUSTOM_DISPLAY", true)
    c.Wrap = utils.GetEnvBool("GBT_CAR_CUSTOM_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_CUSTOM_SEP", "\000")
}
