package exectime

import (
	"fmt"
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

// Returns duration of the execution time.
func getDuration() string {
	precision := utils.GetEnvInt("GBT_CAR_EXECTIME_PRECISION", 0)
	now := float64(tnow().UnixNano()) / float64(1e9)
	execs := now - utils.GetEnvFloat("GBT_CAR_EXECTIME_SECS", now)
	subsecs := execs - float64(int(execs))

	hours := int(execs / 3600)
	mins := int((execs - float64(hours)*3600) / 60)
	secs := int(execs - float64(hours)*3600 - float64(mins)*60)
	millis := 0
	micros := 0
	nanos := 0

	durationtime := ""

	if precision > 0 {
		subsecs *= 1e3
		millis = int(subsecs)

		if precision > 3 || (secs == 0 && millis == 0) {
			subsecs = (subsecs - float64(int(subsecs))) * 1e3
			micros = int(subsecs)

			if precision > 6 || (secs == 0 && millis == 0 && micros == 0) {
				subsecs = (subsecs - float64(int(subsecs))) * 1e3
				nanos = int(subsecs)
			}
		}
	}

	if hours > 0 {
		durationtime += fmt.Sprintf("%dh", hours)
	}

	if mins > 0 {
		durationtime += fmt.Sprintf("%dm", mins)
	}

	if secs > 0 || precision == 0 {
		durationtime += fmt.Sprintf("%ds", secs)
	}

	if millis > 0 {
		durationtime += fmt.Sprintf("%dms", millis)
	}

	if micros > 0 {
		durationtime += fmt.Sprintf("%dµs", micros)
	}

	if nanos > 0 {
		durationtime += fmt.Sprintf("%dns", nanos)
	}

	return durationtime
}

// Returns execution time in seconds.
func getSeconds() string {
	precision := utils.GetEnvInt("GBT_CAR_EXECTIME_PRECISION", 0)
	now := float64(tnow().UnixNano()) / float64(1e9)
	execs := now - utils.GetEnvFloat("GBT_CAR_EXECTIME_SECS", now)

	secondstime := fmt.Sprintf("%0.*f", precision, execs)

	return secondstime
}

// Returns the execution time.
func getTime() string {
	precision := utils.GetEnvInt("GBT_CAR_EXECTIME_PRECISION", 0)
	now := float64(tnow().UnixNano()) / float64(1e9)
	execs := now - utils.GetEnvFloat("GBT_CAR_EXECTIME_SECS", now)
	subsecs := ""

	if precision > 0 {
		subsecs = "."
		subsecs += fmt.Sprintf("%0.*f", precision, execs-float64(int(execs)))[2:]
	}

	hours := int(execs / 3600)
	mins := int((execs - float64(hours)*3600) / 60)
	secs := int(execs - float64(hours)*3600 - float64(mins)*60)

	exectime := fmt.Sprintf("%.2d:%.2d:%02d%s", hours, mins, secs, subsecs)

	return exectime
}

// Init initializes the car.
func (c *Car) Init() {
	defaultRootBg := utils.GetEnv("GBT_CAR_BG", "light_gray")
	defaultRootFg := utils.GetEnv("GBT_CAR_FG", "black")
	defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
	defaultSep := "\000"

	c.Model = map[string]car.ModelElement{
		"root": {
			Bg:   utils.GetEnv("GBT_CAR_EXECTIME_BG", defaultRootBg),
			Fg:   utils.GetEnv("GBT_CAR_EXECTIME_FG", defaultRootFg),
			Fm:   utils.GetEnv("GBT_CAR_EXECTIME_FM", defaultRootFm),
			Text: utils.GetEnv("GBT_CAR_EXECTIME_FORMAT", " {{ Time }} "),
		},
		"Duration": {
			Bg: utils.GetEnv(
				"GBT_CAR_EXECTIME_DURATION_BG", utils.GetEnv(
					"GBT_CAR_EXECTIME_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_EXECTIME_DURATION_FG", utils.GetEnv(
					"GBT_CAR_EXECTIME_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_EXECTIME_DURATION_FM", utils.GetEnv(
					"GBT_CAR_EXECTIME_FM", defaultRootFm)),
			Text: utils.GetEnv("GBT_CAR_EXECTIME_DURATION_TEXT", getDuration()),
		},
		"Seconds": {
			Bg: utils.GetEnv(
				"GBT_CAR_EXECTIME_SECONDS_BG", utils.GetEnv(
					"GBT_CAR_EXECTIME_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_EXECTIME_SECONDS_FG", utils.GetEnv(
					"GBT_CAR_EXECTIME_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_EXECTIME_SECONDS_FM", utils.GetEnv(
					"GBT_CAR_EXECTIME_FM", defaultRootFm)),
			Text: utils.GetEnv("GBT_CAR_EXECTIME_SECONDS_TEXT", getSeconds()),
		},
		"Time": {
			Bg: utils.GetEnv(
				"GBT_CAR_EXECTIME_TIME_BG", utils.GetEnv(
					"GBT_CAR_EXECTIME_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_EXECTIME_TIME_FG", utils.GetEnv(
					"GBT_CAR_EXECTIME_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_EXECTIME_TIME_FM", utils.GetEnv(
					"GBT_CAR_EXECTIME_FM", defaultRootFm)),
			Text: utils.GetEnv("GBT_CAR_EXECTIME_TIME_TEXT", getTime()),
		},
		"Sep": {
			Bg: utils.GetEnv(
				"GBT_CAR_EXECTIME_SEP_BG", utils.GetEnv(
					"GBT_SEPARATOR_BG", defaultSep)),
			Fg: utils.GetEnv(
				"GBT_CAR_EXECTIME_SEP_FG", utils.GetEnv(
					"GBT_SEPARATOR_FG", defaultSep)),
			Fm: utils.GetEnv(
				"GBT_CAR_EXECTIME_SEP_FM", utils.GetEnv(
					"GBT_SEPARATOR_FM", defaultSep)),
			Text: utils.GetEnv(
				"GBT_CAR_EXECTIME_SEP", utils.GetEnv(
					"GBT_CAR_EXECTIME_SEP_TEXT", utils.GetEnv(
						"GBT_SEPARATOR", defaultSep))),
		},
	}

	c.Display = utils.GetEnvBool("GBT_CAR_EXECTIME_DISPLAY", true)
	c.Wrap = utils.GetEnvBool("GBT_CAR_EXECTIME_WRAP", false)
}
