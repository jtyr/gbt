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
	defaultSep := "\000"

	curUser, _ := user.Current()
	hostname, _ := os.Hostname()
	hostname = strings.Split(hostname, ".")[0]

	uaFormat := "{{ User }}"

	if curUser.Uid == adminUID {
		uaFormat = "{{ Admin }}"
	}

	c.Model = map[string]car.ModelElement{
		"root": {
			Bg:   utils.GetEnv("GBT_CAR_HOSTNAME_BG", defaultRootBg),
			Fg:   utils.GetEnv("GBT_CAR_HOSTNAME_FG", defaultRootFg),
			Fm:   utils.GetEnv("GBT_CAR_HOSTNAME_FM", defaultRootFm),
			Text: utils.GetEnv("GBT_CAR_HOSTNAME_FORMAT", " {{ UserHost }} "),
		},
		"UserHost": {
			Bg: utils.GetEnv(
				"GBT_CAR_HOSTNAME_USERHOST_BG", utils.GetEnv(
					"GBT_CAR_HOSTNAME_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_HOSTNAME_USERHOST_FG", utils.GetEnv(
					"GBT_CAR_HOSTNAME_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_HOSTNAME_USERHOST_FM", utils.GetEnv(
					"GBT_CAR_HOSTNAME_FM", defaultRootFm)),
			Text: utils.GetEnv(
				"GBT_CAR_HOSTNAME_USERHOST_FORMAT", fmt.Sprintf("%s@{{ Host }}", uaFormat)),
		},
		"Admin": {
			Bg: utils.GetEnv(
				"GBT_CAR_HOSTNAME_ADMIN_BG", utils.GetEnv(
					"GBT_CAR_HOSTNAME_USER_BG", utils.GetEnv(
						"GBT_CAR_HOSTNAME_BG", defaultRootBg))),
			Fg: utils.GetEnv(
				"GBT_CAR_HOSTNAME_ADMIN_FG", utils.GetEnv(
					"GBT_CAR_HOSTNAME_USER_FG", utils.GetEnv(
						"GBT_CAR_HOSTNAME_FG", defaultRootFg))),
			Fm: utils.GetEnv(
				"GBT_CAR_HOSTNAME_ADMIN_FM", utils.GetEnv(
					"GBT_CAR_HOSTNAME_USER_FM", utils.GetEnv(
						"GBT_CAR_HOSTNAME_FM", defaultRootFm))),
			Text: utils.GetEnv("GBT_CAR_HOSTNAME_ADMIN_TEXT", curUser.Username),
		},
		"User": {
			Bg: utils.GetEnv(
				"GBT_CAR_HOSTNAME_USER_BG", utils.GetEnv(
					"GBT_CAR_HOSTNAME_USER_BG", utils.GetEnv(
						"GBT_CAR_HOSTNAME_BG", defaultRootBg))),
			Fg: utils.GetEnv(
				"GBT_CAR_HOSTNAME_USER_FG", utils.GetEnv(
					"GBT_CAR_HOSTNAME_USER_FG", utils.GetEnv(
						"GBT_CAR_HOSTNAME_FG", defaultRootFg))),
			Fm: utils.GetEnv(
				"GBT_CAR_HOSTNAME_USER_FM", utils.GetEnv(
					"GBT_CAR_HOSTNAME_USER_FM", utils.GetEnv(
						"GBT_CAR_HOSTNAME_FM", defaultRootFm))),
			Text: utils.GetEnv("GBT_CAR_HOSTNAME_USER_TEXT", curUser.Username),
		},
		"Host": {
			Bg: utils.GetEnv(
				"GBT_CAR_HOSTNAME_HOST_BG", utils.GetEnv(
					"GBT_CAR_HOSTNAME_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_HOSTNAME_HOST_FG", utils.GetEnv(
					"GBT_CAR_HOSTNAME_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_HOSTNAME_HOST_FM", utils.GetEnv(
					"GBT_CAR_HOSTNAME_FM", defaultRootFm)),
			Text: utils.GetEnv("GBT_CAR_HOSTNAME_HOST_TEXT", hostname),
		},
		"Sep": {
			Bg: utils.GetEnv(
				"GBT_CAR_HOSTNAME_SEP_BG", utils.GetEnv(
					"GBT_SEPARATOR_BG", defaultSep)),
			Fg: utils.GetEnv(
				"GBT_CAR_HOSTNAME_SEP_FG", utils.GetEnv(
					"GBT_SEPARATOR_FG", defaultSep)),
			Fm: utils.GetEnv(
				"GBT_CAR_HOSTNAME_SEP_FM", utils.GetEnv(
					"GBT_SEPARATOR_FM", defaultSep)),
			Text: utils.GetEnv(
				"GBT_CAR_HOSTNAME_SEP", utils.GetEnv(
					"GBT_CAR_HOSTNAME_SEP_TEXT", utils.GetEnv(
						"GBT_SEPARATOR", defaultSep))),
		},
	}

	c.Display = utils.GetEnvBool("GBT_CAR_HOSTNAME_DISPLAY", true)
	c.Wrap = utils.GetEnvBool("GBT_CAR_HOSTNAME_WRAP", false)
}
