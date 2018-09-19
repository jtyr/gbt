package docker

import (
    "os"
    "strings"

    "path/filepath"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

var runDockerVersion = []string{"docker", "version", "-f", "{{.Server.Version}}"}
var defaultIcon = "\uf308"

func getDockerVersion() string {
    rc, output, _ := utils.Run(runDockerVersion)

    if rc == 0 {
        return strings.TrimSpace(output)
    }

    return ""
}

func getDockerComposeFiles() []string {
    files := []string{}
    composeFilesStr := utils.GetEnv("COMPOSE_FILE", "")
    if composeFilesStr != "" {
        files = strings.Split(composeFilesStr, string(os.PathListSeparator))
    }

    return append(append(files, "Dockerfile"), "docker-compose.yml")
}

func shouldDisplayDocker(dockerCmd string, alwaysShow bool) bool {
    if !utils.CommandExists(dockerCmd) {
        return false
    }

    if alwaysShow {
        return true
    }

    composeFiles := getDockerComposeFiles()

    wd, _ := os.Getwd()

    wd = filepath.Clean(wd)
    for {
        wdStat, err := os.Stat(wd)
        if os.IsNotExist(err) || !wdStat.IsDir() {
            return false
        }

        for _, fileName := range composeFiles {
            dfPath := filepath.Join(wd, fileName)
            dockerFileStat, err := os.Stat(dfPath)
            exists := !os.IsNotExist(err)
            isNotDir := exists && !dockerFileStat.IsDir()
            if exists && isNotDir {
                return true
            }
        }

        wd = filepath.Dir(wd)
        if wd == "/" {
            break
        }
    }

    return false
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "26")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "white")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultIconBg := defaultRootBg
    defaultIconFg := defaultRootFg
    defaultIconFm := defaultRootFm
    defaultVersionBg := defaultRootBg
    defaultVersionFg := defaultRootFg
    defaultVersionFm := defaultRootFm

    dockerCmd := utils.GetEnv("GBT_CAR_DOCKER_CMD", "docker")
    alwaysShow := utils.GetEnvBool("GBT_CAR_DOCKER_ALWAYS", false)

    c.Display = utils.GetEnvBool("GBT_CAR_DOCKER_DISPLAY", shouldDisplayDocker(dockerCmd, alwaysShow))

    c.Model = map[string]car.ModelElement{
        "root": {
            Bg:   utils.GetEnv("GBT_CAR_DOCKER_BG", defaultRootBg),
            Fg:   utils.GetEnv("GBT_CAR_DOCKER_FG", defaultRootFg),
            Fm:   utils.GetEnv("GBT_CAR_DOCKER_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_DOCKER_FORMAT", " {{ Icon }} {{ Version }} "),
        },
        "Icon": {
            Bg: utils.GetEnv(
                "GBT_CAR_DOCKER_ICON_BG", utils.GetEnv(
                    "GBT_CAR_DOCKER_BG", defaultIconBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_DOCKER_ICON_FG", utils.GetEnv(
                    "GBT_CAR_DOCKER_FG", defaultIconFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_DOCKER_ICON_FM", utils.GetEnv(
                    "GBT_CAR_DOCKER_FM", defaultIconFm)),
            Text: utils.GetEnv("GBT_CAR_DOCKER_ICON_TEXT", defaultIcon),
        },
        "Version": {
            Bg: utils.GetEnv(
                "GBT_CAR_DOCKER_VERSION_BG", utils.GetEnv(
                    "GBT_CAR_DOCKER_BG", defaultVersionBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_DOCKER_VERSION_FG", utils.GetEnv(
                    "GBT_CAR_DOCKER_FG", defaultVersionFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_DOCKER_VERSION_FM", utils.GetEnv(
                    "GBT_CAR_DOCKER_FM", defaultVersionFm)),
            Text: utils.GetEnv("GBT_CAR_DOCKER_VERSION_TEXT", getDockerVersion()),
        },
    }

    c.Wrap = utils.GetEnvBool("GBT_CAR_DOCKER_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_DOCKER_SEP", "\000")
}
