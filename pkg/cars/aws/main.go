package aws

import (
    "fmt"
    "os"
    "os/user"
    "path/filepath"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"

    "gopkg.in/go-ini/ini.v1"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

// To be able to fake the user's home directory
var usr, _ = user.Current()

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "180;85;10")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "white")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultSep := "\000"

    profile := utils.GetEnv("AWS_PROFILE", "default")
    region := os.Getenv("AWS_DEFAULT_REGION")

    c.Display = utils.GetEnvBool("GBT_CAR_AWS_DISPLAY", true)

    if c.Display {
        configFile := filepath.Join(usr.HomeDir, ".aws", "config")
        cfg, err := ini.Load(configFile)

        if err == nil {
            if region == "" {
                profileSection := profile

                if profile != "default" {
                    profileSection = fmt.Sprintf("profile %s", profile)
                }

                region = cfg.Section(profileSection).Key("region").String()
            }
        }
    }

    c.Model = map[string]car.ModelElement{
        "root": {
            Bg:   utils.GetEnv("GBT_CAR_AWS_BG", defaultRootBg),
            Fg:   utils.GetEnv("GBT_CAR_AWS_FG", defaultRootFg),
            Fm:   utils.GetEnv("GBT_CAR_AWS_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_AWS_FORMAT", " {{ Icon }} {{ Profile }} "),
        },
        "Icon": {
            Bg: utils.GetEnv(
                "GBT_CAR_AWS_ICON_BG", utils.GetEnv(
                    "GBT_CAR_AWS_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_AWS_ICON_FG", utils.GetEnv(
                    "GBT_CAR_AWS_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_AWS_ICON_FM", utils.GetEnv(
                    "GBT_CAR_AWS_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_AWS_ICON_TEXT", "\uf52d"),
        },
        "Profile": {
            Bg: utils.GetEnv(
                "GBT_CAR_AWS_CONFIG_BG", utils.GetEnv(
                    "GBT_CAR_AWS_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_AWS_CONFIG_FG", utils.GetEnv(
                    "GBT_CAR_AWS_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_AWS_CONFIG_FM", utils.GetEnv(
                    "GBT_CAR_AWS_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_AWS_CONFIG_TEXT", profile),
        },
        "Region": {
            Bg: utils.GetEnv(
                "GBT_CAR_AWS_ACCOUNT_BG", utils.GetEnv(
                    "GBT_CAR_AWS_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_AWS_ACCOUNT_FG", utils.GetEnv(
                    "GBT_CAR_AWS_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_AWS_ACCOUNT_FM", utils.GetEnv(
                    "GBT_CAR_AWS_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_AWS_ACCOUNT_TEXT", region),
        },
        "Sep": {
            Bg: utils.GetEnv(
                "GBT_CAR_AWS_SEP_BG", utils.GetEnv(
                    "GBT_SEPARATOR_BG", defaultSep)),
            Fg: utils.GetEnv(
                "GBT_CAR_AWS_SEP_FG", utils.GetEnv(
                    "GBT_SEPARATOR_FG", defaultSep)),
            Fm: utils.GetEnv(
                "GBT_CAR_AWS_SEP_FM", utils.GetEnv(
                    "GBT_SEPARATOR_FM", defaultSep)),
            Text: utils.GetEnv(
                "GBT_CAR_AWS_SEP", utils.GetEnv(
                    "GBT_CAR_AWS_SEP_TEXT", utils.GetEnv(
                        "GBT_SEPARATOR", defaultSep))),
        },
    }

    c.Wrap = utils.GetEnvBool("GBT_CAR_AWS_WRAP", false)
}
