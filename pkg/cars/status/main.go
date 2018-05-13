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

    if ! argsExist {
        return "?"
    }

    // The bellow statuses are based on the following URLs:
    // https://github.com/bric3/nice-exit-code/blob/master/nice-exit-code.plugin.zsh
    // http://tldp.org/LDP/abs/html/exitcodes.html
    // https://unix.stackexchange.com/a/254747/53489
    switch c.Params["args"] {
        // Usual exit codes
        case  "-1": return "FATAL"
        case   "0": return "OK"
        case   "1": return "FAIL"
        case   "2": return "BLTINMUSE"
        case   "6": return "UNKADDR"

        // Issue with the actual command being invoked
        case "126": return "NOEXEC"
        case "127": return "NOTFOUND"

        // Signal errors (128 + signal)
        case "129": return "SIGHUP"
        case "130": return "SIGINT"
        case "131": return "SIGQUIT"
        case "132": return "SIGILL"
        case "133": return "SIGTRAP"
        case "134": return "SIGABRT"
        case "135": return "SIGBUS"
        case "136": return "SIGFPE"
        case "137": return "SIGKILL"
        case "138": return "SIGUSR1"
        case "139": return "SIGSEGV"
        case "140": return "SIGUSR2"
        case "141": return "SIGPIPE"
        case "142": return "SIGALRM"
        case "143": return "SIGTERM"
        case "145": return "SIGCHLD"
        case "146": return "SIGCONT"
        case "147": return "SIGSTOP"
        case "148": return "SIGTSTP"
        case "149": return "SIGTTIN"
        case "150": return "SIGTTOU"

        // Anything else is unknown
        default:    return "UNK"
    }
}

// Checks for the return code.
func (c *Car) isOk() (ret bool) {
    _, argsExist := c.Params["args"]

    if ! argsExist || c.Params["args"] == "0" {
        ret = true
    } else {
        ret = false
    }

    return
}

// Init initializes the car.
func (c *Car) Init() {
    defaultErrorBg := "red"
    defaultErrorFg := "light_gray"
    defaultErrorFm := "none"
    defaultOkBg := "green"
    defaultOkFg := "light_gray"
    defaultOkFm := "none"
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", defaultErrorBg)
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", defaultErrorFg)
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", defaultErrorFm)
    defaultSymbolBg := defaultRootBg
    defaultSymbolFg := defaultRootFg
    defaultSymbolFm := defaultRootFm
    defaultDetailsBg := defaultRootBg
    defaultDetailsFg := defaultRootFg
    defaultDetailsFm := defaultRootFm
    defaultCodeBg := defaultRootBg
    defaultCodeFg := defaultRootFg
    defaultCodeFm := defaultRootFm
    defaultSignalBg := defaultRootBg
    defaultSignalFg := defaultRootFg
    defaultSignalFm := defaultRootFm

    defaultDetailsFormat := " {{ Signal }}"
    defaultSymbolFormat := "{{ Error }}"
    defaultCodeText := "?"

    if val, ok := c.Params["args"]; ok {
        defaultCodeText = val.(string)
    }

    if c.isOk() {
        defaultRootBg = utils.GetEnv("GBT_CAR_BG", defaultOkBg)
        defaultRootFg = utils.GetEnv("GBT_CAR_FG", defaultOkFg)
        defaultRootFm = utils.GetEnv("GBT_CAR_FM", defaultOkFm)
        defaultDetailsFormat = ""
        defaultSymbolFormat = "{{ Ok }}"
    } else {
        defaultDetailsFormat = utils.GetEnv(
            "GBT_CAR_STATUS_DETAILS_FORMAT", defaultDetailsFormat)
    }

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_STATUS_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_STATUS_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_STATUS_FM", defaultRootFm),
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
            Text: utils.GetEnv("GBT_CAR_STATUS_ERROR_TEXT", "✘"),
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
            Text: utils.GetEnv("GBT_CAR_STATUS_OK_TEXT", "✔"),
        },
        "Symbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_STATUS_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_STATUS_BG", defaultSymbolBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_STATUS_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_STATUS_FG", defaultSymbolFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_STATUS_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_STATUS_FM", defaultSymbolFm)),
            Text: utils.GetEnv(
                "GBT_CAR_STATUS_SYMBOL_FORMAT", defaultSymbolFormat),
        },
        "Details": {
            Bg: utils.GetEnv(
                "GBT_CAR_STATUS_DETAILS_BG", utils.GetEnv(
                    "GBT_CAR_STATUS_BG", defaultDetailsBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_STATUS_DETAILS_FG", utils.GetEnv(
                    "GBT_CAR_STATUS_FG", defaultDetailsFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_STATUS_DETAILS_FM", utils.GetEnv(
                    "GBT_CAR_STATUS_FM", defaultDetailsFm)),
            Text: defaultDetailsFormat,
        },
        "Code": {
            Bg: utils.GetEnv(
                "GBT_CAR_STATUS_CODE_BG", utils.GetEnv(
                    "GBT_CAR_STATUS_BG", defaultCodeBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_STATUS_CODE_FG", utils.GetEnv(
                    "GBT_CAR_STATUS_FG", defaultCodeFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_STATUS_CODE_FM", utils.GetEnv(
                    "GBT_CAR_STATUS_FM", defaultCodeFm)),
            Text: utils.GetEnv(
                "GBT_CAR_STATUS_CODE_TEXT", defaultCodeText),
        },
        "Signal": {
            Bg: utils.GetEnv(
                "GBT_CAR_STATUS_SIGNAL_BG", utils.GetEnv(
                    "GBT_CAR_STATUS_BG", defaultSignalBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_STATUS_SIGNAL_FG", utils.GetEnv(
                    "GBT_CAR_STATUS_FG", defaultSignalFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_STATUS_SIGNAL_FM", utils.GetEnv(
                    "GBT_CAR_STATUS_FM", defaultSignalFm)),
            Text: utils.GetEnv(
                "GBT_CAR_STATUS_SIGNAL_TEXT", c.getSignal()),
        },
    }

    if c.isOk() {
        c.Display = utils.GetEnvBool("GBT_CAR_STATUS_DISPLAY", false)
    } else {
        c.Display = utils.GetEnvBool("GBT_CAR_STATUS_DISPLAY", true)
    }

    c.Wrap = utils.GetEnvBool("GBT_CAR_STATUS_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_STATUS_SEP", "\000")
}
