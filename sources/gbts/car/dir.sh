function GbtCarDir() {
    local text=''
    local pwd_len=0

    for d in ${PWD//\// }; do
        pwd_len=$(( pwd_len + 1 ))
    done

    local homesign=${GBT_CAR_DIR_HOMESIGN:-'~'}

    if [ -n "$homesign" ] && [[ $PWD == $HOME ]]; then
        text=$homesign
    elif [[ $PWD == '/' ]]; then
        text=${GBT_CAR_DIR_DIRSEP-/}
    elif [[ $PWD == '//' ]]; then
        text=${GBT_CAR_DIR_DIRSEP-/}${GBT_CAR_DIR_DIRSEP-/}
    else
        local first=1
        local cur=1

        for d in ${PWD//\// }; do
            if (( $pwd_len - ${GBT_CAR_DIR_DEPTH:-1} < $cur )); then
                if (( $pwd_len <= ${GBT_CAR_DIR_DEPTH:-1} )); then
                    first=0
                fi

                if [[ $first != 1 ]]; then
                    text+=${GBT_CAR_DIR_DIRSEP-/}
                fi

                text+=$d

                first=0
            fi

            cur=$(($cur + 1))
        done
    fi

    local defaultRootBg=${GBT_CAR_BG:-blue}
    local defaultRootFg=${GBT_CAR_FG:-light_gray}
    local defaultRootFm=${GBT_CAR_FM:-none}

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_DIR_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_DIR_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_DIR_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_DIR_FORMAT-' {{ Dir }} '}

        [model-Dir-Bg]=${GBT_CAR_DIR_DIR_BG:-${GBT_CAR_DIR_BG:-$defaultRootBg}}
        [model-Dir-Fg]=${GBT_CAR_DIR_DIR_FG:-${GBT_CAR_DIR_FG:-$defaultRootFg}}
        [model-Dir-Fm]=${GBT_CAR_DIR_DIR_FM:-${GBT_CAR_DIR_FM:-$defaultRootFm}}
        [model-Dir-Text]=${GBT_CAR_DIR_DIR_TEXT-$text}

        [display]=${GBT_CAR_DIR_DISPLAY:-1}
        [wrap]=${GBT_CAR_DIR_WRAP:-0}
        [sep]=${GBT_CAR_DIR_SEP-'\x00'}
    )
}
