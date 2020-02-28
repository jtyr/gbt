function GbtCarPyVirtEnv() {
    local defaultRootBg=${GBT_CAR_BG:-222}
    local defaultRootFg=${GBT_CAR_FG:-black}
    local defaultRootFm=${GBT_CAR_FM:-none}

    GbtDecorateUnicode ${GBT_CAR_PYVIRTENV_NAME_TEXT-'\xee\x9c\xbc'}
    local defaultIconText=$GBT__RETVAL

    local defaultSep="\x00"

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_PYVIRTENV_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_PYVIRTENV_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_PYVIRTENV_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_PYVIRTENV_FORMAT-' {{ Icon }} {{ Name }} '}

        [model-Icon-Bg]=${GBT_CAR_PYVIRTENV_ICON_BG:-${GBT_CAR_PYVIRTENV_BG:-$defaultRootBg}}
        [model-Icon-Fg]=${GBT_CAR_PYVIRTENV_ICON_FG:-${GBT_CAR_PYVIRTENV_FG:-33}}
        [model-Icon-Fm]=${GBT_CAR_PYVIRTENV_ICON_FM:-${GBT_CAR_PYVIRTENV_FM:-$defaultRootFm}}
        [model-Icon-Text]=$defaultIconText

        [model-Name-Bg]=${GBT_CAR_PYVIRTENV_NAME_BG:-${GBT_CAR_PYVIRTENV_BG:-$defaultRootBg}}
        [model-Name-Fg]=${GBT_CAR_PYVIRTENV_NAME_FG:-${GBT_CAR_PYVIRTENV_FG:-$defaultRootFg}}
        [model-Name-Fm]=${GBT_CAR_PYVIRTENV_NAME_FM:-${GBT_CAR_PYVIRTENV_FM:-$defaultRootFm}}
        [model-Name-Text]=${GBT_CAR_PYVIRTENV_NAME_TEXT-${VIRTUAL_ENV##/*/}}

        [model-Sep-Bg]=${GBT_CAR_PYVIRTENV_SEP_BG:-$defaultSep}
        [model-Sep-Fg]=${GBT_CAR_PYVIRTENV_SEP_FG:-$defaultSep}
        [model-Sep-Fm]=${GBT_CAR_PYVIRTENV_SEP_FM:-$defaultSep}
        [model-Sep-Text]=${GBT_CAR_PYVIRTENV_SEP_TEXT:-${GBT_CAR_PYVIRTENV_SEP:-$defaultSep}}

        [wrap]=${GBT_CAR_PYVIRTENV_WRAP:-0}
    )

    if [ -n "$VIRTUAL_ENV" ]; then
        GBT_CAR[display]=${GBT_CAR_PYVIRTENV_DISPLAY:-1}
    else
        GBT_CAR[display]=${GBT_CAR_PYVIRTENV_DISPLAY:-0}
    fi
}
