package kubectl

import (
    "io/ioutil"
    "os"
    "testing"

    ct "github.com/jtyr/gbt/pkg/core/testing"
)

func TestInit(t *testing.T) {
    ct.ResetEnv()

    tests := []struct {
        config                   string
        expectedDisplay          bool
        expectedContext          string
        expectedCluster          string
        expectedAuthInfo         string
        expectedNamespace        string
    }{
        {
            config:            "{apiVersion: v1, kind: Config, current-context: kubename, contexts: [{name: kubename, context: {cluster: kubecluster, user: kubeauth}}], clusters: [{name: kubecluster}], users: [{name: kubeauth}]}",
            expectedDisplay:   true,
            expectedContext:   "kubename",
            expectedCluster:   "kubecluster",
            expectedAuthInfo:  "kubeauth",
            expectedNamespace: "default",
        },
        {
            config:            "{apiVersion: v1, kind: Config, current-context: context, contexts: [{name: context, context: {cluster: cluster, user: authinfo, namespace: namespace}}], clusters: [{name: cluster}], users: [{name: authinfo}]}",
            expectedDisplay:   true,
            expectedContext:   "context",
            expectedCluster:   "cluster",
            expectedAuthInfo:  "authinfo",
            expectedNamespace: "namespace",
        },
        {
            config:            ": : :",
            expectedDisplay:   false,
            expectedContext:   "",
            expectedCluster:   "",
            expectedAuthInfo:  "",
            expectedNamespace: "",
        },
        {
            config:            "",
            expectedDisplay:   false,
            expectedContext:   "",
            expectedCluster:   "",
            expectedAuthInfo:  "",
            expectedNamespace: "",
        },
        {
            config:            "",
            expectedDisplay:   false,
            expectedContext:   "",
            expectedCluster:   "",
            expectedAuthInfo:  "",
            expectedNamespace: "",
        },
    }

    for i, test := range tests {
        config, err := ioutil.TempFile("", "")
        if err != nil {
            t.Errorf("failed to create config file: %s", err)
        }

        if err := ioutil.WriteFile(config.Name(), []byte(test.config), 0644); err != nil {
            t.Errorf("failed to write config file: %s", err)
        }

        os.Setenv("KUBECONFIG", config.Name())

        car := Car{}

        car.Init()

        if car.Display != test.expectedDisplay {
            t.Errorf("Test [%d]: Expected car.Display to be %t, got %t.", i, test.expectedDisplay, car.Display)
        }

        if test.expectedContext != "" && car.Model["Context"].Text != test.expectedContext {
            t.Errorf("Test [%d]: Expected car.Model.Context.Text to be '%s', got '%s'.", i, test.expectedContext, car.Model["Context"].Text)
        }

        if test.expectedCluster != "" && car.Model["Cluster"].Text != test.expectedCluster {
            t.Errorf("Test [%d]: Expected car.Model.Cluster.Text to be '%s', got '%s'.", i, test.expectedCluster, car.Model["Cluster"].Text)
        }

        if test.expectedAuthInfo != "" && car.Model["AuthInfo"].Text != test.expectedAuthInfo {
            t.Errorf("Test [%d]: Expected car.Model.AuthInfo.Text to be '%s', got '%s'.", i, test.expectedAuthInfo, car.Model["AuthInfo"].Text)
        }

        if car.Model["Namespace"].Text != test.expectedNamespace {
            t.Errorf("Test [%d]: Expected car.Model.Namespace.Text to be '%s', got '%s'.", i, test.expectedNamespace, car.Model["Namespace"].Text)
        }
    }
}
