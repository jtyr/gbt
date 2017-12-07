package git

import (
    "strings"

    "github.com/jtyr/gbt/gbt/core/car"
    "github.com/jtyr/gbt/gbt/core/utils"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

// Returns true if the current directory is a Git repo.
func isGitDir() bool {
    rc, _, _ := utils.Run([]string{"git", "rev-parse", "--git-dir"})

    return rc == 0
}

// Returns name of the head.
func getHead(display bool) string {
    if ! display {
        return ""
    }

    // Remote branch name which the local branch is tracking
    rc, out, _ := utils.Run([]string{"git", "symbolic-ref", "HEAD"})

    if rc > 0 {
        // Get tag name
        rc, out, _ = utils.Run(
            []string{"git", "describe", "--tags", "--exact-match", "HEAD"})

        if rc > 0 {
            // Get commit ID
            _, out, _ = utils.Run([]string{"git", "rev-parse", "--short", "HEAD"})
        }
    }

    return strings.Replace(out, "refs/heads/", "", 1)
}

// Returns true if the repo si dirty.
func isDirty(display bool) bool {
    if ! display {
        return false
    }

    _, out, _ := utils.Run([]string{"git", "status", "--porcelain"})

    return len(out) > 0
}

// Returns true if the repo si ahead/behind.
func compareRemote(display bool, ahead bool) bool {
    if ! display {
        return false
    }

    ret := false

    direction := "HEAD..@{upstream}"

    if ahead {
        direction = "@{upstream}..HEAD"
    }

    rc, out, _ := utils.Run([]string{"git", "rev-list", "--count", direction})

    if rc == 0 && out != "0" {
        ret = true
    }

    return ret
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "light_gray")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "black")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultIconBg := defaultRootBg
    defaultIconFg := defaultRootFg
    defaultIconFm := defaultRootFm
    defaultHeadBg := defaultRootBg
    defaultHeadFg := defaultRootFg
    defaultHeadFm := defaultRootFm
    defaultStatusBg := defaultRootBg
    defaultStatusFg := defaultRootFg
    defaultStatusFm := defaultRootFm
    defaultDirtyBg := defaultRootBg
    defaultDirtyFg := "red"
    defaultDirtyFm := defaultRootFm
    defaultCleanBg := defaultRootBg
    defaultCleanFg := "green"
    defaultCleanFm := defaultRootFm
    defaultAheadBg := defaultRootBg
    defaultAheadFg := defaultRootFg
    defaultAheadFm := defaultRootFm
    defaultBehindBg := defaultRootBg
    defaultBehindFg := defaultRootFg
    defaultBehindFm := defaultRootFm

    c.Display = utils.GetEnvBool("GBT_CAR_GIT_DISPLAY", isGitDir())

    defaultStatusFormat := "{{ Clean }}"
    defaultAheadText := ""
    defaultBehindText := ""

    if isDirty(c.Display) {
        defaultStatusFormat = "{{ Dirty }}"
    }

    if compareRemote(c.Display, true) {
        defaultAheadText = utils.GetEnv("GBT_CAR_GIT_AHEAD_SYMBOL", " ⬆")
    }

    if compareRemote(c.Display, false) {
        defaultBehindText = utils.GetEnv("GBT_CAR_GIT_BEHIND_SYMBOL", " ⬇")
    }

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_GIT_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_GIT_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_GIT_FM", defaultRootFm),
            Text: utils.GetEnv(
                "GBT_CAR_GIT_FORMAT",
                " {{ Icon }} {{ Head }} {{ Status }}{{ Ahead }}{{ Behind }} "),
        },
        "Icon": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_ICON_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultIconBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_ICON_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultIconFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_ICON_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultIconFm)),
            Text: utils.GetEnv("GBT_CAR_GIT_ICON_TEXT", ""),
        },
        "Head": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_HEAD_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultHeadBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_HEAD_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultHeadFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_HEAD_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultHeadFm)),
            Text: utils.GetEnv("GBT_CAR_GIT_HEAD_TEXT", getHead(c.Display)),
        },
        "Status": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultStatusBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultStatusFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultStatusFm)),
            Text: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_FORMAT", defaultStatusFormat),
        },
        "Dirty": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_DIRTY_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultDirtyBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_DIRTY_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultDirtyFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_DIRTY_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultDirtyFm)),
            Text: utils.GetEnv("GBT_CAR_GIT_DIRTY_TEXT", "✘"),
        },
        "Clean": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_CLEAN_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultCleanBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_CLEAN_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultCleanFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_CLEAN_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultCleanFm)),
            Text: utils.GetEnv("GBT_CAR_GIT_CLEAN_TEXT", "✔"),
        },
        "Ahead": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultAheadBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultAheadFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultAheadFm)),
            Text: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_TEXT", defaultAheadText),
        },
        "Behind": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultBehindBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultBehindFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultBehindFm)),
            Text: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_TEXT", defaultBehindText),
        },
    }

    c.Wrap = utils.GetEnvBool("GBT_CAR_GIT_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_GIT_SEP", "\000")
}
