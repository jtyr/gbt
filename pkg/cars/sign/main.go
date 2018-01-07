package sign

import (
    "os/user"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

// Admin UID.
var adminUID = "0"

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "default")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "default")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultSymbolBg := defaultRootBg
    defaultSymbolFg := "green"
    defaultSymbolFm := "bold"
    defaultUserBg := defaultRootBg
    defaultUserFm := defaultSymbolFm
    defaultAdminBg := defaultRootBg
    defaultAdminFm := defaultSymbolFm

    symbolFormat := "{{ User }}"
    curUser, _ := user.Current()

    if curUser.Uid == adminUID {
        symbolFormat = "{{ Admin }}"
    }

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_SIGN_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_SIGN_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_SIGN_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_SIGN_FORMAT", " {{ Symbol }} "),
        },
        "Symbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_SIGN_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_SIGN_BG", defaultSymbolBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_SIGN_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_SIGN_FG", defaultSymbolFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_SIGN_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_SIGN_FM", defaultSymbolFm)),
            Text: utils.GetEnv("GBT_CAR_SIGN_SYMBOL_FORMAT", symbolFormat),
        },
        "User": {
            Bg: utils.GetEnv(
                "GBT_CAR_SIGN_USER_BG", utils.GetEnv(
                    "GBT_CAR_SIGN_BG", defaultUserBg)),
            Fg: utils.GetEnv("GBT_CAR_SIGN_USER_FG", "light_green"),
            Fm: utils.GetEnv(
                "GBT_CAR_SIGN_USER_FM", utils.GetEnv(
                    "GBT_CAR_SIGN_FM", defaultUserFm)),
            Text: utils.GetEnv("GBT_CAR_SIGN_USER_TEXT", "$"),
        },
        "Admin": {
            Bg: utils.GetEnv(
                "GBT_CAR_SIGN_ADMIN_BG", utils.GetEnv(
                    "GBT_CAR_SIGN_BG", defaultAdminBg)),
            Fg: utils.GetEnv("GBT_CAR_SIGN_ADMIN_FG", "red"),
            Fm: utils.GetEnv(
                "GBT_CAR_SIGN_ADMIN_FM", utils.GetEnv(
                    "GBT_CAR_SIGN_FM", defaultAdminFm)),
            Text: utils.GetEnv("GBT_CAR_SIGN_ADMIN_TEXT", "#"),
        },
    }

    c.Display = utils.GetEnvBool("GBT_CAR_SIGN_DISPLAY", true)
    c.Wrap = utils.GetEnvBool("GBT_CAR_SIGN_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_SIGN_SEP", "\000")
}
