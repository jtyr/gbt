package kubectl

import (
    "os"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"

    "k8s.io/client-go/tools/clientcmd"
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

// Return the current context information for kubectl.
func getCurrentContext() *kubeContextInfo {
    info := &kubeContextInfo{}

    loadingRules := clientcmd.ClientConfigLoadingRules{
        Precedence:       []string{os.Getenv(clientcmd.RecommendedConfigPathEnvVar), clientcmd.RecommendedHomeFile},
        WarnIfAllMissing: false,
    }

    mergedConfig, err := loadingRules.Load()
    if err != nil {
        return info
    }

    info.context = mergedConfig.CurrentContext

    for k, c := range mergedConfig.Contexts {
        if k == mergedConfig.CurrentContext {
            info.cluster   = c.Cluster
            info.authInfo  = c.AuthInfo

            if len(c.Namespace) == 0 {
                info.namespace = "default"
            } else {
                info.namespace = c.Namespace
            }
        }
    }

    return info
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "26")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "white")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
    defaultSep := "\000"

    contextInfo := getCurrentContext()

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

    c.Display = utils.GetEnvBool("GBT_CAR_KUBECTL_DISPLAY", len(contextInfo.context) > 0)
    c.Wrap = utils.GetEnvBool("GBT_CAR_KUBECTL_WRAP", false)
}
