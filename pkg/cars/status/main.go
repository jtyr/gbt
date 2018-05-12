package status

import (
    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

func getMsgMapping(exitStatus string) (msg string) {
    // Inspired by code from zsh-prompt-powerline:
    // https://github.com/bric3/nice-exit-code/blob/master/nice-exit-code.plugin.zsh
    //
    // Statuses based on above script + the following URLs:
    // Advanced Bash-Scripting Guide: http://tldp.org/LDP/abs/html/exitcodes.html
    // StackExchange: https://unix.stackexchange.com/a/254747/53489
    //
    // Unfortunately no good resources that provide the same details for zsh and/or other shells
    switch exitStatus {
    // usual exit codes
    case "-1": return "FATAL"
    case "0": return "OK"
    case "1": return "FAIL" // Miscellaneous errors, such as "divide by zero"
    case "2": return "BUILTINMISUSE" // misuse of shell builtins (pretty rare)
    case "6": return "UNKADDR" // Unknown address or device, e.g.: curl foo; echo $?

    // issue with the actual command being invoked
    case "126": return "NOEXEC" // cannot invoke requested command (ex : source script_with_syntax_error)
    case "127": return "NOTFOUND" // command not found (ex : source script_not_existing)

    // errors from signal (error code = 128 + signal). These signals are based for "x86, arm, and most other architectures"
    // see "man 7 signal" on Linux or "man signal" on BSD
    case "129":  return "SIGHUP"  //  1
    case "130":  return "SIGINT"  //  2
    case "131":  return "SIGQUIT" //  3
    case "132":  return "SIGILL"  //  4
    case "133":  return "SIGTRAP"  // 5
    case "134":  return "SIGABRT" //  6
    case "135":  return "SIGBUS"  //  7
    case "136":  return "SIGFPE"  //  8
    case "137":  return "SIGKILL" //  9
    case "138":  return "SIGUSR1" // 10
    case "139":  return "SIGSEGV" // 11
    case "140":  return "SIGUSR2" // 12
    case "141":  return "SIGPIPE" // 13
    case "142":  return "SIGALRM" // 14
    case "143":  return "SIGTERM" // 15
    case "145":  return "SIGCHLD" // 17
    case "146":  return "SIGCONT" // 18
    case "147":  return "SIGSTOP" // 19
    case "148":  return "SIGTSTP" // 20
    case "149":  return "SIGTTIN" // 21
    case "150":  return "SIGTTOU" // 22

    // anything else is unknown
    default: return "UNK"
    }
}

// Checks for the return code.
func getStatus(c *Car) (isOk bool, msg string) {
    _, argsExist := c.Params["args"]

    if ! argsExist || c.Params["args"] == "0" {
        isOk = true
        msg = ""
    } else {
        isOk = false
        msg = getMsgMapping(c.Params["args"].(string))
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
    defaultCodeBg := defaultRootBg
    defaultCodeFg := defaultRootFg
    defaultCodeFm := defaultRootFm
    defaultMsgBg := defaultRootBg
    defaultMsgFg := defaultRootFg
    defaultMsgFm := defaultRootFm

    defaultSymbolFormat := "{{ Error }}"
    defaultCodeText := "?"

    if val, ok := c.Params["args"]; ok {
        defaultCodeText = val.(string)
    }

    isOk, msg := getStatus(c)
    if isOk {
        defaultRootBg = utils.GetEnv("GBT_CAR_BG", defaultOkBg)
        defaultRootFg = utils.GetEnv("GBT_CAR_FG", defaultOkFg)
        defaultRootFm = utils.GetEnv("GBT_CAR_FM", defaultOkFm)
        defaultSymbolFormat = "{{ Ok }}"
    }

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_STATUS_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_STATUS_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_STATUS_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_STATUS_FORMAT", " {{ Symbol }} "),
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
        "Msg": {
            Bg: utils.GetEnv(
                "GBT_CAR_STATUS_MSG_BG", utils.GetEnv(
                    "GBT_CAR_STATUS_BG", defaultMsgBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_STATUS_MSG_FG", utils.GetEnv(
                    "GBT_CAR_STATUS_FG", defaultMsgFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_STATUS_MSG_FM", utils.GetEnv(
                    "GBT_CAR_STATUS_FM", defaultMsgFm)),
            Text: utils.GetEnv(
                "GBT_CAR_STATUS_MSG_TEXT", msg),
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
    }

    if isOk {
        c.Display = utils.GetEnvBool("GBT_CAR_STATUS_DISPLAY", false)
    } else {
        c.Display = utils.GetEnvBool("GBT_CAR_STATUS_DISPLAY", true)
    }

    c.Wrap = utils.GetEnvBool("GBT_CAR_STATUS_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_STATUS_SEP", "\000")
}
