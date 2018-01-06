package git

import (
    "testing"
)

func TestInit(t *testing.T) {
    tests := []struct {
        runIsGitDir []string
        runGetBranch []string
        runGetTag []string
        runGetCommit []string
        runIsDirty []string
        runCompareRemote []string
        field string
        expectedDisplay bool
        expectedOutput string
    }{
        {
            runIsGitDir: []string{"echo"},
            expectedDisplay: true,
        },
        {
            runIsGitDir: []string{"nothing"},
            expectedDisplay: false,
        },
        {
            runIsGitDir: []string{"echo"},
            runIsDirty: []string{"echo", "-n"},
            field: "Status",
            expectedOutput: "{{ Clean }}",
            expectedDisplay: true,
        },
        {
            runIsGitDir: []string{"echo"},
            runIsDirty: []string{"echo", "YES"},
            field: "Status",
            expectedOutput: "{{ Dirty }}",
            expectedDisplay: true,
        },
        {
            runIsGitDir: []string{"echo"},
            runGetBranch: []string{"echo", "refs/heads/master"},
            field: "Head",
            expectedOutput: "master",
            expectedDisplay: true,
        },
        {
            runIsGitDir: []string{"echo"},
            runGetBranch: []string{"nothing"},
            runGetTag: []string{"echo", "v1.2.3"},
            field: "Head",
            expectedOutput: "v1.2.3",
            expectedDisplay: true,
        },
        {
            runIsGitDir: []string{"echo"},
            runGetBranch: []string{"nothing"},
            runGetTag: []string{"nothing"},
            runGetCommit: []string{"echo", "92387be"},
            field: "Head",
            expectedOutput: "92387be",
            expectedDisplay: true,
        },
        {
            runIsGitDir: []string{"echo"},
            runCompareRemote: []string{"nothing"},
            field: "Ahead",
            expectedOutput: "",
            expectedDisplay: true,
        },
        {
            runIsGitDir: []string{"echo"},
            runCompareRemote: []string{"nothing"},
            field: "Behind",
            expectedOutput: "",
            expectedDisplay: true,
        },
        {
            runIsGitDir: []string{"echo"},
            runCompareRemote: []string{"echo", "1"},
            field: "Ahead",
            expectedOutput: " ⬆",
            expectedDisplay: true,
        },
        {
            runIsGitDir: []string{"echo"},
            runCompareRemote: []string{"echo", "1"},
            field: "Behind",
            expectedOutput: " ⬇",
            expectedDisplay: true,
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
