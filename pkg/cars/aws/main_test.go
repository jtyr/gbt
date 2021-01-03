package aws

import (
    "io/ioutil"
    "os"
    "os/user"
    "path/filepath"
    "testing"

    ct "github.com/jtyr/gbt/pkg/core/testing"
)

func TestInit(t *testing.T) {
    ct.ResetEnv()

    tests := []struct {
        configContent   string
        profileName     string
        expectedProfile string
        expectedRegion  string
    }{
        {
            configContent:   "[default]\nregion = eu-west-2\n",
            profileName:     "",
            expectedProfile: "default",
            expectedRegion:  "eu-west-2",
        },
        {
            configContent:   "[default]\nregion = eu-west-2\n[profile test]\nregion = eu-west-2\n",
            profileName:     "test",
            expectedProfile: "test",
            expectedRegion:  "eu-west-2",
        },
        {
            configContent:   "[default]\nregion = eu-west-2\n[profile test]\noutput = json\n",
            profileName:     "test",
            expectedProfile: "test",
            expectedRegion:  "",
        },
    }

    for i, test := range tests {
        tmp := os.TempDir()

        tmpHomeDir := ""

        // Write config
        if test.configContent != "" {
            tmpDir, err := ioutil.TempDir(tmp, "testAws")
            tmpHomeDir = tmpDir
            if err != nil {
                t.Errorf("failed to create temp home dir: %s", err)
            }
            defer os.RemoveAll(tmpDir)

            err = os.MkdirAll(filepath.Join(tmpDir, ".aws"), 0755)
            if err != nil {
                t.Errorf("failed to create temp dir: %s", err)
            }

            config := filepath.Join(tmpDir, ".aws", "config")
            if err := ioutil.WriteFile(config, []byte(test.configContent), 0644); err != nil {
                t.Errorf("failed to write config file: %s", err)
            }
        }

        usr = &user.User{
            HomeDir: tmpHomeDir,
        }

        if test.profileName != "" {
            os.Setenv("AWS_PROFILE", test.profileName)
        } else {
            os.Unsetenv("AWS_PROFILE")
        }

        car := Car{}
        car.Init()

        if car.Model["Profile"].Text != test.expectedProfile {
            t.Errorf("Test [%d]: Expected Profile '%s', found '%s'.", i, test.expectedProfile, car.Model["Profile"].Text)
        }

        if car.Model["Region"].Text != test.expectedRegion {
            t.Errorf("Test [%d]: Expected Region '%s', found '%s'.", i, test.expectedRegion, car.Model["Region"].Text)
        }
    }
}
