# shellcheck shell=bash
function GbtCarHostname() {
    local defaultRootBg=${GBT_CAR_BG:-dark_gray}
    local defaultRootFg=${GBT_CAR_FG:-252}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local defaultSep="\x00"

    local hostname
    hostname=$(hostname)

    if [ -z "$hostname" ]; then
        hostname='localhost'
    fi

    local uaFormat='{{ User }}'

    if [[ $UID == 0 ]]; then
        uaFormat='{{ Admin }}'
    fi

    # shellcheck disable=SC2034
    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_HOSTNAME_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_HOSTNAME_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_HOSTNAME_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_HOSTNAME_FORMAT-' {{ UserHost }} '}

        [model-UserHost-Bg]=${GBT_CAR_HOSTNAME_USERHOST_BG:-${GBT_CAR_HOSTNAME_BG:-$defaultRootBg}}
        [model-UserHost-Fg]=${GBT_CAR_HOSTNAME_USERHOST_FG:-${GBT_CAR_HOSTNAME_FG:-$defaultRootFg}}
        [model-UserHost-Fm]=${GBT_CAR_HOSTNAME_USERHOST_FM:-${GBT_CAR_HOSTNAME_FM:-$defaultRootFm}}
        [model-UserHost-Text]=${GBT_CAR_HOSTNAME_USERHOST_FORMAT-"$uaFormat@{{ Host }}"}

        [model-Admin-Bg]=${GBT_CAR_HOSTNAME_ADMIN_BG:-${GBT_CAR_HOSTNAME_USERHOST_BG:-${GBT_CAR_HOSTNAME_BG:-$defaultRootBg}}}
        [model-Admin-Fg]=${GBT_CAR_HOSTNAME_ADMIN_FG:-${GBT_CAR_HOSTNAME_USERHOST_FG:-${GBT_CAR_HOSTNAME_FG:-$defaultRootFg}}}
        [model-Admin-Fm]=${GBT_CAR_HOSTNAME_ADMIN_FM:-${GBT_CAR_HOSTNAME_USERHOST_FM:-${GBT_CAR_HOSTNAME_FM:-$defaultRootFm}}}
        [model-Admin-Text]=${GBT_CAR_HOSTNAME_ADMIN_TEXT-${USER:-$UID}}

        [model-User-Bg]=${GBT_CAR_HOSTNAME_USER_BG:-${GBT_CAR_HOSTNAME_USERHOST_BG:-${GBT_CAR_HOSTNAME_BG:-$defaultRootBg}}}
        [model-User-Fg]=${GBT_CAR_HOSTNAME_USER_FG:-${GBT_CAR_HOSTNAME_USERHOST_FG:-${GBT_CAR_HOSTNAME_FG:-$defaultRootFg}}}
        [model-User-Fm]=${GBT_CAR_HOSTNAME_USER_FM:-${GBT_CAR_HOSTNAME_USERHOST_FM:-${GBT_CAR_HOSTNAME_FM:-$defaultRootFm}}}
        [model-User-Text]=${GBT_CAR_HOSTNAME_USER_TEXT-${USER:-$UID}}

        [model-Host-Bg]=${GBT_CAR_HOSTNAME_HOST_BG:-${GBT_CAR_HOSTNAME_BG:-$defaultRootBg}}
        [model-Host-Fg]=${GBT_CAR_HOSTNAME_HOST_FG:-${GBT_CAR_HOSTNAME_FG:-$defaultRootFg}}
        [model-Host-Fm]=${GBT_CAR_HOSTNAME_HOST_FM:-${GBT_CAR_HOSTNAME_FM:-$defaultRootFm}}
        [model-Host-Text]=${GBT_CAR_HOSTNAME_HOST_TEXT-${hostname%%.*}}

        [model-Sep-Bg]=${GBT_CAR_HOSTNAME_SEP_BG:-$defaultSep}
        [model-Sep-Fg]=${GBT_CAR_HOSTNAME_SEP_FG:-$defaultSep}
        [model-Sep-Fm]=${GBT_CAR_HOSTNAME_SEP_FM:-$defaultSep}
        [model-Sep-Text]=${GBT_CAR_HOSTNAME_SEP_TEXT:-${GBT_CAR_HOSTNAME_SEP:-${GBT_SEPARATOR:-$defaultSep}}}

        [display]=${GBT_CAR_HOSTNAME_DISPLAY:-1}
        [wrap]=${GBT_CAR_HOSTNAME_WRAP:-0}
    )
}
