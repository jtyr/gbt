package os

import (
    "bufio"
    "os"
    "runtime"
    "strings"

    "github.com/jtyr/gbt/gbt/core/car"
    "github.com/jtyr/gbt/gbt/core/utils"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

// Type for the symbols icon and color.
type iconColor struct {
    icon string
    color string
}

// List of names and symbols.
var symbols = map[string]iconColor {
    "amzn":       { icon: "", color: "208",   },
    "android":    { icon: "", color: "113",   },
    "arch":       { icon: "", color: "25",    },
    "archarm":    { icon: "", color: "125",   },
    "centos":     { icon: "", color: "27",    },
    "cloud":      { icon: "", color: "39",    },
    "coreos":     { icon: "", color: "white", },
    "darwin":     { icon: "", color: "white", },
    "debian":     { icon: "", color: "88",    },
    "docker":     { icon: "", color: "26",    },
    "elementary": { icon: "", color: "33",    },
    "fedora":     { icon: "", color: "32",    },
    "freebsd":    { icon: "", color: "red",   },
    "gentoo":     { icon: "", color: "62"     },
    "linux":      { icon: "", color: "white", },
    "linuxmint":  { icon: "", color: "47",    },
    "mageia":     { icon: "", color: "24",    },
    "mandriva":   { icon: "", color: "208",   },
    "opensuse":   { icon: "", color: "113",   },
    "raspbian":   { icon: "", color: "125",   },
    "redhat":     { icon: "", color: "red",   },
    "sabayon":    { icon: "", color: "white", },
    "slackware":  { icon: "", color: "white", },
    "ubuntu":     { icon: "", color: "166",   },
    "windows":    { icon: "", color: "cyan",  },
}

// Holds the OS name.
var osName string

// Returns the OS name.
func getOsName() string {
    if osName != "" {
        return osName
    }

    osName = runtime.GOOS

    if _, err := os.Stat("/etc/os-release"); ! os.IsNotExist(err) {
        file, err := os.Open("/etc/os-release")

        if err != nil {
            return osName
        }

        defer file.Close()

        scanner := bufio.NewScanner(file)
        scanner.Split(bufio.ScanLines)

        for scanner.Scan() {
            line := scanner.Text()

            if len(line) > 3 && line[:3] == "ID=" {
                osName = strings.Replace(
                    strings.Replace(line[3:], "\"", "", -1), "'", "", -1)
            }
        }
    }

    return osName
}

// Returns the OS symbol.
func getOsSymbol() (ret string) {
    name := strings.ToLower(utils.GetEnv("GBT_CAR_OS_NAME", getOsName()))

    if val, ok := symbols[name]; ok {
        ret = val.icon
    } else {
        ret = "?"
    }

    return
}

// Returns the OS color.
func getOsColor() (ret string) {
    name := strings.ToLower(utils.GetEnv("GBT_CAR_OS_NAME", getOsName()))

    if val, ok := symbols[name]; ok {
        ret = val.color
    } else {
        ret = "white"
    }

    return
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "235")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "white")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultSymbolBg := defaultRootBg
    defaultSymbolFg := getOsColor()
    defaultSymbolFm := defaultRootFm

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_OS_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_OS_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_OS_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_OS_FORMAT", " {{ Symbol }} "),
        },
        "Symbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_OS_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_OS_BG", defaultSymbolBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_OS_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_OS_FG", defaultSymbolFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_OS_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_OS_FM", defaultSymbolFm)),
            Text: utils.GetEnv(
                "GBT_CAR_OS_SYMBOL_TEXT", getOsSymbol()),
        },
    }

    c.Display = utils.GetEnvBool("GBT_CAR_OS_DISPLAY", true)
    c.Wrap = utils.GetEnvBool("GBT_CAR_OS_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_OS_SEP", "\000")
}
