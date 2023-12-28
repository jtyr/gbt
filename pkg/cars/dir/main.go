package dir

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

// Max length of a dir name
const maxDirNameLen = 255

// To be able to mock the OS path separator
var osSep = string(os.PathSeparator)

// getDir returns the directory name.
func getDir() (ret string) {
	wd, _ := os.Getwd()
	sep := osSep

	pwd := utils.GetEnv("PWD", wd)
	dirSep := utils.GetEnv("GBT_CAR_DIR_DIRSEP", sep)
	userDirSign := utils.GetEnv("GBT_CAR_DIR_HOMESIGN", "~")
	depth := utils.GetEnvInt("GBT_CAR_DIR_DEPTH", 1)
	nonCurLen := utils.GetEnvInt("GBT_CAR_DIR_NONCURLEN", maxDirNameLen)

	if userDirSign != "" {
		usr, _ := user.Current()
		pwd = strings.Replace(pwd, usr.HomeDir, userDirSign, 1)
	}

	dirs := strings.Split(pwd, sep)
	dirsLen := len(dirs)

	if depth > dirsLen {
		depth = dirsLen
	}

	if depth > 1 && nonCurLen < maxDirNameLen {
		for i := 1; i < dirsLen-1; i++ {
			l := nonCurLen

			dirs[i] = dirs[i][:l]
		}
	} else if depth == 1 && dirsLen == 2 && len(dirs[1]) == 0 {
		ret += dirs[0]
	}

	if pwd == sep {
		ret += dirSep
	} else if pwd == fmt.Sprintf("%s%s", sep, sep) {
		ret += fmt.Sprintf("%s%s", dirSep, dirSep)
	} else if pwd == "~" {
		ret += pwd
	} else {
		ret += strings.Join(dirs[(dirsLen-depth):], dirSep)
	}

	return
}

// Init initializes the car.
func (c *Car) Init() {
	defaultRootBg := utils.GetEnv("GBT_CAR_BG", "blue")
	defaultRootFg := utils.GetEnv("GBT_CAR_FG", "light_gray")
	defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
	defaultSep := "\000"

	c.Model = map[string]car.ModelElement{
		"root": {
			Bg:   utils.GetEnv("GBT_CAR_DIR_BG", defaultRootBg),
			Fg:   utils.GetEnv("GBT_CAR_DIR_FG", defaultRootFg),
			Fm:   utils.GetEnv("GBT_CAR_DIR_FM", defaultRootFm),
			Text: utils.GetEnv("GBT_CAR_DIR_FORMAT", " {{ Dir }} "),
		},
		"Dir": {
			Bg: utils.GetEnv(
				"GBT_CAR_DIR_DIR_BG", utils.GetEnv(
					"GBT_CAR_DIR_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_DIR_DIR_FG", utils.GetEnv(
					"GBT_CAR_DIR_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_DIR_DIR_FM", utils.GetEnv(
					"GBT_CAR_DIR_FM", defaultRootFm)),
			Text: utils.GetEnv("GBT_CAR_DIR_DIR_TEXT", getDir()),
		},
		"Sep": {
			Bg: utils.GetEnv(
				"GBT_CAR_DIR_SEP_BG", utils.GetEnv(
					"GBT_SEPARATOR_BG", defaultSep)),
			Fg: utils.GetEnv(
				"GBT_CAR_DIR_SEP_FG", utils.GetEnv(
					"GBT_SEPARATOR_FG", defaultSep)),
			Fm: utils.GetEnv(
				"GBT_CAR_DIR_SEP_FM", utils.GetEnv(
					"GBT_SEPARATOR_FM", defaultSep)),
			Text: utils.GetEnv(
				"GBT_CAR_DIR_SEP", utils.GetEnv(
					"GBT_CAR_DIR_SEP_TEXT", utils.GetEnv(
						"GBT_SEPARATOR", defaultSep))),
		},
	}

	c.Display = utils.GetEnvBool("GBT_CAR_DIR_DISPLAY", true)
	c.Wrap = utils.GetEnvBool("GBT_CAR_DIR_WRAP", false)
}
