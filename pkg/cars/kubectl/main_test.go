package kubectl

import (
    "testing"
)

func TestInit(t *testing.T) {
    tests := []struct {
        runKubectlCurrentContext []string
        runGetContexts           []string
        expectedDisplay          bool
        expectedContext          string
        expectedCluster          string
        expectedAuthInfo         string
        expectedNamespace        string
    }{
        {
            runKubectlCurrentContext: []string{"echo", "minikube"},
            runGetContexts:           []string{"echo", "CURRENT   NAME            CLUSTER         AUTHINFO        NAMESPACE\n*         kubename        kubecluster     kubeauth\n"},
            expectedDisplay:          true,
            expectedContext:          "kubename",
            expectedCluster:          "kubecluster",
            expectedAuthInfo:         "kubeauth",
            expectedNamespace:        "",
        },
        {
            runKubectlCurrentContext: []string{"echo", "minikube"},
            runGetContexts:           []string{"echo", "CURRENT   NAME            CLUSTER         AUTHINFO        NAMESPACE\n*         context        cluster        authinfo        namespace\n"},
            expectedDisplay:          true,
            expectedContext:          "context",
            expectedCluster:          "cluster",
            expectedAuthInfo:         "authinfo",
            expectedNamespace:        "namespace",
        },
        {
            runKubectlCurrentContext: []string{"commandnotexists"},
            runGetContexts:           []string{"nothing"},
            expectedDisplay:          false,
            expectedContext:          "",
            expectedCluster:          "",
            expectedAuthInfo:         "",
            expectedNamespace:        "",
        },
        {
            runKubectlCurrentContext: []string{"echo"}, // no output
            runGetContexts:           []string{"nothing"},
            expectedDisplay:          false,
            expectedContext:          "",
            expectedCluster:          "",
            expectedAuthInfo:         "",
            expectedNamespace:        "",
        },
    }

    for i, test := range tests {
        if test.runGetContexts != nil {
            runGetContexts = test.runGetContexts
        }

        if len(test.runKubectlCurrentContext) > 0 {
            runKubectlCurrentContext = test.runKubectlCurrentContext
        }

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

        if test.expectedNamespace != "" && car.Model["Namespace"].Text != test.expectedNamespace {
            t.Errorf("Test [%d]: Expected car.Model.Namespace.Text to be '%s', got '%s'.", i, test.expectedNamespace, car.Model["Namespace"].Text)
        }
    }
}
