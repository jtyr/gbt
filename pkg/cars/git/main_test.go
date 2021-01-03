package git

import (
    "testing"

    ct "github.com/jtyr/gbt/pkg/core/testing"
)

func TestInit(t *testing.T) {
    ct.ResetEnv()

    tests := []struct {
        runIsGitDir      []string
        runGetBranch     []string
        runGetTag        []string
        runGetCommit     []string
        runIsDirty       []string
        runCompareRemote []string
        runStash         []string
        field            string
        expectedDisplay  bool
        expectedOutput   string
    }{
        {
            runIsGitDir:     []string{"echo"},
            expectedDisplay: true,
        },
        {
            runIsGitDir:     []string{"nothing"},
            expectedDisplay: false,
        },
        {
            runIsGitDir:     []string{"echo"},
            runIsDirty:      []string{"echo", "-n"},
            field:           "Status",
            expectedOutput:  "{{ StatusClean }}",
            expectedDisplay: true,
        },
        {
            runIsGitDir:     []string{"echo"},
            runIsDirty:      []string{"echo", "-e", "AA added\n C copied\n D deleted\n ! ignored\n M modified\n R renamed\n   staged\n U unmerged\n?? untracked"},
            field:           "Status",
            expectedOutput:  "{{ StatusDirty }}",
            expectedDisplay: true,
        },
        {
            runIsGitDir:     []string{"echo"},
            runGetBranch:    []string{"echo", "refs/heads/master"},
            field:           "Head",
            expectedOutput:  "master",
            expectedDisplay: true,
        },
        {
            runIsGitDir:     []string{"echo"},
            runGetBranch:    []string{"nothing"},
            runGetTag:       []string{"echo", "v1.2.3"},
            field:           "Head",
            expectedOutput:  "v1.2.3",
            expectedDisplay: true,
        },
        {
            runIsGitDir:     []string{"echo"},
            runGetBranch:    []string{"nothing"},
            runGetTag:       []string{"nothing"},
            runGetCommit:    []string{"echo", "92387be"},
            field:           "Head",
            expectedOutput:  "92387be",
            expectedDisplay: true,
        },
        {
            runIsGitDir:      []string{"echo"},
            runCompareRemote: []string{"nothing"},
            field:            "Ahead",
            expectedOutput:   "",
            expectedDisplay:  true,
        },
        {
            runIsGitDir:      []string{"echo"},
            runCompareRemote: []string{"echo", "4"},
            field:            "AheadCount",
            expectedOutput:   "4",
            expectedDisplay:  true,
        },
        {
            runIsGitDir:      []string{"echo"},
            runCompareRemote: []string{"echo", "1"},
            field:            "AheadSymbol",
            expectedOutput:   " ⬆",
            expectedDisplay:  true,
        },
        {
            runIsGitDir:      []string{"echo"},
            runCompareRemote: []string{"nothing"},
            field:            "Behind",
            expectedOutput:   "",
            expectedDisplay:  true,
        },
        {
            runIsGitDir:      []string{"echo"},
            runCompareRemote: []string{"echo", "3"},
            field:            "BehindCount",
            expectedOutput:   "3",
            expectedDisplay:  true,
        },
        {
            runIsGitDir:      []string{"echo"},
            runCompareRemote: []string{"echo", "1"},
            field:            "BehindSymbol",
            expectedOutput:   " ⬇",
            expectedDisplay:  true,
        },
        {
            runIsGitDir:      []string{"echo"},
            runStash:         []string{"nothing"},
            field:            "Stash",
            expectedOutput:   "",
            expectedDisplay:  true,
        },
        {
            runIsGitDir:      []string{"echo"},
            runStash:         []string{"echo", "-e", "a\nb\nc"},
            field:            "StashCount",
            expectedOutput:   "3",
            expectedDisplay:  true,
        },
        {
            runIsGitDir:      []string{"echo"},
            runStash:         []string{"echo", "1"},
            field:            "StashSymbol",
            expectedOutput:   " ⚑",
            expectedDisplay:  true,
        },
    }

    for i, test := range tests {
        if test.runIsGitDir != nil {
            runIsGitDir = test.runIsGitDir
        } else {
            runIsGitDir = []string{"git", "rev-parse", "--git-dir"}
        }

        if test.runGetBranch != nil {
            runGetBranch = test.runGetBranch
        } else {
            runGetBranch = []string{"git", "symbolic-ref", "HEAD"}
        }

        if test.runGetTag != nil {
            runGetTag = test.runGetTag
        } else {
            runGetTag = []string{"git", "describe", "--tags", "--exact-match", "HEAD"}
        }

        if test.runGetCommit != nil {
            runGetCommit = test.runGetCommit
        } else {
            runGetCommit = []string{"git", "rev-parse", "--short", "HEAD"}
        }

        if test.runIsDirty != nil {
            runIsDirty = test.runIsDirty
        } else {
            runIsDirty = []string{"git", "status", "--porcelain"}
        }

        if test.runCompareRemote != nil {
            runCompareRemote = test.runCompareRemote
        } else {
            runCompareRemote = []string{"git", "rev-list", "--count"}
        }

        if test.runStash != nil {
            runStash = test.runStash
        } else {
            runStash = []string{"git", "stash", "list"}
        }

        defaultRootFormat = " {{ Icon }} {{ Head }} {{ Status }}{{ Ahead }}{{ Behind }}{{ Stash }} "

        car := Car{}

        car.Init()

        if car.Display != test.expectedDisplay {
            t.Errorf("Test [%d]: Expected car.Display to be %t, got %t.", i, test.expectedDisplay, car.Display)
        }

        if test.field != "" && car.Model[test.field].Text != test.expectedOutput {
            t.Errorf("Test [%d]: Expected car.Model.%s.Text to be '%s', got '%s'.", i, test.field, test.expectedOutput, car.Model[test.field].Text)
        }
    }
}
