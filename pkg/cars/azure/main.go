package azure

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "os/user"
    "strings"
    "bytes"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"

    "gopkg.in/go-ini/ini.v1"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

// Top level of the azureProfile.json file
type azureProfile struct {
    Subscriptions []azureProfileSubscription `json:"subscriptions"`
}

// Subscription in the azureProfile.json file
type azureProfileSubscription struct {
    ID    string                       `json:"id"`
    Name  string                       `json:"name"`
    State string                       `json:"state"`
    User  azureProfileSubscriptionUser `json:"user"`
}

// User of the specific Subscription in the azureProfile.json file
type azureProfileSubscriptionUser struct {
    Name string `json:"name"`
    Type string `json:"type"`
}

// OS-specific path separator.
const osSep = string(os.PathSeparator)

// To be able to fake the user's home directory.
var usr, _ = user.Current()

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "32")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "white")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultSep := "\000"

    cloud := os.Getenv("AZURE_CLOUD_NAME")
    subscription := ""
    userName := ""
    userType := ""
    state := ""
    defaultsGroup := os.Getenv("AZURE_DEFAULTS_GROUP")

    c.Display = utils.GetEnvBool("GBT_CAR_AZURE_DISPLAY", true)

    if c.Display {
        confDir := os.Getenv("AZURE_CONFIG_DIR")

        if confDir == "" {
            userHome := usr.HomeDir
            confDir = strings.Join([]string{userHome, ".azure"}, osSep)
        }

        if cloud == "" {
            configFile := strings.Join([]string{confDir, "config"}, osSep)
            cfg, err := ini.Load(configFile)

            if err == nil {
                // Get the current Cloud Name
                if cloudSection, sErr := cfg.GetSection("cloud"); sErr == nil {
                    if cName, kErr := cloudSection.GetKey("name"); kErr == nil {
                        cloud = cName.String()
                    }
                }

                // Get the default Resource Group
                if defaultsGroup == "" {
                    if cloudSection, sErr := cfg.GetSection("defaults"); sErr == nil {
                        if cName, kErr := cloudSection.GetKey("group"); kErr == nil {
                            defaultsGroup = cName.String()
                        }
                    }
                }
            }

            // Default Cloud Name
            if cloud == "" {
                cloud = "AzureCloud"
            }
        }

        // Get the Subscription ID
        if cloud != "" {
            subscrID := ""
            cloudsConfigFile := strings.Join([]string{confDir, "clouds.config"}, osSep)

            cfg, err := ini.Load(cloudsConfigFile)

            if err == nil {
                if cloudSection, sErr := cfg.GetSection(cloud); sErr == nil {
                    if sID, kErr := cloudSection.GetKey("subscription"); kErr == nil {
                        subscrID = sID.String()
                    }
                }
            }

            // Get the Subscription Name, User Name, User Type and State
            if subscrID != "" {
                azureProfileFile := strings.Join([]string{confDir, "azureProfile.json"}, osSep)

                if byteValue, err := ioutil.ReadFile(azureProfileFile); err == nil {
                    // There is some garbage at the beginning of the file!
                    if bytes.Compare(byteValue[:3], []byte{'\xEF', '\xBB', '\xBF'}) == 0 {
                        byteValue = byteValue[3:]
                    }

                    var data azureProfile
                    json.Unmarshal(byteValue, &data)

                    for _, s := range data.Subscriptions {
                        if s.ID == subscrID {
                            subscription = s.Name
                            userName = s.User.Name
                            userType = s.User.Type
                            state = s.State

                            break
                        }
                    }
                }
            }
        }
    }

    c.Model = map[string]car.ModelElement{
        "root": {
            Bg:   utils.GetEnv("GBT_CAR_AZURE_BG", defaultRootBg),
            Fg:   utils.GetEnv("GBT_CAR_AZURE_FG", defaultRootFg),
            Fm:   utils.GetEnv("GBT_CAR_AZURE_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_AZURE_FORMAT", " {{ Icon }} {{ Subscription }} "),
        },
        "Icon": {
            Bg: utils.GetEnv(
                "GBT_CAR_AZURE_ICON_BG", utils.GetEnv(
                    "GBT_CAR_AZURE_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_AZURE_ICON_FG", utils.GetEnv(
                    "GBT_CAR_AZURE_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_AZURE_ICON_FM", utils.GetEnv(
                    "GBT_CAR_AZURE_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_AZURE_ICON_TEXT", "\ufd03"),
        },
        "Cloud": {
            Bg: utils.GetEnv(
                "GBT_CAR_AZURE_CLOUD_BG", utils.GetEnv(
                    "GBT_CAR_AZURE_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_AZURE_CLOUD_FG", utils.GetEnv(
                    "GBT_CAR_AZURE_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_AZURE_CLOUD_FM", utils.GetEnv(
                    "GBT_CAR_AZURE_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_AZURE_CLOUD_TEXT", cloud),
        },
        "Subscription": {
            Bg: utils.GetEnv(
                "GBT_CAR_AZURE_SUBSCRIPTION_BG", utils.GetEnv(
                    "GBT_CAR_AZURE_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_AZURE_SUBSCRIPTION_FG", utils.GetEnv(
                    "GBT_CAR_AZURE_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_AZURE_SUBSCRIPTION_FM", utils.GetEnv(
                    "GBT_CAR_AZURE_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_AZURE_SUBSCRIPTION_TEXT", subscription),
        },
        "UserName": {
            Bg: utils.GetEnv(
                "GBT_CAR_AZURE_USERNAME_BG", utils.GetEnv(
                    "GBT_CAR_AZURE_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_AZURE_USERNAME_FG", utils.GetEnv(
                    "GBT_CAR_AZURE_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_AZURE_USERNAME_FM", utils.GetEnv(
                    "GBT_CAR_AZURE_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_AZURE_USERNAME_TEXT", userName),
        },
        "UserType": {
            Bg: utils.GetEnv(
                "GBT_CAR_AZURE_USERTYPE_BG", utils.GetEnv(
                    "GBT_CAR_AZURE_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_AZURE_USERTYPE_FG", utils.GetEnv(
                    "GBT_CAR_AZURE_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_AZURE_USERTYPE_FM", utils.GetEnv(
                    "GBT_CAR_AZURE_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_AZURE_USERTYPE_TEXT", userType),
        },
        "State": {
            Bg: utils.GetEnv(
                "GBT_CAR_AZURE_STATE_BG", utils.GetEnv(
                    "GBT_CAR_AZURE_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_AZURE_STATE_FG", utils.GetEnv(
                    "GBT_CAR_AZURE_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_AZURE_STATE_FM", utils.GetEnv(
                    "GBT_CAR_AZURE_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_AZURE_STATE_TEXT", state),
        },
        "DefaultsGroup": {
            Bg: utils.GetEnv(
                "GBT_CAR_AZURE_DEFAULTS_GROUP_BG", utils.GetEnv(
                    "GBT_CAR_AZURE_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_AZURE_DEFAULTS_GROUP_FG", utils.GetEnv(
                    "GBT_CAR_AZURE_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_AZURE_DEFAULTS_GROUP_FM", utils.GetEnv(
                    "GBT_CAR_AZURE_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_AZURE_DEFAULTS_GROUP_TEXT", defaultsGroup),
        },
        "Sep": {
            Bg: utils.GetEnv(
                "GBT_CAR_AZURE_SEP_BG", utils.GetEnv(
                    "GBT_SEPARATOR_BG", defaultSep)),
            Fg: utils.GetEnv(
                "GBT_CAR_AZURE_SEP_FG", utils.GetEnv(
                    "GBT_SEPARATOR_FG", defaultSep)),
            Fm: utils.GetEnv(
                "GBT_CAR_AZURE_SEP_FM", utils.GetEnv(
                    "GBT_SEPARATOR_FM", defaultSep)),
            Text: utils.GetEnv(
                "GBT_CAR_AZURE_SEP", utils.GetEnv(
                    "GBT_CAR_AZURE_SEP_TEXT", utils.GetEnv(
                        "GBT_SEPARATOR", defaultSep))),
        },
    }

    c.Wrap = utils.GetEnvBool("GBT_CAR_AZURE_WRAP", false)
}
