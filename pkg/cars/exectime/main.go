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

// Returns the execution time.
func getTime() string {
    precision := utils.GetEnvInt("GBT_CAR_EXECTIME_PRECISION", 0)
    now := float64(tnow().UnixNano())/float64(1e9)
    execs := now - utils.GetEnvFloat("GBT_CAR_EXECTIME_SECS", now)
    subsecs := ""

    if precision > 0 {
        subsecs = "."
        subsecs += fmt.Sprintf("%0.*f", precision, execs - float64(int(execs)))[2:]
    }

    hours := int(execs/3600)
    mins := int((execs - float64(hours)*3600)/60)
    secs := int(execs - float64(hours)*3600 - float64(mins)*60)

    exectime := fmt.Sprintf("%.2d:%.2d:%02d%s", hours, mins, secs, subsecs)

    return exectime
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "light_gray")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "black")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultTimeBg := defaultRootBg
    defaultTimeFg := defaultRootFg
    defaultTimeFm := defaultRootFm

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_EXECTIME_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_EXECTIME_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_EXECTIME_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_EXECTIME_FORMAT", " {{ Time }} "),
        },
        "Time": {
            Bg: utils.GetEnv(
                "GBT_CAR_EXECTIME_TIME_BG", utils.GetEnv(
                    "GBT_CAR_EXECTIME_BG", defaultTimeBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_EXECTIME_TIME_FG", utils.GetEnv(
                    "GBT_CAR_EXECTIME_FG", defaultTimeFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_EXECTIME_TIME_FM", utils.GetEnv(
                    "GBT_CAR_EXECTIME_FM", defaultTimeFm)),
            Text: utils.GetEnv("GBT_CAR_EXECTIME_TIME_TEXT", getTime()),
        },
    }

    c.Display = utils.GetEnvBool("GBT_CAR_EXECTIME_DISPLAY", true)
    c.Wrap = utils.GetEnvBool("GBT_CAR_EXECTIME_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_EXECTIME_SEP", "\000")
}
