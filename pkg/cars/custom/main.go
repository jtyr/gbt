package custom

import (
	"fmt"
	"strings"

	"github.com/jtyr/gbt/pkg/core/car"
	"github.com/jtyr/gbt/pkg/core/utils"
)

// Car inherits the core.Car.
type Car struct {
	car.Car
}

// Init initializes the car.
func (c *Car) Init() {
	defaultRootBg := utils.GetEnv("GBT_CAR_BG", "130")
	defaultRootFg := utils.GetEnv("GBT_CAR_FG", "white")
	defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
	defaultSep := "\000"

	prefix := fmt.Sprintf("GBT_CAR_CUSTOM%s", strings.ToUpper(c.Params["name"].(string)))
	defaultTextText := "?"
	defaultTextCmd := utils.GetEnv(fmt.Sprintf("%s_TEXT_CMD", prefix), "")
	defaultDisplayCmd := utils.GetEnv(fmt.Sprintf("%s_DISPLAY_CMD", prefix), "")
	defaultDisplay := true

	if defaultTextCmd != "" {
		shellExecutor := utils.GetEnv(
			fmt.Sprintf("%s_TEXT_EXECUTOR", prefix),
			utils.GetEnv("GBT_CAR_CUSTOM_EXECUTOR", "sh"))
		shellExecutorParam := utils.GetEnv(
			fmt.Sprintf("%s_TEXT_EXECUTOR_PARAM", prefix),
			utils.GetEnv("GBT_CAR_CUSTOM_EXECUTOR_PARAM", "-c"))
		_, defaultTextText, _ = utils.Run([]string{shellExecutor, shellExecutorParam, defaultTextCmd})
	}

	if defaultDisplayCmd != "" {
		shellExecutor := utils.GetEnv(
			fmt.Sprintf("%s_DISPLAY_EXECUTOR", prefix),
			utils.GetEnv("GBT_CAR_CUSTOM_EXECUTOR", "sh"))
		shellExecutorParam := utils.GetEnv(
			fmt.Sprintf("%s_TEXT_DISPLAY_EXECUTOR_PARAM", prefix),
			utils.GetEnv("GBT_CAR_CUSTOM_EXECUTOR_PARAM", "-c"))
		_, defaultDisplayOutput, _ := utils.Run([]string{shellExecutor, shellExecutorParam, defaultDisplayCmd})

		if !utils.IsTrue(defaultDisplayOutput) {
			defaultDisplay = false
		}
	}

	c.Model = map[string]car.ModelElement{
		"root": {
			Bg:   utils.GetEnv(fmt.Sprintf("%s_BG", prefix), defaultRootBg),
			Fg:   utils.GetEnv(fmt.Sprintf("%s_FG", prefix), defaultRootFg),
			Fm:   utils.GetEnv(fmt.Sprintf("%s_FM", prefix), defaultRootFm),
			Text: utils.GetEnv(fmt.Sprintf("%s_FORMAT", prefix), " {{ Text }} "),
		},
		"Text": {
			Bg: utils.GetEnv(
				fmt.Sprintf("%s_TEXT_BG", prefix), utils.GetEnv(
					fmt.Sprintf("%s_BG", prefix), defaultRootBg)),
			Fg: utils.GetEnv(
				fmt.Sprintf("%s_TEXT_FG", prefix), utils.GetEnv(
					fmt.Sprintf("%s_FG", prefix), defaultRootFg)),
			Fm: utils.GetEnv(
				fmt.Sprintf("%s_TEXT_FM", prefix), utils.GetEnv(
					fmt.Sprintf("%s_FM", prefix), defaultRootFm)),
			Text: utils.GetEnv(
				fmt.Sprintf("%s_TEXT_TEXT", prefix), defaultTextText),
		},
		"Sep": {
			Bg: utils.GetEnv(
				fmt.Sprintf("%s_SEP_BG", prefix), utils.GetEnv(
					"GBT_SEPARATOR_BG", defaultSep)),
			Fg: utils.GetEnv(
				fmt.Sprintf("%s_SEP_FG", prefix), utils.GetEnv(
					"GBT_SEPARATOR_FG", defaultSep)),
			Fm: utils.GetEnv(
				fmt.Sprintf("%s_SEP_FM", prefix), utils.GetEnv(
					"GBT_SEPARATOR_FM", defaultSep)),
			Text: utils.GetEnv(
				fmt.Sprintf("%s_SEP", prefix), utils.GetEnv(
					fmt.Sprintf("%s_SEP_TEXT", prefix), utils.GetEnv(
						"GBT_SEPARATOR", defaultSep))),
		},
	}

	c.Display = utils.GetEnvBool(fmt.Sprintf("%s_DISPLAY", prefix), defaultDisplay)
	c.Wrap = utils.GetEnvBool(fmt.Sprintf("%s_WRAP", prefix), false)
}
