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
	defaultSep := "\000"

	symbolFormat := "{{ User }}"
	curUser, _ := user.Current()

	if curUser.Uid == adminUID {
		symbolFormat = "{{ Admin }}"
	}

	c.Model = map[string]car.ModelElement{
		"root": {
			Bg:   utils.GetEnv("GBT_CAR_SIGN_BG", defaultRootBg),
			Fg:   utils.GetEnv("GBT_CAR_SIGN_FG", defaultRootFg),
			Fm:   utils.GetEnv("GBT_CAR_SIGN_FM", defaultRootFm),
			Text: utils.GetEnv("GBT_CAR_SIGN_FORMAT", " {{ Symbol }} "),
		},
		"Symbol": {
			Bg: utils.GetEnv(
				"GBT_CAR_SIGN_SYMBOL_BG", utils.GetEnv(
					"GBT_CAR_SIGN_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_SIGN_SYMBOL_FG", utils.GetEnv(
					"GBT_CAR_SIGN_FG", "green")),
			Fm: utils.GetEnv(
				"GBT_CAR_SIGN_SYMBOL_FM", utils.GetEnv(
					"GBT_CAR_SIGN_FM", "bold")),
			Text: utils.GetEnv("GBT_CAR_SIGN_SYMBOL_FORMAT", symbolFormat),
		},
		"User": {
			Bg: utils.GetEnv(
				"GBT_CAR_SIGN_USER_BG", utils.GetEnv(
					"GBT_CAR_SIGN_BG", defaultRootBg)),
			Fg: utils.GetEnv("GBT_CAR_SIGN_USER_FG", "light_green"),
			Fm: utils.GetEnv(
				"GBT_CAR_SIGN_USER_FM", utils.GetEnv(
					"GBT_CAR_SIGN_FM", defaultRootFm)),
			Text: utils.GetEnv("GBT_CAR_SIGN_USER_TEXT", "$"),
		},
		"Admin": {
			Bg: utils.GetEnv(
				"GBT_CAR_SIGN_ADMIN_BG", utils.GetEnv(
					"GBT_CAR_SIGN_BG", defaultRootBg)),
			Fg: utils.GetEnv("GBT_CAR_SIGN_ADMIN_FG", "red"),
			Fm: utils.GetEnv(
				"GBT_CAR_SIGN_ADMIN_FM", utils.GetEnv(
					"GBT_CAR_SIGN_FM", defaultRootFm)),
			Text: utils.GetEnv("GBT_CAR_SIGN_ADMIN_TEXT", "#"),
		},
		"Sep": {
			Bg: utils.GetEnv(
				"GBT_CAR_SIGN_SEP_BG", utils.GetEnv(
					"GBT_SEPARATOR_BG", defaultSep)),
			Fg: utils.GetEnv(
				"GBT_CAR_SIGN_SEP_FG", utils.GetEnv(
					"GBT_SEPARATOR_FG", defaultSep)),
			Fm: utils.GetEnv(
				"GBT_CAR_SIGN_SEP_FM", utils.GetEnv(
					"GBT_SEPARATOR_FM", defaultSep)),
			Text: utils.GetEnv(
				"GBT_CAR_SIGN_SEP", utils.GetEnv(
					"GBT_CAR_SIGN_SEP_TEXT", utils.GetEnv(
						"GBT_SEPARATOR", defaultSep))),
		},
	}

	c.Display = utils.GetEnvBool("GBT_CAR_SIGN_DISPLAY", true)
	c.Wrap = utils.GetEnvBool("GBT_CAR_SIGN_WRAP", false)
}
