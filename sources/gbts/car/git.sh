function GbtCarGit() {
    if [[ $GBT_CAR_GIT_DISPLAY == 0 ]]; then
        return
    fi

    local defaultRootBg=${GBT_CAR_BG:-light_gray}
    local defaultRootFg=${GBT_CAR_FG:-black}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local defaultHeadText=''
    local defaultStatusFormat='{{ Clean }}'
    local defaultAheadText=''
    local defaultBehindText=''

    GbtDecorateUnicode ${GBT_CAR_GIT_ICON_TEXT:-'\xee\x82\xa0'}
    local defaultIconText=$GBT__RETVAL
    GbtDecorateUnicode ${GBT_CAR_GIT_DIRTY_TEXT:-'\xe2\x9c\x98'}
    local defaultDirtyText=$GBT__RETVAL
    GbtDecorateUnicode ${GBT_CAR_GIT_CLEAN_TEXT:-'\xe2\x9c\x94'}
    local defaultCleanText=$GBT__RETVAL

    local isGitDir=0

    git rev-parse --git-dir 1>/dev/null 2>/dev/null

    if [[ $? == 0 ]]; then
        isGitDir=1
    fi

    if [[ $isGitDir == 1 ]]; then
        defaultHeadText=$(git symbolic-ref HEAD 2>/dev/null)

        if [[ -z "$defaultHeadText" ]]; then
            defaultHeadText=$(git describe --tags --exact-match HEAD 2>/dev/null)

            if [[ -z "$defaultHeadText" ]]; then
                defaultHeadText=$(git rev-parse --short HEAD 2>/dev/null)
            fi
        fi

        defaultHeadText=${defaultHeadText#refs/heads/}

        local dirty=$(git status --porcelain 2>/dev/null)

        if [ -n "$dirty" ]; then
            defaultStatusFormat='{{ Dirty }}'
        fi

        local ahead=$(git rev-list --count HEAD..@{upstream} 2>/dev/null || echo E)

        if [[ $ahead != 0 ]] && [[ $ahead != 'E' ]]; then
            GbtDecorateUnicode ${GBT_CAR_GIT_AHEAD_SYMBOL:-' \xe2\xac\x86'}
            defaultAheadText=$GBT__RETVAL
        fi

        local behind=$(git rev-list --count @{upstream}..HEAD 2>/dev/null || echo E)

        if [[ $behind != 0 ]] && [[ $behind != 'E' ]]; then
            GbtDecorateUnicode ${GBT_CAR_GIT_BEHIND_SYMBOL:-' \xe2\xac\x87'}
            defaultBehindText=$GBT__RETVAL
        fi
    fi

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_GIT_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_GIT_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_GIT_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_GIT_FORMAT:-' {{ Icon }} {{ Head }} {{ Status }}{{ Ahead }}{{ Behind }} '}

        [model-Icon-Bg]=${GBT_CAR_GIT_ICON_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Icon-Fg]=${GBT_CAR_GIT_ICON_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Icon-Fm]=${GBT_CAR_GIT_ICON_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Icon-Text]=$defaultIconText

        [model-Head-Bg]=${GBT_CAR_GIT_HEAD_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Head-Fg]=${GBT_CAR_GIT_HEAD_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Head-Fm]=${GBT_CAR_GIT_HEAD_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Head-Text]=${GBT_CAR_GIT_HEAD_TEXT:-$defaultHeadText}

        [model-Status-Bg]=${GBT_CAR_GIT_STATUS_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Status-Fg]=${GBT_CAR_GIT_STATUS_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Status-Fm]=${GBT_CAR_GIT_STATUS_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Status-Text]=${GBT_CAR_GIT_STATUS_FORMAT:-$defaultStatusFormat}

        [model-Dirty-Bg]=${GBT_CAR_GIT_DIRTY_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Dirty-Fg]=${GBT_CAR_GIT_DIRTY_FG:-${GBT_CAR_GIT_FG:-red}}
        [model-Dirty-Fm]=${GBT_CAR_GIT_DIRTY_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Dirty-Text]=$defaultDirtyText

        [model-Clean-Bg]=${GBT_CAR_GIT_CLEAN_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Clean-Fg]=${GBT_CAR_GIT_CLEAN_FG:-${GBT_CAR_GIT_FG:-green}}
        [model-Clean-Fm]=${GBT_CAR_GIT_CLEAN_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Clean-Text]=$defaultCleanText

        [model-Ahead-Bg]=${GBT_CAR_GIT_AHEAD_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Ahead-Fg]=${GBT_CAR_GIT_AHEAD_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Ahead-Fm]=${GBT_CAR_GIT_AHEAD_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Ahead-Text]=${GBT_CAR_GIT_AHEAD_TEXT:-$defaultAheadText}

        [model-Behind-Bg]=${GBT_CAR_GIT_BEHIND_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Behind-Fg]=${GBT_CAR_GIT_BEHIND_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Behind-Fm]=${GBT_CAR_GIT_BEHIND_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Behind-Text]=${GBT_CAR_GIT_BEHIND_TEXT:-$defaultBehindText}

        [display]=${GBT_CAR_GIT_DISPLAY:-$isGitDir}
        [wrap]=${GBT_CAR_GIT_WRAP:-0}
        [sep]=${GBT_CAR_GIT_SEP:-'\x00'}
    )
}
