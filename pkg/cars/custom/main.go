package custom

import (
    "fmt"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"
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

    prefix := fmt.Sprintf("GBT_CAR_CUSTOM%s", c.Params["name"].(string))
    defaultTextText := "?"
    defaultTextCmd := utils.GetEnv(fmt.Sprintf("%s_TEXT_CMD", prefix), "")
    defaultDisplayCmd := utils.GetEnv(fmt.Sprintf("%s_DISPLAY_CMD", prefix), "")
    defaultDisplay := true

    if defaultTextCmd != "" {
        _, defaultTextText, _ = utils.Run([]string{"sh", "-c", defaultTextCmd})
    }

    if defaultDisplayCmd != "" {
        _, defaultDisplayOutput, _ := utils.Run([]string{"sh", "-c", defaultDisplayCmd})

        if ! utils.IsTrue(defaultDisplayOutput) {
            defaultDisplay = false
        }
    }

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv(fmt.Sprintf("%s_BG", prefix), defaultRootBg),
            Fg: utils.GetEnv(fmt.Sprintf("%s_FG", prefix), defaultRootFg),
            Fm: utils.GetEnv(fmt.Sprintf("%s_FM", prefix), defaultRootFm),
            Text: utils.GetEnv(fmt.Sprintf("%s_FORMAT", prefix), " {{ Text }} "),
        },
        "Text": {
            Bg: utils.GetEnv(
                fmt.Sprintf("%s_TEXT_BG", prefix), utils.GetEnv(
                    fmt.Sprintf("%s_BG", prefix), defaultTextBg)),
            Fg: utils.GetEnv(
                fmt.Sprintf("%s_TEXT_FG", prefix), utils.GetEnv(
                    fmt.Sprintf("%s_FG", prefix), defaultTextFg)),
            Fm: utils.GetEnv(
                fmt.Sprintf("%s_TEXT_FM", prefix), utils.GetEnv(
                    fmt.Sprintf("%s_FM", prefix), defaultTextFm)),
            Text: utils.GetEnv(
                fmt.Sprintf("%s_TEXT_TEXT", prefix), defaultTextText),
        },
    }

    c.Display = utils.GetEnvBool(fmt.Sprintf("%s_DISPLAY", prefix), defaultDisplay)
    c.Wrap = utils.GetEnvBool(fmt.Sprintf("%s_WRAP", prefix), false)
    c.Sep = utils.GetEnv(fmt.Sprintf("%s_SEP", prefix), "\000")
}
