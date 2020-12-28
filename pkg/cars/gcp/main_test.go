package dir

import (
    "fmt"
    "io/ioutil"
    "os"
    "os/user"
    "path/filepath"
    "testing"
)

func TestInit(t *testing.T) {
    tests := []struct {
        configName      string
        configContent   string
        credFileContent string
        projectAliases  string
        os              string
        expectedConfig  string
        expectedAccount string
        expectedProject string
    }{
        {
            configName:      "test",
            configContent:   "[core]\naccount = test@example.com\nproject = my-test-project-123456\n",
            os:              "linux",
            expectedConfig:  "test",
            expectedAccount: "test@example.com",
            expectedProject: "my-test-project-123456",
        },
        {
            configName:      "test",
            configContent:   "[core]\naccount = test@example.com\nproject = my-test-project-123456\n",
            os:              "windows",
            expectedConfig:  "test",
            expectedAccount: "test@example.com",
            expectedProject: "my-test-project-123456",
        },
        {
            configName:      "test",
            configContent:   "[core]\naccount = test@example.com\nproject = my-test-project-123456\n",
            projectAliases:  "my-test-project-123456=my-test",
            os:              "linux",
            expectedConfig:  "test",
            expectedAccount: "test@example.com",
            expectedProject: "my-test",
        },
        {
            configName:    "test",
            configContent: "[core]\naccount = test@example.com\nproject = my-test-project-123456\n",
            credFileContent: `{
                "type": "service_account",
                "project_id": "my-test-project-6543231",
                "private_key_id": "",
                "private_key": "",
                "client_email": "test2@example.com",
                "client_id": "",
                "auth_uri": "",
                "token_uri": "",
                "auth_provider_x509_cert_url": "",
                "client_x509_cert_url": ""}`,
            os:              "windows",
            expectedConfig:  "test",
            expectedAccount: "test2@example.com",
            expectedProject: "my-test-project-123456",
        },
    }

    for i, test := range tests {
        tmp := os.TempDir()

        tmpHomeDir := ""
        credFile := ""

        // Write config
        if test.configContent != "" {
            tmpDir, err := ioutil.TempDir(tmp, "testGcp")
            tmpHomeDir = tmpDir
            if err != nil {
                t.Errorf("failed to create temp home dir: %s", err)
            }
            defer os.RemoveAll(tmpDir)

            err = os.MkdirAll(filepath.Join(tmpDir, ".config", "gcloud", "configurations"), 0755)
            if err != nil {
                t.Errorf("failed to create temp dir: %s", err)
            }

            activeConfig := filepath.Join(tmpDir, ".config", "gcloud", "active_config")
            if err := ioutil.WriteFile(activeConfig, []byte(test.configName), 0644); err != nil {
                t.Errorf("failed to write active_config file: %s", err)
            }

            config := filepath.Join(tmpDir, ".config", "gcloud", "configurations", fmt.Sprintf("config_%s", test.configName))
            if err := ioutil.WriteFile(config, []byte(test.configContent), 0644); err != nil {
                t.Errorf("failed to write config file: %s", err)
            }
        }

        // Write cred file
        if test.credFileContent != "" {
            tmpFile, err := ioutil.TempFile(tmp, "testGcpCredFile")
            credFile = tmpFile.Name()
            if err != nil {
                t.Errorf("failed to crete cred file: %s", err)
            }
            defer os.Remove(tmpFile.Name())

            if _, err := tmpFile.Write([]byte(test.credFileContent)); err != nil {
                t.Errorf("failed to write cred file: %s", err)
            }

            if err := tmpFile.Close(); err != nil {
                t.Errorf("failed to close cred file: %s", err)
            }
        }

        car := Car{}
        runtimeOs = test.os

        if test.credFileContent != "" {
            os.Setenv("CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE", credFile)
        } else {
            os.Unsetenv("CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE")
        }

        if test.os == "linux" {
            usr = &user.User{
                HomeDir: tmpHomeDir,
            }
        } else {
            os.Setenv("APPDATA", fmt.Sprintf("%s/.config", tmpHomeDir))
        }

        if test.projectAliases != "" {
            os.Setenv("GBT_CAR_GCP_PROJECT_ALIASES", test.projectAliases)
        } else {
            os.Unsetenv("GBT_CAR_GCP_PROJECT_ALIASES")
        }

        car.Init()

        if car.Model["Config"].Text != test.expectedConfig {
            t.Errorf("Test [%d]: Expected Config '%s', found '%s'.", i, test.expectedConfig, car.Model["Config"].Text)
        }

        if car.Model["Account"].Text != test.expectedAccount {
            t.Errorf("Test [%d]: Expected Account '%s', found '%s'.", i, test.expectedAccount, car.Model["Account"].Text)
        }

        if car.Model["Project"].Text != test.expectedProject {
            t.Errorf("Test [%d]: Expected Project '%s', found '%s'.", i, test.expectedProject, car.Model["Project"].Text)
        }
    }
}
