package ttime

import (
	"time"

	"github.com/jtyr/gbt/pkg/core/car"
	"github.com/jtyr/gbt/pkg/core/utils"
)

// Car inherits the core.Car.
type Car struct {
	car.Car
}

// Reference to the time.Now() function.
var tnow = time.Now

// Init initializes the car.
func (c *Car) Init() {
	defaultRootBg := utils.GetEnv("GBT_CAR_BG", "light_blue")
	defaultRootFg := utils.GetEnv("GBT_CAR_FG", "light_gray")
	defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
	defaultSep := "\000"

	c.Model = map[string]car.ModelElement{
		"root": {
			Bg:   utils.GetEnv("GBT_CAR_TIME_BG", defaultRootBg),
			Fg:   utils.GetEnv("GBT_CAR_TIME_FG", defaultRootFg),
			Fm:   utils.GetEnv("GBT_CAR_TIME_FM", defaultRootFm),
			Text: utils.GetEnv("GBT_CAR_TIME_FORMAT", " {{ DateTime }} "),
		},
		"DateTime": {
			Bg: utils.GetEnv(
				"GBT_CAR_TIME_DATETIME_BG", utils.GetEnv(
					"GBT_CAR_TIME_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_TIME_DATETIME_FG", utils.GetEnv(
					"GBT_CAR_TIME_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_TIME_DATETIME_FM", utils.GetEnv(
					"GBT_CAR_TIME_FM", defaultRootFm)),
			Text: utils.GetEnv(
				"GBT_CAR_TIME_DATETIME_FORMAT", "{{ Date }} {{ Time }}"),
		},
		"Date": {
			Bg: utils.GetEnv(
				"GBT_CAR_TIME_DATE_BG", utils.GetEnv(
					"GBT_CAR_TIME_DATETIME_BG", utils.GetEnv(
						"GBT_CAR_TIME_BG", defaultRootBg))),
			Fg: utils.GetEnv(
				"GBT_CAR_TIME_DATE_FG", utils.GetEnv(
					"GBT_CAR_TIME_DATETIME_FG", utils.GetEnv(
						"GBT_CAR_TIME_FG", defaultRootFg))),
			Fm: utils.GetEnv(
				"GBT_CAR_TIME_DATE_FM", utils.GetEnv(
					"GBT_CAR_TIME_DATETIME_FM", utils.GetEnv(
						"GBT_CAR_TIME_FM", defaultRootFm))),
			Text: tnow().Format(
				utils.GetEnv("GBT_CAR_TIME_DATE_FORMAT", "Mon 02 Jan")),
		},
		"Time": {
			Bg: utils.GetEnv(
				"GBT_CAR_TIME_TIME_BG", utils.GetEnv(
					"GBT_CAR_TIME_DATETIME_BG", utils.GetEnv(
						"GBT_CAR_TIME_BG", defaultRootBg))),
			Fg: utils.GetEnv(
				"GBT_CAR_TIME_TIME_FG", utils.GetEnv(
					"GBT_CAR_TIME_DATETIME_FG", utils.GetEnv(
						"GBT_CAR_TIME_FG", "light_yellow"))),
			Fm: utils.GetEnv(
				"GBT_CAR_TIME_TIME_FM", utils.GetEnv(
					"GBT_CAR_TIME_DATETIME_FM", utils.GetEnv(
						"GBT_CAR_TIME_FM", defaultRootFm))),
			Text: tnow().Format(
				utils.GetEnv("GBT_CAR_TIME_TIME_FORMAT", "15:04:05")),
		},
		"Sep": {
			Bg: utils.GetEnv(
				"GBT_CAR_TIME_SEP_BG", utils.GetEnv(
					"GBT_SEPARATOR_BG", defaultSep)),
			Fg: utils.GetEnv(
				"GBT_CAR_TIME_SEP_FG", utils.GetEnv(
					"GBT_SEPARATOR_FG", defaultSep)),
			Fm: utils.GetEnv(
				"GBT_CAR_TIME_SEP_FM", utils.GetEnv(
					"GBT_SEPARATOR_FM", defaultSep)),
			Text: utils.GetEnv(
				"GBT_CAR_TIME_SEP", utils.GetEnv(
					"GBT_CAR_TIME_SEP_TEXT", utils.GetEnv(
						"GBT_SEPARATOR", defaultSep))),
		},
	}

	c.Display = utils.GetEnvBool("GBT_CAR_TIME_DISPLAY", true)
	c.Wrap = utils.GetEnvBool("GBT_CAR_TIME_WRAP", false)
}
