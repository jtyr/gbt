package docker

import (
    "testing"
    "runtime"
    "path/filepath"
    "os"
)

func TestInit(t *testing.T) {
    _, filename, _, _ := runtime.Caller(0)
    path := filepath.Dir(filename)

    testDataExists := filepath.Join(path, "testdata", "exists")
    testDataNotExists := filepath.Join(path, "testdata", "not-exists")

    tests := []struct {
        runDockerVersion []string
        composeFiles     string
        workdir          string
        expectedDisplay  bool
        expectedVersion  string
    }{
        {
            runDockerVersion: []string{"echo", "18.06.0-ce"},
            composeFiles:     "",
            workdir:          filepath.Join(testDataExists, "subdir"),
            expectedDisplay:  true,
            expectedVersion:  "18.06.0-ce",
        },
        {
            runDockerVersion: []string{"echo", "18.06.0-ce"},
            composeFiles:     "",
            workdir:          filepath.Join(testDataExists, "subdir", "subdir"),
            expectedDisplay:  true,
            expectedVersion:  "18.06.0-ce",
        },
        {
            runDockerVersion: []string{"echo", "18.06.0-ce"},
            composeFiles:     "",
            workdir:          filepath.Join(testDataExists, "subdir", "subdir", "subdir"),
            expectedDisplay:  true,
            expectedVersion:  "18.06.0-ce",
        },
        {
            runDockerVersion: []string{"echo", "18.06.0-ce"},
            composeFiles:     "",
            workdir:          filepath.Join(testDataExists, "subdir2"),
            expectedDisplay:  true,
            expectedVersion:  "18.06.0-ce",
        },
        {
            runDockerVersion: []string{"echo", "18.06.0-ce"},
            composeFiles:     "",
            workdir:          filepath.Join(testDataNotExists, "subdir"),
            expectedDisplay:  false,
            expectedVersion:  "",
        },
        {
            runDockerVersion: []string{"echo", "18.06.0-ce"},
            composeFiles:     "",
            workdir:          filepath.Join(testDataNotExists, "subdir", "subdir"),
            expectedDisplay:  false,
            expectedVersion:  "",
        },
        {
            runDockerVersion: []string{"echo", "18.06.0-ce"},
            composeFiles:     "",
            workdir:          filepath.Join(testDataNotExists, "subdir", "subdir", "subdir"),
            expectedDisplay:  false,
            expectedVersion:  "",
        },
        {
            runDockerVersion: []string{"echo", "18.06.0-ce"},
            composeFiles:     "",
            workdir:          filepath.Join(testDataNotExists, "subdir2"),
            expectedDisplay:  false,
            expectedVersion:  "",
        },
        {
            runDockerVersion: []string{"commandnotexists"},
            composeFiles:     "",
            workdir:          "",
            expectedDisplay:  false,
            expectedVersion:  "",
        },
    }

    oldCwd, _ := os.Getwd()
    for i, test := range tests {
        if test.runDockerVersion != nil {
            runDockerVersion = test.runDockerVersion
        }

        car := Car{}

        os.Chdir(test.workdir)
        car.Init()
        os.Chdir(oldCwd)

        if car.Display != test.expectedDisplay {
            t.Errorf("Test [%d]: Expected car.Display to be %t, got %t.", i, test.expectedDisplay, car.Display)
        }

        if test.expectedVersion != "" && car.Model["Version"].Text != test.expectedVersion {
            t.Errorf("Test [%d]: Expected car.Model.Version.Text to be '%s', got '%s'.", i, test.expectedVersion, car.Model["Version"].Text)
        }
    }
}
