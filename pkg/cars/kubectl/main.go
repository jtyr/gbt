package kubectl

import (
	"github.com/jtyr/gbt/pkg/core/car"
	"github.com/jtyr/gbt/pkg/core/utils"
	"regexp"
	"os/exec"
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

var kubectlLookupCmd = "kubectl"
var runGetContexts = []string{"kubectl", "config", "get-contexts"}

// Returns true if kubectl command exists
func isKubectlExists() bool {
	_, err := exec.LookPath(kubectlLookupCmd)

	return err == nil
}

// return the current context information for kubectl
func getCurrentContext(display bool) *kubeContextInfo {
	if !display {
		return &kubeContextInfo{
			namespace: "",
			context:   "",
			cluster:   "",
			authInfo:  "",
		}
	}

	rc, out, _ := utils.Run(runGetContexts)

	matchGroupMapping := make(map[string]string)
	if rc == 0 {
		var re = regexp.MustCompile(`(?m)^\*[\t\f\v ]+(?P<Context>[\w\-]+)[\t\f\v ]+(?P<Cluster>[\w\-]+)\s+(?P<AuthInfo>[\w\-]+)(?:[\t\f\v ]+(?P<Namespace>[\w\-]+))?$`)
		groupNames := re.SubexpNames()
		for _, match := range re.FindAllStringSubmatch(out, -1) {
			for groupIdx, value := range match {
				name := groupNames[groupIdx]
				if name == "" {
					name = "*"
				}

				matchGroupMapping[name] = value
			}
		}
	}

	return &kubeContextInfo{
		cluster:   matchGroupMapping["Cluster"],
		context:   matchGroupMapping["Context"],
		authInfo:  matchGroupMapping["AuthInfo"],
		namespace: matchGroupMapping["Namespace"],
	}
}

// Init initializes the car.
func (c *Car) Init() {
	defaultRootBg := utils.GetEnv("GBT_CAR_BG", "light_blue")
	defaultRootFg := utils.GetEnv("GBT_CAR_FG", "white")
	defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")
	defaultIconBg := defaultRootBg
	defaultIconFg := defaultRootFg
	defaultIconFm := defaultRootFm
	defaultContextBg := defaultRootBg
	defaultContextFg := defaultRootFg
	defaultContextFm := defaultRootFm
	defaultClusterBg := defaultRootBg
	defaultClusterFg := defaultRootFm
	defaultClusterFm := defaultRootFm
	defaultAuthInfoBg := defaultRootBg
	defaultAuthInfoFg := defaultRootFm
	defaultAuthInfoFm := defaultRootFm
	defaultNamespaceBg := defaultRootBg
	defaultNamespaceFg := defaultRootFm
	defaultNamespaceFm := defaultRootFm

	c.Display = utils.GetEnvBool("GBT_CAR_KUBECTL_DISPLAY", isKubectlExists())
	contextInfo := getCurrentContext(c.Display)

	c.Model = map[string]car.ModelElement{
		"root": {
			Bg:   utils.GetEnv("GBT_CAR_KUBECTL_BG", defaultRootBg),
			Fg:   utils.GetEnv("GBT_CAR_KUBECTL_FG", defaultRootFg),
			Fm:   utils.GetEnv("GBT_CAR_KUBECTL_FM", defaultRootFm),
			Text: utils.GetEnv("GBT_CAR_KUBECTL_FORMAT", " {{ Icon }} {{ Context }}|{{ Cluster }} "),
		},
		"Icon": {
			Bg:   utils.GetEnv("GBT_CAR_KUBECTL_ICON_BG", utils.GetEnv("GBT_CAR_KUBECTL_BG", defaultIconBg)),
			Fg:   utils.GetEnv("GBT_CAR_KUBECTL_ICON_FG", utils.GetEnv("GBT_CAR_KUBECTL_FG", defaultIconFg)),
			Fm:   utils.GetEnv("GBT_CAR_KUBECTL_ICON_FM", utils.GetEnv("GBT_CAR_KUBECTL_FM", defaultIconFm)),
			Text: utils.GetEnv("GBT_CAR_KUBECTL_ICON_TEXT", "⎈"),
		},
		"Context": {
			Bg:   utils.GetEnv("GBT_CAR_KUBECTL_CONTEXT_BG", utils.GetEnv("GBT_CAR_KUBECTL_BG", defaultContextBg)),
			Fg:   utils.GetEnv("GBT_CAR_KUBECTL_CONTEXT_FG", utils.GetEnv("GBT_CAR_KUBECTL_FG", defaultContextFg)),
			Fm:   utils.GetEnv("GBT_CAR_KUBECTL_CONTEXT_FM", utils.GetEnv("GBT_CAR_KUBECTL_FM", defaultContextFm)),
			Text: utils.GetEnv("GBT_CAR_KUBECTL_CONTEXT_TEXT", contextInfo.context),
		},
		"Cluster": {
			Bg:   utils.GetEnv("GBT_CAR_KUBECTL_CLUSTER_BG", utils.GetEnv("GBT_CAR_KUBECTL_BG", defaultClusterBg)),
			Fg:   utils.GetEnv("GBT_CAR_KUBECTL_CLUSTER_FG", utils.GetEnv("GBT_CAR_KUBECTL_FG", defaultClusterFg)),
			Fm:   utils.GetEnv("GBT_CAR_KUBECTL_CLUSTER_FM", utils.GetEnv("GBT_CAR_KUBECTL_FM", defaultClusterFm)),
			Text: utils.GetEnv("GBT_CAR_KUBECTL_CLUSTER_TEXT", contextInfo.cluster),
		},
		"AuthInfo": {
			Bg:   utils.GetEnv("GBT_CAR_KUBECTL_AUTHINFO_BG", utils.GetEnv("GBT_CAR_KUBECTL_BG", defaultAuthInfoBg)),
			Fg:   utils.GetEnv("GBT_CAR_KUBECTL_AUTHINFO_FG", utils.GetEnv("GBT_CAR_KUBECTL_FG", defaultAuthInfoFg)),
			Fm:   utils.GetEnv("GBT_CAR_KUBECTL_AUTHINFO_FM", utils.GetEnv("GBT_CAR_KUBECTL_FM", defaultAuthInfoFm)),
			Text: utils.GetEnv("GBT_CAR_KUBECTL_AUTHINFO_TEXT", contextInfo.authInfo),
		},
		"Namespace": {
			Bg:   utils.GetEnv("GBT_CAR_KUBECTL_NAMESPACE_BG", utils.GetEnv("GBT_CAR_KUBECTL_BG", defaultNamespaceBg)),
			Fg:   utils.GetEnv("GBT_CAR_KUBECTL_NAMESPACE_FG", utils.GetEnv("GBT_CAR_KUBECTL_FG", defaultNamespaceFg)),
			Fm:   utils.GetEnv("GBT_CAR_KUBECTL_NAMESPACE_FM", utils.GetEnv("GBT_CAR_KUBECTL_FM", defaultNamespaceFm)),
			Text: utils.GetEnv("GBT_CAR_KUBECTL_NAMESPACE_TEXT", contextInfo.namespace),
		},
	}

	c.Wrap = utils.GetEnvBool("GBT_CAR_KUBECTL_WRAP", false)
	c.Sep = utils.GetEnv("GBT_CAR_KUBECTL_SEP", "\000")
}
