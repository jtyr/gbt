function GbtCarAws() {
    if [[ $GBT_CAR_AWS_DISPLAY == 0 ]]; then
        return
    fi

    local defaultRootBg=${GBT_CAR_BG:-180;85;10}
    local defaultRootFg=${GBT_CAR_FG:-white}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local defaultRootFormat=${GBT_CAR_AWS_FORMAT-' {{ Icon }} {{ Profile }} '}
    local defaultProfileText=${AWS_PROFILE:-default}
    local defaultRegionText=$AWS_DEFAULT_REGION
    local defaultSep="\x00"

    configFile="$HOME/.aws/config"

    profileSection=$defaultProfileText

    if [[ $defaultProfileText != 'default' ]]; then
        profileSection="profile $defaultProfileText"
    fi

    defaultRegionText=$(sed -nr "/^\[$profileSection\]/ { :l /^region[ ]*=/ { s/.*=[ ]*//; p; q;}; n; b l;}" $configFile)

    GbtDecorateUnicode ${GBT_CAR_AWS_ICON_TEXT-'\xef\x94\xad'}
    local defaultIconText=$GBT__RETVAL

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_AWS_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_AWS_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_AWS_FM:-$defaultRootFm}
        [model-root-Text]=$defaultRootFormat

        [model-Icon-Bg]=${GBT_CAR_AWS_ICON_BG:-${GBT_CAR_AWS_BG:-$defaultRootBg}}
        [model-Icon-Fg]=${GBT_CAR_AWS_ICON_FG:-${GBT_CAR_AWS_FG:-$defaultRootFg}}
        [model-Icon-Fm]=${GBT_CAR_AWS_ICON_FM:-${GBT_CAR_AWS_FM:-$defaultRootFm}}
        [model-Icon-Text]=$defaultIconText

        [model-Profile-Bg]=${GBT_CAR_AWS_PROFILE_BG:-${GBT_CAR_AWS_BG:-$defaultRootBg}}
        [model-Profile-Fg]=${GBT_CAR_AWS_PROFILE_FG:-${GBT_CAR_AWS_FG:-$defaultRootFg}}
        [model-Profile-Fm]=${GBT_CAR_AWS_PROFILE_FM:-${GBT_CAR_AWS_FM:-$defaultRootFm}}
        [model-Profile-Text]=${GBT_CAR_AWS_PROFILE_TEXT-$defaultProfileText}

        [model-Region-Bg]=${GBT_CAR_AWS_REGION_BG:-${GBT_CAR_AWS_BG:-$defaultRootBg}}
        [model-Region-Fg]=${GBT_CAR_AWS_REGION_FG:-${GBT_CAR_AWS_FG:-$defaultRootFg}}
        [model-Region-Fm]=${GBT_CAR_AWS_REGION_FM:-${GBT_CAR_AWS_FM:-$defaultRootFm}}
        [model-Region-Text]=${GBT_CAR_AWS_REGION_TEXT-$defaultRegionText}

        [model-Sep-Bg]=${GBT_CAR_AWS_SEP_BG:-$defaultSep}
        [model-Sep-Fg]=${GBT_CAR_AWS_SEP_FG:-$defaultSep}
        [model-Sep-Fm]=${GBT_CAR_AWS_SEP_FM:-$defaultSep}
        [model-Sep-Text]=${GBT_CAR_AWS_SEP_TEXT:-${GBT_CAR_AWS_SEP:-${GBT_SEPARATOR:-$defaultSep}}}

        [display]=${GBT_CAR_AWS_DISPLAY:-1}
        [wrap]=${GBT_CAR_AWS_WRAP:-0}
    )
}
