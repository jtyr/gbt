function GbtCarSign() {
    local symbolFormat='{{ User }}'

    if [[ $UID == 0 ]]; then
        symbolFormat='{{ Admin }}'
    fi

    local defaultRootBg=${GBT_CAR_BG:-default}
    local defaultRootFg=${GBT_CAR_FG:-default}
    local defaultRootFm=${GBT_CAR_FM:-none}

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_SIGN_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_SIGN_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_SIGN_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_SIGN_FORMAT-' {{ Symbol }} '}

        [model-Symbol-Bg]=${GBT_CAR_SIGN_SYMBOL_BG:-${GBT_CAR_SIGN_BG:-$defaultRootBg}}
        [model-Symbol-Fg]=${GBT_CAR_SIGN_SYMBOL_FG:-${GBT_CAR_SIGN_FG:-green}}
        [model-Symbol-Fm]=${GBT_CAR_SIGN_SYMBOL_FM:-${GBT_CAR_SIGN_FM:-bold}}
        [model-Symbol-Text]=${GBT_CAR_SIGN_SYMBOL_FORMAT-$symbolFormat}

        [model-User-Bg]=${GBT_CAR_SIGN_USER_BG:-${GBT_CAR_SIGN_BG:-$defaultRootBg}}
        [model-User-Fg]=${GBT_CAR_SIGN_USER_FG:-light_green}
        [model-User-Fm]=${GBT_CAR_SIGN_USER_FM:-${GBT_CAR_SIGN_FM:-bold}}
        [model-User-Text]=${GBT_CAR_SIGN_USER_TEXT-'$'}

        [model-Admin-Bg]=${GBT_CAR_SIGN_ADMIN_BG:-${GBT_CAR_SIGN_BG:-$defaultRootBg}}
        [model-Admin-Fg]=${GBT_CAR_SIGN_ADMIN_FG:-red}
        [model-Admin-Fm]=${GBT_CAR_SIGN_ADMIN_FM:-${GBT_CAR_SIGN_FM:-bold}}
        [model-Admin-Text]=${GBT_CAR_SIGN_ADMIN_TEXT-'#'}

        [display]=${GBT_CAR_SIGN_DISPLAY:-1}
        [wrap]=${GBT_CAR_SIGN_WRAP:-0}
        [sep]=${GBT_CAR_SIGN_SEP-'\x00'}
    )
}
