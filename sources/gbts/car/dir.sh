function GbtCarDir() {
    local text=''

    local homesign=${GBT_CAR_DIR_HOMESIGN-'~'}
    local dirsep=${GBT_CAR_DIR_DIRSEP-/}

    if [ -n "$homesign" ] && [[ $PWD == $HOME ]]; then
        text=$homesign
    elif [[ $PWD == '/' ]]; then
        text=$dirsep
    elif [[ $PWD == '//' ]]; then
        text=$dirsep$dirsep
    else
        local first=1
        local cur=1
        local pwd=$PWD
        local pwd_len=0

        local depth=${GBT_CAR_DIR_DEPTH:-1}

        if [[ -n $homesign ]]; then
            pwd=${pwd/$HOME/$homesign}
        fi

        for d in ${pwd//\// }; do
            pwd_len=$(( pwd_len + 1 ))
        done

        for d in ${pwd//\// }; do
            if (( $pwd_len - $depth < $cur )); then
                if (( $pwd_len <= $depth )); then
                    first=0
                fi

                if [[ $first != 1 ]] && [[ $d != $homesign ]]; then
                    text+=$dirsep
                fi

                if (( $cur < $pwd_len )); then
                    text+=${d::${GBT_CAR_DIR_NONCURLEN:-255}}
                else
                    text+=$d
                fi

                first=0
            fi

            cur=$(( $cur + 1 ))
        done
    fi

    local defaultRootBg=${GBT_CAR_BG:-blue}
    local defaultRootFg=${GBT_CAR_FG:-light_gray}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local defaultSep="\x00"

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_DIR_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_DIR_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_DIR_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_DIR_FORMAT-' {{ Dir }} '}

        [model-Dir-Bg]=${GBT_CAR_DIR_DIR_BG:-${GBT_CAR_DIR_BG:-$defaultRootBg}}
        [model-Dir-Fg]=${GBT_CAR_DIR_DIR_FG:-${GBT_CAR_DIR_FG:-$defaultRootFg}}
        [model-Dir-Fm]=${GBT_CAR_DIR_DIR_FM:-${GBT_CAR_DIR_FM:-$defaultRootFm}}
        [model-Dir-Text]=${GBT_CAR_DIR_DIR_TEXT-$text}

        [model-Sep-Bg]=${GBT_CAR_DIR_SEP_BG:-$defaultSep}
        [model-Sep-Fg]=${GBT_CAR_DIR_SEP_FG:-$defaultSep}
        [model-Sep-Fm]=${GBT_CAR_DIR_SEP_FM:-$defaultSep}
        [model-Sep-Text]=${GBT_CAR_DIR_SEP_TEXT:-${GBT_CAR_DIR_SEP:-${GBT_SEPARATOR:-$defaultSep}}}

        [display]=${GBT_CAR_DIR_DISPLAY:-1}
        [wrap]=${GBT_CAR_DIR_WRAP:-0}
    )
}
