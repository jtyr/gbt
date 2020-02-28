package git

import (
    "regexp"
    "strconv"
    "strings"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

type statusCount struct {
    added     int
    copied    int
    deleted   int
    ignored   int
    modified  int
    renamed   int
    staged    int
    unmerged  int
    untracked int
}

// Git commands.
var runIsGitDir = []string{"git", "rev-parse", "--git-dir"}
var runGetBranch = []string{"git", "symbolic-ref", "HEAD"}
var runGetTag = []string{"git", "describe", "--tags", "--exact-match", "HEAD"}
var runGetCommit = []string{"git", "rev-parse", "--short", "HEAD"}
var runIsDirty = []string{"git", "status", "--porcelain"}
var runCompareRemote = []string{"git", "rev-list", "--count"}
var runStash = []string{"git", "stash", "list"}

// Regexps for matching the templating
var reTemplatingHead = regexp.MustCompile(`{{\s*Head\s*}}`)
var reTemplatingStatus = regexp.MustCompile(`{{\s*Status.*\s*}}`)
var reTemplatingAhead = regexp.MustCompile(`{{\s*Ahead.*\s*}}`)
var reTemplatingBehind = regexp.MustCompile(`{{\s*Behind.*\s*}}`)
var reTemplatingStash = regexp.MustCompile(`{{\s*Stash.*\s*}}`)

var defaultRootFormat = utils.GetEnv("GBT_CAR_GIT_FORMAT", " {{ Icon }} {{ Head }} {{ Status }}{{ Ahead }}{{ Behind }} ")

// Returns true if the current directory is a Git repo.
func isGitDir() bool {
    rc, _, _ := utils.Run(runIsGitDir)

    return rc == 0
}

// Returns name of the head.
func getHead() string {
    // Remote branch name which the local branch is tracking
    rc, out, _ := utils.Run(runGetBranch)

    if rc > 0 {
        // Get tag name
        rc, out, _ = utils.Run(runGetTag)

        if rc > 0 {
            // Get commit ID
            _, out, _ = utils.Run(runGetCommit)
        }
    }

    return strings.Replace(out, "refs/heads/", "", 1)
}

// Returns true and the status field counts if the repo is dirty.
func isDirty() (bool, statusCount) {
    rc, out, _ := utils.Run(runIsDirty, false)

    dirty := false
    var status statusCount

    if rc == 0 && len(out) > 0 {
        dirty = true

        // Normalize new lines (is it needed?)
        out = strings.Replace(out, "\r\n", "\n", -1)
        out = strings.Replace(out, "\r", "\n", -1)

        for _, line := range strings.Split(out, "\n") {
            if len(line) > 4 {
                switch line[1] {
                case 'A':
                    status.added++
                case 'C':
                    status.copied++
                case 'D':
                    status.deleted++
                case '!':
                    status.ignored++
                case 'M':
                    status.modified++
                case 'R':
                    status.renamed++
                case ' ':
                    status.staged++
                case 'U':
                    status.unmerged++
                case '?':
                    status.untracked++
                }
            }
        }
    }

    return dirty, status
}

// Returns number of commits the repo is ahead/behind.
func compareRemote(ahead bool) string {
    ret := "0"

    direction := "HEAD..@{upstream}"

    if ahead {
        direction = "@{upstream}..HEAD"
    }

    rc, out, _ := utils.Run(append(runCompareRemote, direction))

    if rc == 0 && out != "0" {
        ret = strings.Split(out, " ")[0]
    }

    return ret
}

// Returns number of stashes.
func getStash() string {
    ret := "0"

    rc, out, _ := utils.Run(runStash)

    if rc == 0 && len(out) > 0 {
        ret = strconv.Itoa(len(strings.Split(out, "\n")))
    }

    return ret
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "light_gray")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "black")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultSep := "\000"

    c.Display = utils.GetEnvBool("GBT_CAR_GIT_DISPLAY", isGitDir())

    defaultHeadText := ""
    defaultStatusFormat := "{{ StatusClean }}"
    defaultStatusAddedFormat := ""
    defaultStatusAddedSymbolText := ""
    defaultStatusAddedCountText := ""
    defaultStatusCopiedFormat := ""
    defaultStatusCopiedSymbolText := ""
    defaultStatusCopiedCountText := ""
    defaultStatusDeletedFormat := ""
    defaultStatusDeletedSymbolText := ""
    defaultStatusDeletedCountText := ""
    defaultStatusIgnoredFormat := ""
    defaultStatusIgnoredSymbolText := ""
    defaultStatusIgnoredCountText := ""
    defaultStatusModifiedFormat := ""
    defaultStatusModifiedSymbolText := ""
    defaultStatusModifiedCountText := ""
    defaultStatusRenamedFormat := ""
    defaultStatusRenamedSymbolText := ""
    defaultStatusRenamedCountText := ""
    defaultStatusStagedFormat := ""
    defaultStatusStagedSymbolText := ""
    defaultStatusStagedCountText := ""
    defaultStatusUnmergedFormat := ""
    defaultStatusUnmergedSymbolText := ""
    defaultStatusUnmergedCountText := ""
    defaultStatusUntrackedFormat := ""
    defaultStatusUntrackedSymbolText := ""
    defaultStatusUntrackedCountText := ""
    defaultAheadFormat := ""
    defaultAheadSymbolText := ""
    defaultAheadCountText := ""
    defaultBehindFormat := ""
    defaultBehindSymbolText := ""
    defaultBehindCountText := ""
    defaultStashFormat := ""
    defaultStashSymbolText := ""
    defaultStashCountText := ""

    if c.Display {
        if ok := reTemplatingHead.MatchString(defaultRootFormat); ok {
            defaultHeadText = getHead()
        }

        if ok := reTemplatingStatus.MatchString(defaultRootFormat); ok {
            dirty, status := isDirty()

            if dirty {
                defaultStatusFormat = "{{ StatusDirty }}"

                if status.added > 0 {
                    defaultStatusAddedFormat = utils.GetEnv("GBT_CAR_GIT_STATUS_ADDED_FORMAT", "{{ StatusAddedSymbol }}")
                    defaultStatusAddedSymbolText = utils.GetEnv("GBT_CAR_GIT_STATUS_ADDED_SYMBOL_TEXT", " \u27f4")
                    defaultStatusAddedCountText = utils.GetEnv("GBT_CAR_GIT_STATUS_ADDED_COUNT_TEXT", strconv.Itoa(status.added))
                }

                if status.copied > 0 {
                    defaultStatusCopiedFormat = utils.GetEnv("GBT_CAR_GIT_STATUS_COPIED_FORMAT", "{{ StatusCopiedSymbol }}")
                    defaultStatusCopiedSymbolText = utils.GetEnv("GBT_CAR_GIT_STATUS_COPIED_SYMBOL_TEXT", " \u2948")
                    defaultStatusCopiedCountText = utils.GetEnv("GBT_CAR_GIT_STATUS_COPIED_COUNT_TEXT", strconv.Itoa(status.copied))
                }

                if status.deleted > 0 {
                    defaultStatusDeletedFormat = utils.GetEnv("GBT_CAR_GIT_STATUS_DELETED_FORMAT", "{{ StatusDeletedSymbol }}")
                    defaultStatusDeletedSymbolText = utils.GetEnv("GBT_CAR_GIT_STATUS_DELETED_SYMBOL_TEXT", " \u2796")
                    defaultStatusDeletedCountText = utils.GetEnv("GBT_CAR_GIT_STATUS_DELETED_COUNT_TEXT", strconv.Itoa(status.deleted))
                }

                if status.ignored > 0 {
                    defaultStatusIgnoredFormat = utils.GetEnv("GBT_CAR_GIT_STATUS_IGNORED_FORMAT", "{{ StatusIgnoredSymbol }}")
                    defaultStatusIgnoredSymbolText = utils.GetEnv("GBT_CAR_GIT_STATUS_IGNORED_SYMBOL_TEXT", " \u25cb")
                    defaultStatusIgnoredCountText = utils.GetEnv("GBT_CAR_GIT_STATUS_IGNORED_COUNT_TEXT", strconv.Itoa(status.ignored))
                }

                if status.modified > 0 {
                    defaultStatusModifiedFormat = utils.GetEnv("GBT_CAR_GIT_STATUS_MODIFIED_FORMAT", "{{ StatusModifiedSymbol }}")
                    defaultStatusModifiedSymbolText = utils.GetEnv("GBT_CAR_GIT_STATUS_MODIFIED_SYMBOL_TEXT", " \u271a")
                    defaultStatusModifiedCountText = utils.GetEnv("GBT_CAR_GIT_STATUS_MODIFIED_COUNT_TEXT", strconv.Itoa(status.modified))
                }

                if status.renamed > 0 {
                    defaultStatusRenamedFormat = utils.GetEnv("GBT_CAR_GIT_STATUS_RENAMED_FORMAT", "{{ StatusRenamedSymbol }}")
                    defaultStatusRenamedSymbolText = utils.GetEnv("GBT_CAR_GIT_STATUS_RENAMED_SYMBOL_TEXT", " \u2972")
                    defaultStatusRenamedCountText = utils.GetEnv("GBT_CAR_GIT_STATUS_RENAMED_COUNT_TEXT", strconv.Itoa(status.renamed))
                }

                if status.staged > 0 {
                    defaultStatusStagedFormat = utils.GetEnv("GBT_CAR_GIT_STATUS_STAGED_FORMAT", "{{ StatusStagedSymbol }}")
                    defaultStatusStagedSymbolText = utils.GetEnv("GBT_CAR_GIT_STATUS_STAGED_SYMBOL_TEXT", " \u25cf")
                    defaultStatusStagedCountText = utils.GetEnv("GBT_CAR_GIT_STATUS_STAGED_COUNT_TEXT", strconv.Itoa(status.staged))
                }

                if status.unmerged > 0 {
                    defaultStatusUnmergedFormat = utils.GetEnv("GBT_CAR_GIT_STATUS_UNMERGED_FORMAT", "{{ StatusUnmergedSymbol }}")
                    defaultStatusUnmergedSymbolText = utils.GetEnv("GBT_CAR_GIT_STATUS_UNMERGED_SYMBOL_TEXT", " \u2716")
                    defaultStatusUnmergedCountText = utils.GetEnv("GBT_CAR_GIT_STATUS_UNMERGED_COUNT_TEXT", strconv.Itoa(status.unmerged))
                }

                if status.untracked > 0 {
                    defaultStatusUntrackedFormat = utils.GetEnv("GBT_CAR_GIT_STATUS_UNTRACKED_FORMAT", "{{ StatusUntrackedSymbol }}")
                    defaultStatusUntrackedSymbolText = utils.GetEnv("GBT_CAR_GIT_STATUS_UNTRACKED_SYMBOL_TEXT", " \u2026")
                    defaultStatusUntrackedCountText = utils.GetEnv("GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_TEXT", strconv.Itoa(status.untracked))
                }
            }
        }

        if ok := reTemplatingAhead.MatchString(defaultRootFormat); ok {
            ahead := true
            aheadCount := compareRemote(ahead)

            if aheadCount != "0" {
                defaultAheadFormat = utils.GetEnv("GBT_CAR_GIT_AHEAD_FORMAT", "{{ AheadSymbol }}")
                defaultAheadSymbolText = utils.GetEnv("GBT_CAR_GIT_AHEAD_SYMBOL_TEXT", " \u2b06")
                defaultAheadCountText = utils.GetEnv("GBT_CAR_GIT_AHEAD_COUNT_TEXT", aheadCount)
            }
        }

        if ok := reTemplatingBehind.MatchString(defaultRootFormat); ok {
            behind := false
            behindCount := compareRemote(behind)

            if behindCount != "0" {
                defaultBehindFormat = utils.GetEnv("GBT_CAR_GIT_BEHIND_FORMAT", "{{ BehindSymbol }}")
                defaultBehindSymbolText = utils.GetEnv("GBT_CAR_GIT_BEHIND_SYMBOL_TEXT", " \u2b07")
                defaultBehindCountText = utils.GetEnv("GBT_CAR_GIT_BEHIND_COUNT_TEXT", behindCount)
            }
        }

        if ok := reTemplatingStash.MatchString(defaultRootFormat); ok {
            stashCount := getStash()

            if stashCount != "0" {
                defaultStashFormat = utils.GetEnv("GBT_CAR_GIT_STASH_FORMAT", "{{ StashSymbol }}")
                defaultStashSymbolText = utils.GetEnv("GBT_CAR_GIT_STASH_SYMBOL_TEXT", " \u2691")
                defaultStashCountText = utils.GetEnv("GBT_CAR_GIT_STASH_COUNT_TEXT", stashCount)
            }
        }
    }

    c.Model = map[string]car.ModelElement{
        "root": {
            Bg: utils.GetEnv("GBT_CAR_GIT_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_GIT_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_GIT_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_GIT_FORMAT", defaultRootFormat),
        },
        "Icon": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_ICON_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_ICON_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_ICON_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_GIT_ICON_TEXT", "\ue0a0"),
        },
        "Head": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_HEAD_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_HEAD_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_HEAD_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_GIT_HEAD_TEXT", defaultHeadText),
        },

        "Status": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_FORMAT", defaultStatusFormat),
        },
        "StatusDirty": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DIRTY_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DIRTY_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", utils.GetEnv(
                            "GBT_CAR_FG", "red")))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DIRTY_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: utils.GetEnv("GBT_CAR_GIT_STATUS_DIRTY_TEXT", "\u2718"),
        },
        "StatusClean": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_CLEAN_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_CLEAN_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", utils.GetEnv(
                            "GBT_CAR_FG", "green")))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_CLEAN_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: utils.GetEnv("GBT_CAR_GIT_STATUS_CLEAN_TEXT", "\u2714"),
        },

        "StatusAdded": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_ADDED_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_ADDED_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_ADDED_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultStatusAddedFormat,
        },
        "StatusAddedSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_ADDED_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_ADDED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_ADDED_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_ADDED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_ADDED_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_ADDED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusAddedSymbolText,
        },
        "StatusAddedCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_ADDED_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_ADDED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_ADDED_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_ADDED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_ADDED_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_ADDED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusAddedCountText,
        },

        "StatusCopied": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_COPIED_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_COPIED_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_COPIED_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultStatusCopiedFormat,
        },
        "StatusCopiedSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_COPIED_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_COPIED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_COPIED_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_COPIED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_COPIED_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_COPIED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusCopiedSymbolText,
        },
        "StatusCopiedCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_COPIED_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_COPIED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_COPIED_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_COPIED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_COPIED_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_COPIED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusCopiedCountText,
        },

        "StatusDeleted": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DELETED_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DELETED_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DELETED_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultStatusDeletedFormat,
        },
        "StatusDeletedSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DELETED_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_DELETED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DELETED_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_DELETED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DELETED_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_DELETED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusDeletedSymbolText,
        },
        "StatusDeletedCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DELETED_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_DELETED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DELETED_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_DELETED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_DELETED_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_DELETED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusDeletedCountText,
        },

        "StatusIgnored": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_IGNORED_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_IGNORED_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_IGNORED_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultStatusIgnoredFormat,
        },
        "StatusIgnoredSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_IGNORED_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_IGNORED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_IGNORED_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_IGNORED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_IGNORED_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_IGNORED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusIgnoredSymbolText,
        },
        "StatusIgnoredCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_IGNORED_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_IGNORED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_IGNORED_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_IGNORED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_IGNORED_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_IGNORED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusIgnoredCountText,
        },


        "StatusModified": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_MODIFIED_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_MODIFIED_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_MODIFIED_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultStatusModifiedFormat,
        },
        "StatusModifiedSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_MODIFIED_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_MODIFIED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_MODIFIED_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_MODIFIED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", utils.GetEnv(
                            "GBT_CAR_FG", "blue")))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_MODIFIED_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_MODIFIED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusModifiedSymbolText,
        },
        "StatusModifiedCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_MODIFIED_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_MODIFIED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_MODIFIED_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_MODIFIED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_MODIFIED_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_MODIFIED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusModifiedCountText,
        },

        "StatusRenamed": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_RENAMED_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_RENAMED_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_RENAMED_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultStatusRenamedFormat,
        },
        "StatusRenamedSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_RENAMED_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_RENAMED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_RENAMED_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_RENAMED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_RENAMED_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_RENAMED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusRenamedSymbolText,
        },
        "StatusRenamedCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_RENAMED_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_RENAMED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_RENAMED_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_RENAMED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_RENAMED_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_RENAMED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusRenamedCountText,
        },

        "StatusStaged": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_STAGED_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_STAGED_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_STAGED_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultStatusStagedFormat,
        },
        "StatusStagedSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_STAGED_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_STAGED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_STAGED_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_STAGED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", utils.GetEnv(
                            "GBT_CAR_FG", "red")))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_STAGED_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_STAGED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusStagedSymbolText,
        },
        "StatusStagedCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_STAGED_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_STAGED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_STAGED_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_STAGED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_STAGED_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_STAGED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusStagedCountText,
        },

        "StatusUnmerged": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNMERGED_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNMERGED_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNMERGED_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultStatusUnmergedFormat,
        },
        "StatusUnmergedSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNMERGED_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNMERGED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNMERGED_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNMERGED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", utils.GetEnv(
                            "GBT_CAR_FG", "red")))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNMERGED_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNMERGED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusUnmergedSymbolText,
        },
        "StatusUnmergedCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNMERGED_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNMERGED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNMERGED_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNMERGED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNMERGED_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNMERGED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusUnmergedCountText,
        },

        "StatusUntracked": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNTRACKED_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNTRACKED_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNTRACKED_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultStatusUntrackedFormat,
        },
        "StatusUntrackedSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNTRACKED_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNTRACKED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNTRACKED_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNTRACKED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNTRACKED_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNTRACKED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusUntrackedSymbolText,
        },
        "StatusUntrackedCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNTRACKED_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNTRACKED_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STATUS_UNTRACKED_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStatusUntrackedCountText,
        },

        "Ahead": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultAheadFormat,
        },
        "AheadSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_AHEAD_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_AHEAD_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_AHEAD_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultAheadSymbolText,
        },
        "AheadCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_AHEAD_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_AHEAD_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_AHEAD_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_AHEAD_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultAheadCountText,
        },

        "Behind": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultBehindFormat,
        },
        "BehindSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BEHIND_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_BEHIND_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_BEHIND_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultBehindSymbolText,
        },
        "BehindCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BEHIND_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_BEHIND_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_BEHIND_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_BEHIND_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultBehindCountText,
        },

        "Stash": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STASH_BG", utils.GetEnv(
                    "GBT_CAR_GIT_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STASH_FG", utils.GetEnv(
                    "GBT_CAR_GIT_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STASH_FM", utils.GetEnv(
                    "GBT_CAR_GIT_FM", defaultRootFm)),
            Text: defaultStashFormat,
        },
        "StashSymbol": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STASH_SYMBOL_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STASH_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STASH_SYMBOL_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STASH_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STASH_SYMBOL_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STASH_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStashSymbolText,
        },
        "StashCount": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_STASH_COUNT_BG", utils.GetEnv(
                    "GBT_CAR_GIT_STASH_BG", utils.GetEnv(
                        "GBT_CAR_GIT_BG", defaultRootBg))),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_STASH_COUNT_FG", utils.GetEnv(
                    "GBT_CAR_GIT_STASH_FG", utils.GetEnv(
                        "GBT_CAR_GIT_FG", defaultRootFg))),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_STASH_COUNT_FM", utils.GetEnv(
                    "GBT_CAR_GIT_STASH_FM", utils.GetEnv(
                        "GBT_CAR_GIT_FM", defaultRootFm))),
            Text: defaultStashCountText,
        },

        "Sep": {
            Bg: utils.GetEnv(
                "GBT_CAR_GIT_SEP_BG", utils.GetEnv(
                    "GBT_SEPARATOR_BG", defaultSep)),
            Fg: utils.GetEnv(
                "GBT_CAR_GIT_SEP_FG", utils.GetEnv(
                    "GBT_SEPARATOR_FG", defaultSep)),
            Fm: utils.GetEnv(
                "GBT_CAR_GIT_SEP_FM", utils.GetEnv(
                    "GBT_SEPARATOR_FM", defaultSep)),
            Text: utils.GetEnv(
                "GBT_CAR_GIT_SEP", utils.GetEnv(
                    "GBT_CAR_GIT_SEP_TEXT", utils.GetEnv(
                        "GBT_SEPARATOR", defaultSep))),
        },
    }

    c.Wrap = utils.GetEnvBool("GBT_CAR_GIT_WRAP", false)
}
