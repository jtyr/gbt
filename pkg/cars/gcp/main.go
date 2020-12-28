package dir

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "os/user"
    "runtime"
    "strings"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"

    "gopkg.in/go-ini/ini.v1"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

// OS-specific path separator
const osSep = string(os.PathSeparator)

// To be able to test Windows results.
var runtimeOs = runtime.GOOS

// To be able to fake the user's home directory
var usr, _ = user.Current()

// getDefaultConfDir returns location of the configuration directory used by gcloud.
func getDefaultConfDir() (defaultDir string) {
    baseDir := ""

    if runtimeOs == "windows" {
        baseDir = utils.GetEnv("APPDATA", utils.GetEnv("SystemDrive", "C:"))
    } else {
        userHome := usr.HomeDir
        baseDir = strings.Join([]string{userHome, ".config"}, osSep)
    }

    defaultDir = strings.Join([]string{baseDir, "gcloud"}, osSep)

    return
}

// getActiveConfig returns active config.
func getActiveConfig(configDir string) (config string) {
    dat, err := ioutil.ReadFile(strings.Join([]string{configDir, "active_config"}, osSep))

    if err == nil {
        config = strings.TrimSpace(string(dat))
    }

    return
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "66;133;244")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "white")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultSep := "\000"

    config := ""
    account := os.Getenv("CLOUDSDK_CORE_ACCOUNT")
    project := os.Getenv("CLOUDSDK_CORE_PROJECT")

    if os.Getenv("CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE") != "" {
        jsonFile := os.Getenv("CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE")
        byteValue, err := ioutil.ReadFile(jsonFile)

        if err == nil {
            var data map[string]interface{}
            json.Unmarshal([]byte(byteValue), &data)

            if val, ok := data["client_email"]; ok {
                account = val.(string)
            }
        }
    }

    configDir := utils.GetEnv("CLOUDSDK_CONFIG", getDefaultConfDir())
    config = getActiveConfig(configDir)

    configFile := strings.Join([]string{configDir, "configurations", fmt.Sprintf("config_%s", config)}, osSep)
    cfg, err := ini.Load(configFile)

    if err == nil {
        if account == "" {
            account = cfg.Section("core").Key("account").String()
        }

        if project == "" {
            project = cfg.Section("core").Key("project").String()
        }
    }

    if os.Getenv("GBT_CAR_GCP_PROJECT_ALIASES") != "" {
        for _, pair := range strings.Split(os.Getenv("GBT_CAR_GCP_PROJECT_ALIASES"), ",") {
            kv := strings.Split(pair, "=")

            if len(kv) == 2 && project == strings.TrimSpace(kv[0]) {
                project = strings.TrimSpace(kv[1])

                break
            }
        }
    }

    c.Model = map[string]car.ModelElement{
        "root": {
            Bg:   utils.GetEnv("GBT_CAR_GCP_BG", defaultRootBg),
            Fg:   utils.GetEnv("GBT_CAR_GCP_FG", defaultRootFg),
            Fm:   utils.GetEnv("GBT_CAR_GCP_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_GCP_FORMAT", " {{ Icon }} {{ Project }} "),
        },
        "Icon": {
            Bg: utils.GetEnv(
                "GBT_CAR_GCP_ICON_BG", utils.GetEnv(
                    "GBT_CAR_GCP_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GCP_ICON_FG", utils.GetEnv(
                    "GBT_CAR_GCP_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GCP_ICON_FM", utils.GetEnv(
                    "GBT_CAR_GCP_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_GCP_ICON_TEXT", "\ue7b2"),
        },
        "Config": {
            Bg: utils.GetEnv(
                "GBT_CAR_GCP_CONFIG_BG", utils.GetEnv(
                    "GBT_CAR_GCP_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GCP_CONFIG_FG", utils.GetEnv(
                    "GBT_CAR_GCP_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GCP_CONFIG_FM", utils.GetEnv(
                    "GBT_CAR_GCP_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_GCP_CONFIG_TEXT", config),
        },
        "Account": {
            Bg: utils.GetEnv(
                "GBT_CAR_GCP_ACCOUNT_BG", utils.GetEnv(
                    "GBT_CAR_GCP_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GCP_ACCOUNT_FG", utils.GetEnv(
                    "GBT_CAR_GCP_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GCP_ACCOUNT_FM", utils.GetEnv(
                    "GBT_CAR_GCP_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_GCP_ACCOUNT_TEXT", account),
        },
        "Project": {
            Bg: utils.GetEnv(
                "GBT_CAR_GCP_PROJECT_BG", utils.GetEnv(
                    "GBT_CAR_GCP_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GCP_PROJECT_FG", utils.GetEnv(
                    "GBT_CAR_GCP_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GCP_PROJECT_FM", utils.GetEnv(
                    "GBT_CAR_GCP_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_GCP_PROJECT_TEXT", project),
        },
        "Sep": {
            Bg: utils.GetEnv(
                "GBT_CAR_GCP_SEP_BG", utils.GetEnv(
                    "GBT_SEPARATOR_BG", defaultSep)),
            Fg: utils.GetEnv(
                "GBT_CAR_GCP_SEP_FG", utils.GetEnv(
                    "GBT_SEPARATOR_FG", defaultSep)),
            Fm: utils.GetEnv(
                "GBT_CAR_GCP_SEP_FM", utils.GetEnv(
                    "GBT_SEPARATOR_FM", defaultSep)),
            Text: utils.GetEnv(
                "GBT_CAR_GCP_SEP", utils.GetEnv(
                    "GBT_CAR_GCP_SEP_TEXT", utils.GetEnv(
                        "GBT_SEPARATOR", defaultSep))),
        },
    }

    c.Display = utils.GetEnvBool("GBT_CAR_GCP_DISPLAY", true)
    c.Wrap = utils.GetEnvBool("GBT_CAR_GCP_WRAP", false)
}
