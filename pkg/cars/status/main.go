package status

import (
	"github.com/jtyr/gbt/pkg/core/car"
	"github.com/jtyr/gbt/pkg/core/utils"
)

// Car inherits the core.Car.
type Car struct {
	car.Car
}

func (c *Car) getSignal() (signal string) {
	_, argsExist := c.Params["args"]

	if !argsExist {
		return "?"
	}

	// The bellow statuses are based on the following URLs:
	// https://github.com/bric3/nice-exit-code/blob/master/nice-exit-code.plugin.zsh
	// http://tldp.org/LDP/abs/html/exitcodes.html
	// https://unix.stackexchange.com/a/254747/53489
	switch c.Params["args"] {
	// Usual exit codes
	case "-1":
		return "FATAL"
	case "0":
		return "OK"
	case "1":
		return "FAIL"
	case "2":
		return "BLTINMUSE"
	case "6":
		return "UNKADDR"

	// Issue with the actual command being invoked
	case "126":
		return "NOEXEC"
	case "127":
		return "NOTFOUND"

	// Signal errors (128 + signal)
	case "129":
		return "SIGHUP"
	case "130":
		return "SIGINT"
	case "131":
		return "SIGQUIT"
	case "132":
		return "SIGILL"
	case "133":
		return "SIGTRAP"
	case "134":
		return "SIGABRT"
	case "135":
		return "SIGBUS"
	case "136":
		return "SIGFPE"
	case "137":
		return "SIGKILL"
	case "138":
		return "SIGUSR1"
	case "139":
		return "SIGSEGV"
	case "140":
		return "SIGUSR2"
	case "141":
		return "SIGPIPE"
	case "142":
		return "SIGALRM"
	case "143":
		return "SIGTERM"
	case "145":
		return "SIGCHLD"
	case "146":
		return "SIGCONT"
	case "147":
		return "SIGSTOP"
	case "148":
		return "SIGTSTP"
	case "149":
		return "SIGTTIN"
	case "150":
		return "SIGTTOU"

	// Anything else is unknown
	default:
		return "UNK"
	}
}

// Checks for the return code.
func (c *Car) isOk() (ret bool) {
	_, argsExist := c.Params["args"]

	if !argsExist || c.Params["args"] == "0" {
		ret = true
	} else {
		ret = false
	}

	return
}

// Init initializes the car.
func (c *Car) Init() {
	defaultErrorBg := utils.GetEnv("GBT_CAR_BG", "red")
	defaultErrorFg := utils.GetEnv("GBT_CAR_FG", "light_gray")
	defaultErrorFm := utils.GetEnv("GBT_CAR_FM", "none")
	defaultOkBg := utils.GetEnv("GBT_CAR_BG", "green")
	defaultOkFg := utils.GetEnv("GBT_CAR_FG", "light_gray")
	defaultOkFm := utils.GetEnv("GBT_CAR_FM", "none")
	defaultRootBg := defaultErrorBg
	defaultRootFg := defaultErrorFg
	defaultRootFm := defaultErrorFm
	defaultSep := "\000"

	defaultDetailsFormat := " {{ Signal }}"
	defaultSymbolFormat := "{{ Error }}"
	defaultCodeText := "?"

	if val, ok := c.Params["args"]; ok {
		defaultCodeText = val.(string)
	}

	if c.isOk() {
		defaultRootBg = defaultOkBg
		defaultRootFg = defaultOkFg
		defaultRootFm = defaultOkFm
		defaultDetailsFormat = ""
		defaultSymbolFormat = "{{ Ok }}"
	} else {
		defaultDetailsFormat = utils.GetEnv(
			"GBT_CAR_STATUS_DETAILS_FORMAT", defaultDetailsFormat)
	}

	c.Model = map[string]car.ModelElement{
		"root": {
			Bg:   utils.GetEnv("GBT_CAR_STATUS_BG", defaultRootBg),
			Fg:   utils.GetEnv("GBT_CAR_STATUS_FG", defaultRootFg),
			Fm:   utils.GetEnv("GBT_CAR_STATUS_FM", defaultRootFm),
			Text: utils.GetEnv("GBT_CAR_STATUS_FORMAT", " {{ Symbol }} "),
		},
		"Error": {
			Bg: utils.GetEnv(
				"GBT_CAR_STATUS_ERROR_BG", utils.GetEnv(
					"GBT_CAR_STATUS_SYMBOL_BG", utils.GetEnv(
						"GBT_CAR_STATUS_BG", defaultErrorBg))),
			Fg: utils.GetEnv(
				"GBT_CAR_STATUS_ERROR_FG", utils.GetEnv(
					"GBT_CAR_STATUS_SYMBOL_FG", utils.GetEnv(
						"GBT_CAR_STATUS_FG", defaultErrorFg))),
			Fm: utils.GetEnv(
				"GBT_CAR_STATUS_ERROR_FM", utils.GetEnv(
					"GBT_CAR_STATUS_SYMBOL_FM", utils.GetEnv(
						"GBT_CAR_STATUS_FM", defaultErrorFm))),
			Text: utils.GetEnv("GBT_CAR_STATUS_ERROR_TEXT", "\u2718"),
		},
		"Ok": {
			Bg: utils.GetEnv(
				"GBT_CAR_STATUS_OK_BG", utils.GetEnv(
					"GBT_CAR_STATUS_SYMBOL_BG", utils.GetEnv(
						"GBT_CAR_STATUS_BG", defaultOkBg))),
			Fg: utils.GetEnv(
				"GBT_CAR_STATUS_OK_FG", utils.GetEnv(
					"GBT_CAR_STATUS_SYMBOL_FG", utils.GetEnv(
						"GBT_CAR_STATUS_FG", defaultOkFg))),
			Fm: utils.GetEnv(
				"GBT_CAR_STATUS_OK_FM", utils.GetEnv(
					"GBT_CAR_STATUS_SYMBOL_FM", utils.GetEnv(
						"GBT_CAR_STATUS_FM", defaultOkFm))),
			Text: utils.GetEnv("GBT_CAR_STATUS_OK_TEXT", "\u2714"),
		},
		"Symbol": {
			Bg: utils.GetEnv(
				"GBT_CAR_STATUS_SYMBOL_BG", utils.GetEnv(
					"GBT_CAR_STATUS_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_STATUS_SYMBOL_FG", utils.GetEnv(
					"GBT_CAR_STATUS_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_STATUS_SYMBOL_FM", utils.GetEnv(
					"GBT_CAR_STATUS_FM", defaultRootFm)),
			Text: utils.GetEnv(
				"GBT_CAR_STATUS_SYMBOL_FORMAT", defaultSymbolFormat),
		},
		"Details": {
			Bg: utils.GetEnv(
				"GBT_CAR_STATUS_DETAILS_BG", utils.GetEnv(
					"GBT_CAR_STATUS_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_STATUS_DETAILS_FG", utils.GetEnv(
					"GBT_CAR_STATUS_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_STATUS_DETAILS_FM", utils.GetEnv(
					"GBT_CAR_STATUS_FM", defaultRootFm)),
			Text: defaultDetailsFormat,
		},
		"Code": {
			Bg: utils.GetEnv(
				"GBT_CAR_STATUS_CODE_BG", utils.GetEnv(
					"GBT_CAR_STATUS_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_STATUS_CODE_FG", utils.GetEnv(
					"GBT_CAR_STATUS_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_STATUS_CODE_FM", utils.GetEnv(
					"GBT_CAR_STATUS_FM", defaultRootFm)),
			Text: utils.GetEnv(
				"GBT_CAR_STATUS_CODE_TEXT", defaultCodeText),
		},
		"Signal": {
			Bg: utils.GetEnv(
				"GBT_CAR_STATUS_SIGNAL_BG", utils.GetEnv(
					"GBT_CAR_STATUS_BG", defaultRootBg)),
			Fg: utils.GetEnv(
				"GBT_CAR_STATUS_SIGNAL_FG", utils.GetEnv(
					"GBT_CAR_STATUS_FG", defaultRootFg)),
			Fm: utils.GetEnv(
				"GBT_CAR_STATUS_SIGNAL_FM", utils.GetEnv(
					"GBT_CAR_STATUS_FM", defaultRootFm)),
			Text: utils.GetEnv(
				"GBT_CAR_STATUS_SIGNAL_TEXT", c.getSignal()),
		},
		"Sep": {
			Bg: utils.GetEnv(
				"GBT_CAR_STATUS_SEP_BG", utils.GetEnv(
					"GBT_SEPARATOR_BG", defaultSep)),
			Fg: utils.GetEnv(
				"GBT_CAR_STATUS_SEP_FG", utils.GetEnv(
					"GBT_SEPARATOR_FG", defaultSep)),
			Fm: utils.GetEnv(
				"GBT_CAR_STATUS_SEP_FM", utils.GetEnv(
					"GBT_SEPARATOR_FM", defaultSep)),
			Text: utils.GetEnv(
				"GBT_CAR_STATUS_SEP", utils.GetEnv(
					"GBT_CAR_STATUS_SEP_TEXT", utils.GetEnv(
						"GBT_SEPARATOR", defaultSep))),
		},
	}

	if c.isOk() {
		c.Display = utils.GetEnvBool("GBT_CAR_STATUS_DISPLAY", false)
	} else {
		c.Display = utils.GetEnvBool("GBT_CAR_STATUS_DISPLAY", true)
	}

	c.Wrap = utils.GetEnvBool("GBT_CAR_STATUS_WRAP", false)
}
