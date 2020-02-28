package kubectl

import (
    "strings"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

type kubeContextInfo struct {
    context   string
    cluster   string
    authInfo  string
    namespace string
}

var runKubectlCurrentContext = []string{"kubectl", "config", "current-context"}
var runGetContexts = []string{"kubectl", "config", "get-contexts"}

// Returns true if kubectl command exists.
func isKubectlCurrentContextSet() bool {
    rc, output, _ := utils.Run(runKubectlCurrentContext)

    return rc == 0 && strings.TrimSpace(output) != ""
}

// Return the current context information for kubectl.
func getCurrentContext(display bool) *kubeContextInfo {
    kubectlInfo := &kubeContextInfo{
        namespace: "",
        context:   "",
        cluster:   "",
        authInfo:  "",
    }

    if ! display {
        return kubectlInfo
    }

    rc, out, _ := utils.Run(runGetContexts)

    if rc == 0 {
        lines := strings.Split(out, "\n")

        for _, line := range lines {
            if strings.HasPrefix(line, "*") {
                fields := strings.Fields(line)

                kubectlInfo.context = fields[1]
                kubectlInfo.cluster = fields[2]
                kubectlInfo.authInfo = fields[3]

                if len(fields) == 5 {
                    kubectlInfo.namespace = fields[4]
                }

                break
            }
        }
    }

    return kubectlInfo
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "26")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "white")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultSep := "\000"

    c.Display = utils.GetEnvBool("GBT_CAR_KUBECTL_DISPLAY", isKubectlCurrentContextSet())
    contextInfo := getCurrentContext(c.Display)

    c.Model = map[string]car.ModelElement{
        "root": {
            Bg: utils.GetEnv("GBT_CAR_KUBECTL_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_KUBECTL_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_KUBECTL_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_KUBECTL_FORMAT", " {{ Icon }} {{ Context }} "),
        },
        "Icon": {
            Bg: utils.GetEnv(
                "GBT_CAR_KUBECTL_ICON_BG", utils.GetEnv(
                    "GBT_CAR_KUBECTL_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_KUBECTL_ICON_FG", utils.GetEnv(
                    "GBT_CAR_KUBECTL_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_KUBECTL_ICON_FM", utils.GetEnv(
                    "GBT_CAR_KUBECTL_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_KUBECTL_ICON_TEXT", "\u2388"),
        },
        "Context": {
            Bg: utils.GetEnv(
                "GBT_CAR_KUBECTL_CONTEXT_BG", utils.GetEnv(
                    "GBT_CAR_KUBECTL_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_KUBECTL_CONTEXT_FG", utils.GetEnv(
                    "GBT_CAR_KUBECTL_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_KUBECTL_CONTEXT_FM", utils.GetEnv(
                    "GBT_CAR_KUBECTL_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_KUBECTL_CONTEXT_TEXT", contextInfo.context),
        },
        "Cluster": {
            Bg: utils.GetEnv(
                "GBT_CAR_KUBECTL_CLUSTER_BG", utils.GetEnv(
                    "GBT_CAR_KUBECTL_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_KUBECTL_CLUSTER_FG", utils.GetEnv(
                    "GBT_CAR_KUBECTL_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_KUBECTL_CLUSTER_FM", utils.GetEnv(
                    "GBT_CAR_KUBECTL_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_KUBECTL_CLUSTER_TEXT", contextInfo.cluster),
        },
        "AuthInfo": {
            Bg: utils.GetEnv(
                "GBT_CAR_KUBECTL_AUTHINFO_BG", utils.GetEnv(
                    "GBT_CAR_KUBECTL_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_KUBECTL_AUTHINFO_FG", utils.GetEnv(
                    "GBT_CAR_KUBECTL_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_KUBECTL_AUTHINFO_FM", utils.GetEnv(
                    "GBT_CAR_KUBECTL_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_KUBECTL_AUTHINFO_TEXT", contextInfo.authInfo),
        },
        "Namespace": {
            Bg: utils.GetEnv(
                "GBT_CAR_KUBECTL_NAMESPACE_BG", utils.GetEnv(
                    "GBT_CAR_KUBECTL_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_KUBECTL_NAMESPACE_FG", utils.GetEnv(
                    "GBT_CAR_KUBECTL_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_KUBECTL_NAMESPACE_FM", utils.GetEnv(
                    "GBT_CAR_KUBECTL_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_KUBECTL_NAMESPACE_TEXT", contextInfo.namespace),
        },
        "Sep": {
            Bg: utils.GetEnv(
                "GBT_CAR_KUBECTL_SEP_BG", utils.GetEnv(
                    "GBT_SEPARATOR_BG", defaultSep)),
            Fg: utils.GetEnv(
                "GBT_CAR_KUBECTL_SEP_FG", utils.GetEnv(
                    "GBT_SEPARATOR_FG", defaultSep)),
            Fm: utils.GetEnv(
                "GBT_CAR_KUBECTL_SEP_FM", utils.GetEnv(
                    "GBT_SEPARATOR_FM", defaultSep)),
            Text: utils.GetEnv(
                "GBT_CAR_KUBECTL_SEP", utils.GetEnv(
                    "GBT_CAR_KUBECTL_SEP_TEXT", utils.GetEnv(
                        "GBT_SEPARATOR", defaultSep))),
        },
    }

    c.Wrap = utils.GetEnvBool("GBT_CAR_KUBECTL_WRAP", false)
}
