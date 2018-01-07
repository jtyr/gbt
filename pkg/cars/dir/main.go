package dir

import (
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

// getDir returns the directory name.
func getDir() (ret string) {
    wd, _ := os.Getwd()
    sep := string(os.PathSeparator)

    pwd := utils.GetEnv("PWD", wd)
    dirSep := utils.GetEnv("GBT_CAR_DIR_DIRSEP", sep)
    userDirSign := utils.GetEnv("GBT_CAR_DIR_HOMESIGN", "~")

    if userDirSign != "" {
        usr, _ := user.Current()
        pwd = strings.Replace(pwd, usr.HomeDir, userDirSign, 1)
    }

    dirs := strings.Split(pwd, sep)
    dirsLen := len(dirs)
    depth := utils.GetEnvInt("GBT_CAR_DIR_DEPTH", 1)

    if depth > dirsLen {
        depth = dirsLen
    }

    if pwd == sep {
        ret = dirSep
    } else if pwd == "~" {
        ret = pwd
    } else {
        ret = strings.Join(dirs[(dirsLen - depth):], dirSep)
    }

    return
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "blue")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "light_gray")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultDirBg := defaultRootBg
    defaultDirFg := defaultRootFg
    defaultDirFm := defaultRootFm

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_DIR_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_DIR_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_DIR_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_DIR_FORMAT", " {{ Dir }} "),
        },
        "Dir": {
            Bg: utils.GetEnv(
                "GBT_CAR_DIR_DIR_BG", utils.GetEnv(
                    "GBT_CAR_DIR_BG", defaultDirBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_DIR_DIR_FG", utils.GetEnv(
                    "GBT_CAR_DIR_FG", defaultDirFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_DIR_DIR_FM", utils.GetEnv(
                    "GBT_CAR_DIR_FM", defaultDirFm)),
            Text: utils.GetEnv("GBT_CAR_DIR_DIR_TEXT", getDir()),
        },
    }

    c.Display = utils.GetEnvBool("GBT_CAR_DIR_DISPLAY", true)
    c.Wrap = utils.GetEnvBool("GBT_CAR_DIR_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_DIR_SEP", "\000")
}
