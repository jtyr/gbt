package hostname

import (
    "fmt"
    "os"
    "os/user"
    "strings"

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
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "dark_gray")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "252")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultUserhostBg := defaultRootBg
    defaultUserhostFg := defaultRootFg
    defaultUserhostFm := defaultRootFm
    defaultUserBg := defaultRootBg
    defaultUserFg := defaultRootFg
    defaultUserFm := defaultRootFm
    defaultAdminBg := defaultRootBg
    defaultAdminFg := defaultRootFg
    defaultAdminFm := defaultRootFm
    defaultHostBg := defaultRootBg
    defaultHostFg := defaultRootFg
    defaultHostFm := defaultRootFm

    curUser, _ := user.Current()
    hostname, _ := os.Hostname()
    hostname = strings.Split(hostname, ".")[0]

    uaFormat := "{{ User }}"

    if curUser.Uid == adminUID {
        uaFormat = "{{ Admin }}"
    }

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_HOSTNAME_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_HOSTNAME_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_HOSTNAME_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_HOSTNAME_FORMAT", " {{ UserHost }} "),
        },
        "UserHost": {
            Bg: utils.GetEnv(
                "GBT_CAR_HOSTNAME_USERHOST_BG", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_BG", defaultUserhostBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_HOSTNAME_USERHOST_FG", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_FG", defaultUserhostFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_HOSTNAME_USERHOST_FM", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_FM", defaultUserhostFm)),
            Text: utils.GetEnv(
                "GBT_CAR_HOSTNAME_USERHOST_FORMAT", fmt.Sprintf("%s@{{ Host }}", uaFormat)),
        },
        "Admin": {
            Bg: utils.GetEnv(
                "GBT_CAR_HOSTNAME_ADMIN_BG", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_BG", defaultAdminBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_HOSTNAME_ADMIN_FG", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_FG", defaultAdminFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_HOSTNAME_ADMIN_FM", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_FM", defaultAdminFm)),
            Text: utils.GetEnv("GBT_CAR_HOSTNAME_ADMIN_TEXT", curUser.Username),
        },
        "User": {
            Bg: utils.GetEnv(
                "GBT_CAR_HOSTNAME_USER_BG", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_BG", defaultUserBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_HOSTNAME_USER_FG", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_FG", defaultUserFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_HOSTNAME_USER_FM", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_FM", defaultUserFm)),
            Text: utils.GetEnv("GBT_CAR_HOSTNAME_USER_TEXT", curUser.Username),
        },
        "Host": {
            Bg: utils.GetEnv(
                "GBT_CAR_HOSTNAME_HOST_BG", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_BG", defaultHostBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_HOSTNAME_HOST_FG", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_FG", defaultHostFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_HOSTNAME_HOST_FM", utils.GetEnv(
                    "GBT_CAR_HOSTNAME_FM", defaultHostFm)),
            Text: utils.GetEnv("GBT_CAR_HOSTNAME_HOST_TEXT", hostname),
        },
    }

    c.Display = utils.GetEnvBool("GBT_CAR_HOSTNAME_DISPLAY", true)
    c.Wrap = utils.GetEnvBool("GBT_CAR_HOSTNAME_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_HOSTNAME_SEP", "\000")
}
