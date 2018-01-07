package status

import (
    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

// Checks for the return code.
func isOk(c *Car) (ret bool) {
    _, argsExist := c.Params["args"]

    if ! argsExist || c.Params["args"] == "0" {
        ret = true
    } else {
        ret = false
    }

    return
}

// Init initializes the car.
func (c *Car) Init() {
    defaultErrorBg := "red"
    defaultErrorFg := "light_gray"
    defaultErrorFm := "none"
    defaultOkBg := "green"
    defaultOkFg := "light_gray"
    defaultOkFm := "none"
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", defaultErrorBg)
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", defaultErrorFg)
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", defaultErrorFm)
    defaultSymbolBg := defaultRootBg
    defaultSymbolFg := defaultRootFg
    defaultSymbolFm := defaultRootFm
    defaultCodeBg := defaultRootBg
    defaultCodeFg := defaultRootFg
    defaultCodeFm := defaultRootFm

    defaultSymbolFormat := "{{ Error }}"
    defaultCodeText := "?"

    if val, ok := c.Params["args"]; ok {
        defaultCodeText = val.(string)
    }

    if isOk(c) {
        defaultRootBg = utils.GetEnv("GBT_CAR_BG", defaultOkBg)
        defaultRootFg = utils.GetEnv("GBT_CAR_FG", defaultOkFg)
        defaultRootFm = utils.GetEnv("GBT_CAR_FM", defaultOkFm)
        defaultSymbolFormat = "{{ Ok }}"
    }

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_STATUS_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_STATUS_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_STATUS_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_STATUS_FORMAT", " {{ Symbol }} "),
        },
        "Symbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_STATUS_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_STATUS_BG", defaultSymbolBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_STATUS_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_STATUS_FG", defaultSymbolFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_STATUS_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_STATUS_FM", defaultSymbolFm)),
            Text: utils.GetEnv(
                "GBT_CAR_STATUS_SYMBOL_FORMAT", defaultSymbolFormat),
        },
        "Code": {
            Bg: utils.GetEnv(
                "GBT_CAR_STATUS_CODE_BG", utils.GetEnv(
                    "GBT_CAR_STATUS_BG", defaultCodeBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_STATUS_CODE_FG", utils.GetEnv(
                    "GBT_CAR_STATUS_FG", defaultCodeFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_STATUS_CODE_FM", utils.GetEnv(
                    "GBT_CAR_STATUS_FM", defaultCodeFm)),
            Text: utils.GetEnv(
                "GBT_CAR_STATUS_CODE_TEXT", defaultCodeText),
        },
        "Error": {
            Bg: utils.GetEnv(
                "GBT_CAR_STATUS_ERROR_BG", utils.GetEnv(
                    "GBT_CAR_STATUS_SYMBOL_BG", utils.GetEnv(
                        "GBT_CAR_STATUS_BG", defaultErrorBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_STATUS_ERROR_FG", utils.GetEnv(
                    "GBT_CAR_STATUS_SYMBOL_FG", utils.GetEnv(
                        "GBT_CAR_STATUS_FG", defaultErrorFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_STATUS_ERROR_FM", utils.GetEnv(
                    "GBT_CAR_STATUS_SYMBOL_FM", utils.GetEnv(
                        "GBT_CAR_STATUS_FM", defaultErrorFm))),
            Text: utils.GetEnv("GBT_CAR_STATUS_ERROR_TEXT", "✘"),
        },
        "Ok": {
            Bg: utils.GetEnv(
                "GBT_CAR_STATUS_OK_BG", utils.GetEnv(
                    "GBT_CAR_STATUS_SYMBOL_BG", utils.GetEnv(
                        "GBT_CAR_STATUS_BG", defaultOkBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_STATUS_OK_FG", utils.GetEnv(
                    "GBT_CAR_STATUS_SYMBOL_FG", utils.GetEnv(
                        "GBT_CAR_STATUS_FG", defaultOkFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_STATUS_OK_FM", utils.GetEnv(
                    "GBT_CAR_STATUS_SYMBOL_FM", utils.GetEnv(
                        "GBT_CAR_STATUS_FM", defaultOkFm))),
            Text: utils.GetEnv("GBT_CAR_STATUS_OK_TEXT", "✔"),
        },
    }

    if isOk(c) {
        c.Display = utils.GetEnvBool("GBT_CAR_STATUS_DISPLAY", false)
    } else {
        c.Display = utils.GetEnvBool("GBT_CAR_STATUS_DISPLAY", true)
    }

    c.Wrap = utils.GetEnvBool("GBT_CAR_STATUS_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_STATUS_SEP", "\000")
}
