package os

import (
    "bufio"
    "os"
    "runtime"
    "strings"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"
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
    // Unicode codes and font names are taken from https://nerdfonts.com
    // If adding a new symbol, always choose the smaller picture if multiple
    // symbols are available.
    "amzn":       { icon: "\uf52c", color: "208", }, // nf-mdi-amazon
    "android":    { icon: "\uf17b", color: "113", }, // nf-fa-android
    "arch":       { icon: "\uf303", color: "25",  }, // nf-linux-archlinux
    "archarm":    { icon: "\uf303", color: "125", }, // nf-linux-archlinux
    "alpine":     { icon: "\uf300", color: "24",  }, // nf-linux-alpine
    "aosc":       { icon: "\uf301", color: "172", }, // nf-linux-aosc
    "centos":     { icon: "\uf304", color: "27",  }, // nf-linux-centos
    "cloud":      { icon: "\uf65e", color: "39",  }, // nf-mdi-cloud
    "coreos":     { icon: "\uf305", color: "32",  }, // nf-linux-coreos
    "darwin":     { icon: "\uf534", color: "15",  }, // nf-mdi-apple
    "debian":     { icon: "\ue77d", color: "88",  }, // nf-dev-debian
    "devuan":     { icon: "\uf307", color: "16",  }, // nf-linux-devuan
    "docker":     { icon: "\ue7b0", color: "26",  }, // nf-dev-docker
    "elementary": { icon: "\uf309", color: "33",  }, // nf-linux-elementary
    "fedora":     { icon: "\uf30a", color: "32",  }, // nf-linux-fedora
    "freebsd":    { icon: "\uf30c", color: "1",   }, // nf-linux-freebsd
    "gentoo":     { icon: "\uf30d", color: "62",  }, // nf-linux-gentoo
    "linux":      { icon: "\uf17c", color: "15",  }, // nf-fa-linux
    "linuxmint":  { icon: "\uf30e", color: "47",  }, // nf-linux-linuxmint
    "mageia":     { icon: "\uf310", color: "24",  }, // nf-linux-mageia
    "mandriva":   { icon: "\uf311", color: "208", }, // nf-linux-mandriva
    "manjaro":    { icon: "\uf312", color: "34",  }, // nf-linux-manjaro
    "mysql":      { icon: "\ue704", color: "30",  }, // nf-dev-mysql
    "nixos":      { icon: "\uf313", color: "88",  }, // nf-linux-nixos
    "opensuse":   { icon: "\uf314", color: "113", }, // nf-linux-opensuse
    "raspbian":   { icon: "\uf315", color: "125", }, // nf-linux-raspberry_pi
    "redhat":     { icon: "\ue7bb", color: "1",   }, // nf-dev-redhat
    "sabayon":    { icon: "\uf317", color: "255", }, // nf-linux-sabayon
    "slackware":  { icon: "\uf318", color: "63",  }, // nf-linux-slackware
    "ubuntu":     { icon: "\uf31b", color: "166", }, // nf-linux-ubuntu
    "windows":    { icon: "\ue62a", color: "6",   }, // nf-custom-windows
}

// Holds the OS name.
var osName string

// Path to the os-release file.
var osReleaseFile = "/etc/os-release"

// Returns the OS name.
func getOsName() string {
    if osName != "" {
        return osName
    }

    osName = runtime.GOOS

    if _, err := os.Stat(osReleaseFile); ! os.IsNotExist(err) {
        file, err := os.Open(osReleaseFile)

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
                    "GBT_CAR_OS_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_OS_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_OS_FG", getOsColor())),
            Fm: utils.GetEnv(
                "GBT_CAR_OS_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_OS_FM", defaultRootFm)),
            Text: utils.GetEnv(
                "GBT_CAR_OS_SYMBOL_TEXT", getOsSymbol()),
        },
    }

    c.Display = utils.GetEnvBool("GBT_CAR_OS_DISPLAY", true)
    c.Wrap = utils.GetEnvBool("GBT_CAR_OS_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_OS_SEP", "\000")
}
