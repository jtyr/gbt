package azure

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
        config                string
        cloudsConfig          string
        azureProfile          string
        envCloud              string
        envDefaultsGroup      string
        expectedCloud         string
        expectedSubscription  string
        expectedUserName      string
        expectedUserType      string
        expectedState         string
        expectedDefaultsGroup string
    }{
        {
            expectedCloud:         "AzureCloud",
        },
        {
            config:                "[cloud]\nname = AzureGermanCloud",
            expectedCloud:         "AzureGermanCloud",
        },
        {
            config:                "[cloud]\nname = AzureCloud",
            cloudsConfig:          "[AzureCloud]\nsubscription = 12345678-1234-1234-1234-123456789012",
            azureProfile:          "\xEF\xBB\xBF" + `{
                "subscriptions": [
                  {
                    "id": "12345678-1234-1234-1234-123456789012",
                    "name": "test-subscription",
                    "user": {
                      "name": "test_name",
                      "type": "user"
                    },
                    "state": "Enabled"
                  }
                ]}`,
            expectedCloud:         "AzureCloud",
            expectedSubscription:  "test-subscription",
            expectedUserName:      "test_name",
            expectedUserType:      "user",
            expectedState:         "Enabled",
        },
        {
            config:                "[cloud]\nname = AzureCloud\n\n[defaults]\ngroup = test-rg",
            expectedCloud:         "AzureCloud",
            expectedDefaultsGroup: "test-rg",
        },
    }

    for i, test := range tests {
        tmp := os.TempDir()

        // Create faked home and config directory
        tmpDir, err := ioutil.TempDir(tmp, "testAzure")
        if err != nil {
            t.Errorf("failed to create temp home dir: %s", err)
        }
        defer os.RemoveAll(tmpDir)

        err = os.MkdirAll(filepath.Join(tmpDir, ".azure"), 0755)
        if err != nil {
            t.Errorf("failed to create .azure dir: %s", err)
        }

        usr = &user.User{
            HomeDir: tmpDir,
        }

        // Write config
        if test.config != "" {
            config := filepath.Join(tmpDir, ".azure", "config")
            if err := ioutil.WriteFile(config, []byte(test.config), 0644); err != nil {
                t.Errorf("failed to write config file: %s", err)
            }
        }

        // Write clouds.config
        if test.cloudsConfig != "" {
            cloudsConfig := filepath.Join(tmpDir, ".azure", "clouds.config")
            if err := ioutil.WriteFile(cloudsConfig, []byte(test.cloudsConfig), 0644); err != nil {
                t.Errorf("failed to write clouds.config file: %s", err)
            }
        }

        // Write azureProfile.json
        if test.azureProfile != "" {
            azureProfile := filepath.Join(tmpDir, ".azure", "azureProfile.json")
            if err := ioutil.WriteFile(azureProfile, []byte(test.azureProfile), 0644); err != nil {
                t.Errorf("failed to write azureProfile.json file: %s", err)
            }
        }

        if test.envCloud != "" {
            os.Setenv("AZURE_CLOUD_NAME", test.envCloud)
        } else {
            os.Unsetenv("AZURE_CLOUD_NAME")
        }

        if test.envDefaultsGroup != "" {
            os.Setenv("AZURE_DEFAULTS_GROUP", test.envDefaultsGroup)
        } else {
            os.Unsetenv("AZURE_DEFAULTS_GROUP")
        }

        car := Car{}
        car.Init()

        if car.Model["Cloud"].Text != test.expectedCloud {
            t.Errorf("Test [%d]: Expected Cloud '%s', found '%s'.", i, test.expectedCloud, car.Model["Cloud"].Text)
        }

        if car.Model["Subscription"].Text != test.expectedSubscription {
            t.Errorf("Test [%d]: Expected Subscription '%s', found '%s'.", i, test.expectedSubscription, car.Model["Subscription"].Text)
        }

        if car.Model["UserName"].Text != test.expectedUserName {
            t.Errorf("Test [%d]: Expected UserName '%s', found '%s'.", i, test.expectedUserName, car.Model["UserName"].Text)
        }

        if car.Model["UserType"].Text != test.expectedUserType {
            t.Errorf("Test [%d]: Expected UserType '%s', found '%s'.", i, test.expectedUserType, car.Model["UserType"].Text)
        }

        if car.Model["State"].Text != test.expectedState {
            t.Errorf("Test [%d]: Expected State '%s', found '%s'.", i, test.expectedState, car.Model["State"].Text)
        }

        if car.Model["DefaultsGroup"].Text != test.expectedDefaultsGroup {
            t.Errorf("Test [%d]: Expected DefaultsGroup '%s', found '%s'.", i, test.expectedDefaultsGroup, car.Model["DefaultsGroup"].Text)
        }
    }
}
