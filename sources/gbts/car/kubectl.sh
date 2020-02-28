function GbtCarKubectl() {
    local defaultRootBg=${GBT_CAR_BG:-26}
    local defaultRootFg=${GBT_CAR_FG:-white}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local defaultSep="\x00"

    local isKubectlCurrentContextSet=0
    local output=$(kubectl config current-context 2>/dev/null)

    if [[ $? == 0 ]] && [ -n "$output" ]; then
        isKubectlCurrentContextSet=1

        declare -A contextInfo
        local i=0
        for N in $(kubectl config get-contexts | grep '^*' | sed -E 's/^\*\ +//'); do
            if [[ $i == 0 ]]; then
                contextInfo[context]=$N
            elif [[ $i == 1 ]]; then
                contextInfo[cluster]=$N
            elif [[ $i == 2 ]]; then
                contextInfo[authInfo]=$N
            elif [[ $i == 3 ]]; then
                contextInfo[namespace]=$N
            fi

            i=$(( i + 1 ))
        done
    fi

    GbtDecorateUnicode ${GBT_CAR_KUBECTL_ICON_TEXT-'\xe2\x8e\x88'}
    local defaultIconText=$GBT__RETVAL

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_KUBECTL_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_KUBECTL_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_KUBECTL_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_KUBECTL_FORMAT-' {{ Icon }} {{ Context }} '}

        [model-Icon-Bg]=${GBT_CAR_KUBECTL_ICON_BG:-${GBT_CAR_KUBECTL_BG:-$defaultRootBg}}
        [model-Icon-Fg]=${GBT_CAR_KUBECTL_ICON_FG:-${GBT_CAR_KUBECTL_FG:-$defaultRootFg}}
        [model-Icon-Fm]=${GBT_CAR_KUBECTL_ICON_FM:-${GBT_CAR_KUBECTL_FM:-$defaultRootFm}}
        [model-Icon-Text]=$defaultIconText

        [model-Context-Bg]=${GBT_CAR_KUBECTL_CONTEXT_BG:-${GBT_CAR_KUBECTL_BG:-$defaultRootBg}}
        [model-Context-Fg]=${GBT_CAR_KUBECTL_CONTEXT_FG:-${GBT_CAR_KUBECTL_FG:-$defaultRootFg}}
        [model-Context-Fm]=${GBT_CAR_KUBECTL_CONTEXT_FM:-${GBT_CAR_KUBECTL_FM:-$defaultRootFm}}
        [model-Context-Text]=${GBT_CAR_KUBECTL_CONTEXT_TEXT-${contextInfo[context]}}

        [model-Cluster-Bg]=${GBT_CAR_KUBECTL_CLUSTER_BG:-${GBT_CAR_KUBECTL_BG:-$defaultRootBg}}
        [model-Cluster-Fg]=${GBT_CAR_KUBECTL_CLUSTER_FG:-${GBT_CAR_KUBECTL_FG:-$defaultRootFg}}
        [model-Cluster-Fm]=${GBT_CAR_KUBECTL_CLUSTER_FM:-${GBT_CAR_KUBECTL_FM:-$defaultRootFm}}
        [model-Cluster-Text]=${GBT_CAR_KUBECTL_CLUSTER_TEXT-${contextInfo[cluster]}}

        [model-AuthInfo-Bg]=${GBT_CAR_KUBECTL_AUTHINFO_BG:-${GBT_CAR_KUBECTL_BG:-$defaultRootBg}}
        [model-AuthInfo-Fg]=${GBT_CAR_KUBECTL_AUTHINFO_FG:-${GBT_CAR_KUBECTL_FG:-$defaultRootFg}}
        [model-AuthInfo-Fm]=${GBT_CAR_KUBECTL_AUTHINFO_FM:-${GBT_CAR_KUBECTL_FM:-$defaultRootFm}}
        [model-AuthInfo-Text]=${GBT_CAR_KUBECTL_AUTHINFO_TEXT-${contextInfo[authInfo]}}

        [model-Namespace-Bg]=${GBT_CAR_KUBECTL_NAMESPACE_BG:-${GBT_CAR_KUBECTL_BG:-$defaultRootBg}}
        [model-Namespace-Fg]=${GBT_CAR_KUBECTL_NAMESPACE_FG:-${GBT_CAR_KUBECTL_FG:-$defaultRootFg}}
        [model-Namespace-Fm]=${GBT_CAR_KUBECTL_NAMESPACE_FM:-${GBT_CAR_KUBECTL_FM:-$defaultRootFm}}
        [model-Namespace-Text]=${GBT_CAR_KUBECTL_NAMESPACE_TEXT-${contextInfo[namespace]}}

        [model-Sep-Bg]=${GBT_CAR_KUBECTL_SEP_BG:-$defaultSep}
        [model-Sep-Fg]=${GBT_CAR_KUBECTL_SEP_FG:-$defaultSep}
        [model-Sep-Fm]=${GBT_CAR_KUBECTL_SEP_FM:-$defaultSep}
        [model-Sep-Text]=${GBT_CAR_KUBECTL_SEP_TEXT:-${GBT_CAR_KUBECTL_SEP:-$defaultSep}}

        [display]=${GBT_CAR_KUBECTL_DISPLAY:-$isKubectlCurrentContextSet}
        [wrap]=${GBT_CAR_KUBECTL_WRAP:-0}
    )
}
