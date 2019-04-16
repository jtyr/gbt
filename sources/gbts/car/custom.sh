function GbtCarCustom() {
    local name=$(echo $1 | tr '[:lower:]' '[:upper:]')

    local prefix="GBT_CAR_CUSTOM$name"

    local C_TEXT_CMD="${prefix}_TEXT_CMD"
    local C_DISPLAY_CMD="${prefix}_DISPLAY_CMD"
    local C_BG="${prefix}_BG"
    local C_FG="${prefix}_FG"
    local C_FM="${prefix}_FM"
    local C_FORMAT="${prefix}_FORMAT"
    local C_TEXT_BG="${prefix}_TEXT_BG"
    local C_TEXT_FG="${prefix}_TEXT_FG"
    local C_TEXT_FM="${prefix}_TEXT_FM"
    local C_TEXT_TEXT="${prefix}_TEXT_TEXT"
    local C_DISPLAY="${prefix}_DISPLAY"
    local C_WRAP="${prefix}_WRAP"
    local C_SEP="${prefix}_SEP"

    local defaultRootBg=${GBT_CAR_BG:-yellow}
    local defaultRootFg=${GBT_CAR_FG:-default}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local defaultTextText='?'
    local defaultTextCmd=${!C_TEXT_CMD}
    local defaultDisplayCmd=${!C_DISPLAY_CMD}
    local defaultDisplay=1

    if [ -n "$defaultTextCmd" ]; then
        defaultTextText=$(sh -c "$defaultTextCmd")
    fi

    if [ -n "$defaultDisplayCmd" ]; then
        local defaultDisplayOutput=$(sh -c "$defaultDisplayCmd")

        if [[ ! $defaultDisplayOutput =~ ^([Yy][Ee][Ss]|[Tt][Rr][Uu][Ee]|1)$ ]]; then
            defaultDisplay=0
        fi
    fi

    GBT_CAR=(
        [model-root-Bg]=${!C_BG:-$defaultRootBg}
        [model-root-Fg]=${!C_FG:-$defaultRootFg}
        [model-root-Fm]=${!C_FM:-$defaultRootFm}
        [model-root-Text]=${!C_FORMAT-' {{ Text }} '}
        [model-Text-Bg]=${!C_TEXT_BG:-${!C_BG:-$defaultRootBg}}
        [model-Text-Fg]=${!C_TEXT_FG:-${!C_FG:-$defaultRootFg}}
        [model-Text-Fm]=${!C_TEXT_FM:-${!C_FM:-$defaultRootFm}}
        [model-Text-Text]=${!C_TEXT_TEXT-$defaultTextText}

        [display]=${!C_DISPLAY:-$defaultDisplay}
        [wrap]=${!C_WRAP:-0}
        [sep]=${!C_SEP-'\x00'}
    )
}
