function GbtCarGit() {
    if [[ $GBT_CAR_GIT_DISPLAY == 0 ]]; then
        return
    fi

    local defaultRootBg=${GBT_CAR_BG:-light_gray}
    local defaultRootFg=${GBT_CAR_FG:-black}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local defaultRootFormat=${GBT_CAR_GIT_FORMAT:-' {{ Icon }} {{ Head }} {{ Status }}{{ Ahead }}{{ Behind }} '}
    local defaultHeadText=''
    local defaultStatusFormat='{{ StatusClean }}'
    local defaultStatusAddedCountText=''
    local defaultStatusAddedSymbolText=''
    local defaultStatusCopiedCountText=''
    local defaultStatusCopiedSymbolText=''
    local defaultStatusDeletedCountText=''
    local defaultStatusDeletedSymbolText=''
    local defaultStatusIgnoredCountText=''
    local defaultStatusIgnoredSymbolText=''
    local defaultStatusModifiedCountText=''
    local defaultStatusModifiedSymbolText=''
    local defaultStatusRenamedCountText=''
    local defaultStatusRenamedSymbolText=''
    local defaultStatusStagedCountText=''
    local defaultStatusStagedSymbolText=''
    local defaultStatusUnmergedCountText=''
    local defaultStatusUnmergedSymbolText=''
    local defaultStatusUntrackedCountText=''
    local defaultStatusUntrackedSymbolText=''
    local defaultAheadCountText=''
    local defaultAheadSymbolText=''
    local defaultBehindCountText=''
    local defaultBehindSymbolText=''
    local defaultStashCountText=''
    local defaultStashSymbolText=''

    GbtDecorateUnicode ${GBT_CAR_GIT_ICON_TEXT:-'\xee\x82\xa0'}
    local defaultIconText=$GBT__RETVAL
    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_DIRTY_TEXT:-'\xe2\x9c\x98'}
    local defaultStatusDirtyText=$GBT__RETVAL
    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_CLEAN_TEXT:-'\xe2\x9c\x94'}
    local defaultStatusCleanText=$GBT__RETVAL

    local isGitDir=0

    git rev-parse --git-dir 1>/dev/null 2>/dev/null

    if [[ $? == 0 ]]; then
        isGitDir=1

        if [[ $defaultRootFormat =~ \{\{\ *Head\ *\}\} ]]; then
            defaultHeadText=$(git symbolic-ref HEAD 2>/dev/null)

            if [[ -z "$defaultHeadText" ]]; then
                defaultHeadText=$(git describe --tags --exact-match HEAD 2>/dev/null)

                if [[ -z "$defaultHeadText" ]]; then
                    defaultHeadText=$(git rev-parse --short HEAD 2>/dev/null)
                fi
            fi

            defaultHeadText=${defaultHeadText#refs/heads/}
        fi

        if [[ $defaultRootFormat =~ \{\{\ *Status.*\ *\}\} ]]; then
            declare -a status

            local IFS='\n'
            for line in $(git status --porcelain 2>/dev/null); do
                case "${line:1:1}" in
                    A)
                        ((status[added]++)) ;;
                    C)
                        ((status[copied]++)) ;;
                    D)
                        ((status[deleted]++)) ;;
                    !)
                        ((status[ignored]++)) ;;
                    M)
                        ((status[modified]++)) ;;
                    R)
                        ((status[renamed]++)) ;;
                    ' ')
                        ((status[staged]++)) ;;
                    U)
                        ((status[unmerged]++)) ;;
                    ?)
                        ((status[untracked]++)) ;;
                esac
            done

            if [ ${#status[@]} -gt 0 ]; then
                defaultStatusFormat='{{ StatusDirty }}'

                if [ ${status[added]} -gt 0 ]; then
                    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_ADDED_SYMBOL_TEXT:-' \xe2\x9f\xb4'}
                    defaultStatusAddedSymbolText=$GBT__RETVAL
                    defaultStatusAddedCountText=${GBT_CAR_GIT_STATUS_ADDED_COUNT_TEXT-${status[added]}}
                fi

                if [ ${status[copied]} -gt 0 ]; then
                    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_COPIED_SYMBOL_TEXT-' \xe2\xa5\x88'}
                    defaultStatusCopiedSymbolText=$GBT__RETVAL
                    defaultStatusCopiedCountText=${GBT_CAR_GIT_STATUS_COPIED_COUNT_TEXT-${status[copied]}}
                fi

                if [ ${status[deleted]} -gt 0 ]; then
                    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_DELETED_SYMBOL_TEXT-' \xe2\x9e\x96'}
                    defaultStatusDeletedSymbolText=$GBT__RETVAL
                    defaultStatusDeletedCountText=${GBT_CAR_GIT_STATUS_DELETED_COUNT_TEXT-${status[deleted]}}
                fi

                if [ ${status[ignored]} -gt 0 ]; then
                    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_IGNORED_SYMBOL_TEXT-' \xe2\x97\x8b'}
                    defaultStatusIgnoredSymbolText=$GBT__RETVAL
                    defaultStatusIgnoredCountText=${GBT_CAR_GIT_STATUS_IGNORED_COUNT_TEXT-${status[ignored]}}
                fi

                if [ ${status[modified]} -gt 0 ]; then
                    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_MODIFIED_SYMBOL_TEXT-' \xe2\x9c\x9a'}
                    defaultStatusModifiedSymbolText=$GBT__RETVAL
                    defaultStatusModifiedCountText=${GBT_CAR_GIT_STATUS_MODIFIED_COUNT_TEXT-${status[modified]}}
                fi

                if [ ${status[renamed]} -gt 0 ]; then
                    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_RENAMED_SYMBOL_TEXT-' \xe2\xa5\xb2'}
                    defaultStatusRenamedSymbolText=$GBT__RETVAL
                    defaultStatusRenamedCountText=${GBT_CAR_GIT_STATUS_RENAMED_COUNT_TEXT-${status[renamed]}}
                fi

                if [ ${status[staged]} -gt 0 ]; then
                    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_STAGED_SYMBOL_TEXT-' \xe2\x97\x8f'}
                    defaultStatusStagedSymbolText=$GBT__RETVAL
                    defaultStatusStagedCountText=${GBT_CAR_GIT_STATUS_STAGED_COUNT_TEXT-${status[staged]}}
                fi

                if [ ${status[unmerged]} -gt 0 ]; then
                    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_UNMERGED_SYMBOL_TEXT-' \xe2\x9c\x96'}
                    defaultStatusUnmergedSymbolText=$GBT__RETVAL
                    defaultStatusUnmergedCountText=${GBT_CAR_GIT_STATUS_UNMERGED_COUNT_TEXT-${status[unmerged]}}
                fi

                if [ ${status[untracked]} -gt 0 ]; then
                    GbtDecorateUnicode ${GBT_CAR_GIT_STATUS_UNTRACKED_SYMBOL_TEXT-' \xe2\x80\xa6'}
                    defaultStatusUntrackedSymbolText=$GBT__RETVAL
                    defaultStatusUntrackedCountText=${GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_TEXT-${status[untracked]}}
                fi
            fi
        fi

        if [[ $defaultRootFormat =~ \{\{\ *Ahead.*\ *\}\} ]]; then
            local ahead=$(git rev-list --count @{upstream}..HEAD 2>/dev/null || echo E)

            if [[ $ahead != 0 ]] && [[ $ahead != 'E' ]]; then
                GbtDecorateUnicode ${GBT_CAR_GIT_AHEAD_SYMBOL:-' \xe2\xac\x86'}
                defaultAheadSymbolText=$GBT__RETVAL
                defaultAheadCountText=${GBT_CAR_GIT_AHEAD_COUNT_TEXT-$ahead}
            fi
        fi

        if [[ $defaultRootFormat =~ \{\{\ *Behind.*\ *\}\} ]]; then
            local behind=$(git rev-list --count HEAD..@{upstream} 2>/dev/null || echo E)

            if [[ $behind != 0 ]] && [[ $behind != 'E' ]]; then
                GbtDecorateUnicode ${GBT_CAR_GIT_BEHIND_SYMBOL:-' \xe2\xac\x87'}
                defaultBehindSymbolText=$GBT__RETVAL
                defaultBehindCountText=${GBT_CAR_GIT_BEHIND_COUNT_TEXT-$behind}
            fi
        fi

        if [[ $defaultRootFormat =~ \{\{\ *Stash.*\ *\}\} ]]; then
            local stash=$(git stash list 2>/dev/null | wc -l)

            if [[ $stash != 0 ]]; then
                GbtDecorateUnicode ${GBT_CAR_GIT_STASH_SYMBOL_TEXT-' \xe2\x9a\x91'}
                defaultStashSymbolText=$GBT__RETVAL
                defaultStashCountText=${GBT_CAR_GIT_STASH_COUNT_TEXT-$stash}
            fi
        fi
    fi

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_GIT_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_GIT_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_GIT_FM:-$defaultRootFm}
        [model-root-Text]=$defaultRootFormat

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

        [model-StatusDirty-Bg]=${GBT_CAR_GIT_STATUS_DIRTY_BG:-${GBT_CAR_GIT_STATUS_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-StatusDirty-Fg]=${GBT_CAR_GIT_STATUS_DIRTY_FG:-${GBT_CAR_GIT_STATUS_FG:-${GBT_CAR_GIT_FG:-red}}}
        [model-StatusDirty-Fm]=${GBT_CAR_GIT_STATUS_DIRTY_FM:-${GBT_CAR_GIT_STATUS_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-StatusDirty-Text]=$defaultStatusDirtyText

        [model-StatusClean-Bg]=${GBT_CAR_GIT_STATUS_CLEAN_BG:-${GBT_CAR_GIT_STATUS_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-StatusClean-Fg]=${GBT_CAR_GIT_STATUS_CLEAN_FG:-${GBT_CAR_GIT_STATUS_FG:-${GBT_CAR_GIT_FG:-green}}}
        [model-StatusClean-Fm]=${GBT_CAR_GIT_STATUS_CLEAN_FM:-${GBT_CAR_GIT_STATUS_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-StatusClean-Text]=$defaultStatusCleanText

        [model-Added-Bg]=${GBT_CAR_GIT_ADDED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Added-Fg]=${GBT_CAR_GIT_ADDED_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Added-Fm]=${GBT_CAR_GIT_ADDED_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Added-Text]=${GBT_CAR_GIT_ADDED_FORMAT:-'{{ AddedSymbol }}'}

        [model-AddedSymbol-Bg]=${GBT_CAR_GIT_ADDED_SYMBOL_BG:-${GBT_CAR_GIT_ADDED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-AddedSymbol-Fg]=${GBT_CAR_GIT_ADDED_SYMBOL_FG:-${GBT_CAR_GIT_ADDED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-AddedSymbol-Fm]=${GBT_CAR_GIT_ADDED_SYMBOL_FM:-${GBT_CAR_GIT_ADDED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-AddedSymbol-Text]=${GBT_CAR_GIT_ADDED_SYMBOL_TEXT:-$defaultAddedSymbolText}

        [model-AddedCount-Bg]=${GBT_CAR_GIT_ADDED_COUNT_BG:-${GBT_CAR_GIT_ADDED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-AddedCount-Fg]=${GBT_CAR_GIT_ADDED_COUNT_FG:-${GBT_CAR_GIT_ADDED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-AddedCount-Fm]=${GBT_CAR_GIT_ADDED_COUNT_FM:-${GBT_CAR_GIT_ADDED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-AddedCount-Text]=${GBT_CAR_GIT_ADDED_COUNT_TEXT:-$defaultAddedCountText}

        [model-Copied-Bg]=${GBT_CAR_GIT_COPIED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Copied-Fg]=${GBT_CAR_GIT_COPIED_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Copied-Fm]=${GBT_CAR_GIT_COPIED_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Copied-Text]=${GBT_CAR_GIT_COPIED_FORMAT:-'{{ CopiedSymbol }}'}

        [model-CopiedSymbol-Bg]=${GBT_CAR_GIT_COPIED_SYMBOL_BG:-${GBT_CAR_GIT_COPIED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-CopiedSymbol-Fg]=${GBT_CAR_GIT_COPIED_SYMBOL_FG:-${GBT_CAR_GIT_COPIED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-CopiedSymbol-Fm]=${GBT_CAR_GIT_COPIED_SYMBOL_FM:-${GBT_CAR_GIT_COPIED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-CopiedSymbol-Text]=${GBT_CAR_GIT_COPIED_SYMBOL_TEXT:-$defaultCopiedSymbolText}

        [model-CopiedCount-Bg]=${GBT_CAR_GIT_COPIED_COUNT_BG:-${GBT_CAR_GIT_COPIED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-CopiedCount-Fg]=${GBT_CAR_GIT_COPIED_COUNT_FG:-${GBT_CAR_GIT_COPIED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-CopiedCount-Fm]=${GBT_CAR_GIT_COPIED_COUNT_FM:-${GBT_CAR_GIT_COPIED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-CopiedCount-Text]=${GBT_CAR_GIT_COPIED_COUNT_TEXT:-$defaultCopiedCountText}

        [model-Deleted-Bg]=${GBT_CAR_GIT_DELETED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Deleted-Fg]=${GBT_CAR_GIT_DELETED_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Deleted-Fm]=${GBT_CAR_GIT_DELETED_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Deleted-Text]=${GBT_CAR_GIT_DELETED_FORMAT:-'{{ DeletedSymbol }}'}

        [model-DeletedSymbol-Bg]=${GBT_CAR_GIT_DELETED_SYMBOL_BG:-${GBT_CAR_GIT_DELETED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-DeletedSymbol-Fg]=${GBT_CAR_GIT_DELETED_SYMBOL_FG:-${GBT_CAR_GIT_DELETED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-DeletedSymbol-Fm]=${GBT_CAR_GIT_DELETED_SYMBOL_FM:-${GBT_CAR_GIT_DELETED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-DeletedSymbol-Text]=${GBT_CAR_GIT_DELETED_SYMBOL_TEXT:-$defaultDeletedSymbolText}

        [model-DeletedCount-Bg]=${GBT_CAR_GIT_DELETED_COUNT_BG:-${GBT_CAR_GIT_DELETED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-DeletedCount-Fg]=${GBT_CAR_GIT_DELETED_COUNT_FG:-${GBT_CAR_GIT_DELETED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-DeletedCount-Fm]=${GBT_CAR_GIT_DELETED_COUNT_FM:-${GBT_CAR_GIT_DELETED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-DeletedCount-Text]=${GBT_CAR_GIT_DELETED_COUNT_TEXT:-$defaultDeletedCountText}

        [model-Ignored-Bg]=${GBT_CAR_GIT_IGNORED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Ignored-Fg]=${GBT_CAR_GIT_IGNORED_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Ignored-Fm]=${GBT_CAR_GIT_IGNORED_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Ignored-Text]=${GBT_CAR_GIT_IGNORED_FORMAT:-'{{ IgnoredSymbol }}'}

        [model-IgnoredSymbol-Bg]=${GBT_CAR_GIT_IGNORED_SYMBOL_BG:-${GBT_CAR_GIT_IGNORED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-IgnoredSymbol-Fg]=${GBT_CAR_GIT_IGNORED_SYMBOL_FG:-${GBT_CAR_GIT_IGNORED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-IgnoredSymbol-Fm]=${GBT_CAR_GIT_IGNORED_SYMBOL_FM:-${GBT_CAR_GIT_IGNORED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-IgnoredSymbol-Text]=${GBT_CAR_GIT_IGNORED_SYMBOL_TEXT:-$defaultIgnoredSymbolText}

        [model-IgnoredCount-Bg]=${GBT_CAR_GIT_IGNORED_COUNT_BG:-${GBT_CAR_GIT_IGNORED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-IgnoredCount-Fg]=${GBT_CAR_GIT_IGNORED_COUNT_FG:-${GBT_CAR_GIT_IGNORED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-IgnoredCount-Fm]=${GBT_CAR_GIT_IGNORED_COUNT_FM:-${GBT_CAR_GIT_IGNORED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-IgnoredCount-Text]=${GBT_CAR_GIT_IGNORED_COUNT_TEXT:-$defaultIgnoredCountText}

        [model-Modified-Bg]=${GBT_CAR_GIT_MODIFIED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Modified-Fg]=${GBT_CAR_GIT_MODIFIED_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Modified-Fm]=${GBT_CAR_GIT_MODIFIED_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Modified-Text]=${GBT_CAR_GIT_MODIFIED_FORMAT:-'{{ ModifiedSymbol }}'}

        [model-ModifiedSymbol-Bg]=${GBT_CAR_GIT_MODIFIED_SYMBOL_BG:-${GBT_CAR_GIT_MODIFIED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-ModifiedSymbol-Fg]=${GBT_CAR_GIT_MODIFIED_SYMBOL_FG:-${GBT_CAR_GIT_MODIFIED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-ModifiedSymbol-Fm]=${GBT_CAR_GIT_MODIFIED_SYMBOL_FM:-${GBT_CAR_GIT_MODIFIED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-ModifiedSymbol-Text]=${GBT_CAR_GIT_MODIFIED_SYMBOL_TEXT:-$defaultModifiedSymbolText}

        [model-ModifiedCount-Bg]=${GBT_CAR_GIT_MODIFIED_COUNT_BG:-${GBT_CAR_GIT_MODIFIED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-ModifiedCount-Fg]=${GBT_CAR_GIT_MODIFIED_COUNT_FG:-${GBT_CAR_GIT_MODIFIED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-ModifiedCount-Fm]=${GBT_CAR_GIT_MODIFIED_COUNT_FM:-${GBT_CAR_GIT_MODIFIED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-ModifiedCount-Text]=${GBT_CAR_GIT_MODIFIED_COUNT_TEXT:-$defaultModifiedCountText}

        [model-Renamed-Bg]=${GBT_CAR_GIT_RENAMED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Renamed-Fg]=${GBT_CAR_GIT_RENAMED_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Renamed-Fm]=${GBT_CAR_GIT_RENAMED_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Renamed-Text]=${GBT_CAR_GIT_RENAMED_FORMAT:-'{{ RenamedSymbol }}'}

        [model-RenamedSymbol-Bg]=${GBT_CAR_GIT_RENAMED_SYMBOL_BG:-${GBT_CAR_GIT_RENAMED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-RenamedSymbol-Fg]=${GBT_CAR_GIT_RENAMED_SYMBOL_FG:-${GBT_CAR_GIT_RENAMED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-RenamedSymbol-Fm]=${GBT_CAR_GIT_RENAMED_SYMBOL_FM:-${GBT_CAR_GIT_RENAMED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-RenamedSymbol-Text]=${GBT_CAR_GIT_RENAMED_SYMBOL_TEXT:-$defaultRenamedSymbolText}

        [model-RenamedCount-Bg]=${GBT_CAR_GIT_RENAMED_COUNT_BG:-${GBT_CAR_GIT_RENAMED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-RenamedCount-Fg]=${GBT_CAR_GIT_RENAMED_COUNT_FG:-${GBT_CAR_GIT_RENAMED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-RenamedCount-Fm]=${GBT_CAR_GIT_RENAMED_COUNT_FM:-${GBT_CAR_GIT_RENAMED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-RenamedCount-Text]=${GBT_CAR_GIT_RENAMED_COUNT_TEXT:-$defaultRenamedCountText}

        [model-Staged-Bg]=${GBT_CAR_GIT_STAGED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Staged-Fg]=${GBT_CAR_GIT_STAGED_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Staged-Fm]=${GBT_CAR_GIT_STAGED_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Staged-Text]=${GBT_CAR_GIT_STAGED_FORMAT:-'{{ StagedSymbol }}'}

        [model-StagedSymbol-Bg]=${GBT_CAR_GIT_STAGED_SYMBOL_BG:-${GBT_CAR_GIT_STAGED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-StagedSymbol-Fg]=${GBT_CAR_GIT_STAGED_SYMBOL_FG:-${GBT_CAR_GIT_STAGED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-StagedSymbol-Fm]=${GBT_CAR_GIT_STAGED_SYMBOL_FM:-${GBT_CAR_GIT_STAGED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-StagedSymbol-Text]=${GBT_CAR_GIT_STAGED_SYMBOL_TEXT:-$defaultStagedSymbolText}

        [model-StagedCount-Bg]=${GBT_CAR_GIT_STAGED_COUNT_BG:-${GBT_CAR_GIT_STAGED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-StagedCount-Fg]=${GBT_CAR_GIT_STAGED_COUNT_FG:-${GBT_CAR_GIT_STAGED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-StagedCount-Fm]=${GBT_CAR_GIT_STAGED_COUNT_FM:-${GBT_CAR_GIT_STAGED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-StagedCount-Text]=${GBT_CAR_GIT_STAGED_COUNT_TEXT:-$defaultStagedCountText}

        [model-Unmerged-Bg]=${GBT_CAR_GIT_UNMERGED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Unmerged-Fg]=${GBT_CAR_GIT_UNMERGED_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Unmerged-Fm]=${GBT_CAR_GIT_UNMERGED_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Unmerged-Text]=${GBT_CAR_GIT_UNMERGED_FORMAT:-'{{ UnmergedSymbol }}'}

        [model-UnmergedSymbol-Bg]=${GBT_CAR_GIT_UNMERGED_SYMBOL_BG:-${GBT_CAR_GIT_UNMERGED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-UnmergedSymbol-Fg]=${GBT_CAR_GIT_UNMERGED_SYMBOL_FG:-${GBT_CAR_GIT_UNMERGED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-UnmergedSymbol-Fm]=${GBT_CAR_GIT_UNMERGED_SYMBOL_FM:-${GBT_CAR_GIT_UNMERGED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-UnmergedSymbol-Text]=${GBT_CAR_GIT_UNMERGED_SYMBOL_TEXT:-$defaultUnmergedSymbolText}

        [model-UnmergedCount-Bg]=${GBT_CAR_GIT_UNMERGED_COUNT_BG:-${GBT_CAR_GIT_UNMERGED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-UnmergedCount-Fg]=${GBT_CAR_GIT_UNMERGED_COUNT_FG:-${GBT_CAR_GIT_UNMERGED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-UnmergedCount-Fm]=${GBT_CAR_GIT_UNMERGED_COUNT_FM:-${GBT_CAR_GIT_UNMERGED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-UnmergedCount-Text]=${GBT_CAR_GIT_UNMERGED_COUNT_TEXT:-$defaultUnmergedCountText}

        [model-Untracked-Bg]=${GBT_CAR_GIT_UNTRACKED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Untracked-Fg]=${GBT_CAR_GIT_UNTRACKED_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Untracked-Fm]=${GBT_CAR_GIT_UNTRACKED_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Untracked-Text]=${GBT_CAR_GIT_UNTRACKED_FORMAT:-'{{ UntrackedSymbol }}'}

        [model-UntrackedSymbol-Bg]=${GBT_CAR_GIT_UNTRACKED_SYMBOL_BG:-${GBT_CAR_GIT_UNTRACKED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-UntrackedSymbol-Fg]=${GBT_CAR_GIT_UNTRACKED_SYMBOL_FG:-${GBT_CAR_GIT_UNTRACKED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-UntrackedSymbol-Fm]=${GBT_CAR_GIT_UNTRACKED_SYMBOL_FM:-${GBT_CAR_GIT_UNTRACKED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-UntrackedSymbol-Text]=${GBT_CAR_GIT_UNTRACKED_SYMBOL_TEXT:-$defaultUntrackedSymbolText}

        [model-UntrackedCount-Bg]=${GBT_CAR_GIT_UNTRACKED_COUNT_BG:-${GBT_CAR_GIT_UNTRACKED_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-UntrackedCount-Fg]=${GBT_CAR_GIT_UNTRACKED_COUNT_FG:-${GBT_CAR_GIT_UNTRACKED_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-UntrackedCount-Fm]=${GBT_CAR_GIT_UNTRACKED_COUNT_FM:-${GBT_CAR_GIT_UNTRACKED_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-UntrackedCount-Text]=${GBT_CAR_GIT_UNTRACKED_COUNT_TEXT:-$defaultUntrackedCountText}

        [model-Ahead-Bg]=${GBT_CAR_GIT_AHEAD_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Ahead-Fg]=${GBT_CAR_GIT_AHEAD_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Ahead-Fm]=${GBT_CAR_GIT_AHEAD_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Ahead-Text]=${GBT_CAR_GIT_AHEAD_FORMAT:-'{{ AheadSymbol }}'}

        [model-AheadSymbol-Bg]=${GBT_CAR_GIT_AHEAD_SYMBOL_BG:-${GBT_CAR_GIT_AHEAD_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-AheadSymbol-Fg]=${GBT_CAR_GIT_AHEAD_SYMBOL_FG:-${GBT_CAR_GIT_AHEAD_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-AheadSymbol-Fm]=${GBT_CAR_GIT_AHEAD_SYMBOL_FM:-${GBT_CAR_GIT_AHEAD_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-AheadSymbol-Text]=${GBT_CAR_GIT_AHEAD_SYMBOL_TEXT:-$defaultAheadSymbolText}

        [model-AheadCount-Bg]=${GBT_CAR_GIT_AHEAD_COUNT_BG:-${GBT_CAR_GIT_AHEAD_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-AheadCount-Fg]=${GBT_CAR_GIT_AHEAD_COUNT_FG:-${GBT_CAR_GIT_AHEAD_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-AheadCount-Fm]=${GBT_CAR_GIT_AHEAD_COUNT_FM:-${GBT_CAR_GIT_AHEAD_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-AheadCount-Text]=${GBT_CAR_GIT_AHEAD_COUNT_TEXT:-$defaultAheadCountText}

        [model-Behind-Bg]=${GBT_CAR_GIT_BEHIND_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Behind-Fg]=${GBT_CAR_GIT_BEHIND_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Behind-Fm]=${GBT_CAR_GIT_BEHIND_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Behind-Text]=${GBT_CAR_GIT_BEHIND_TEXT:-'{{ BehindSymbol }}'}

        [model-BehindSymbol-Bg]=${GBT_CAR_GIT_BEHIND_SYMBOL_BG:-${GBT_CAR_GIT_BEHIND_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-BehindSymbol-Fg]=${GBT_CAR_GIT_BEHIND_SYMBOL_FG:-${GBT_CAR_GIT_BEHIND_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-BehindSymbol-Fm]=${GBT_CAR_GIT_BEHIND_SYMBOL_FM:-${GBT_CAR_GIT_BEHIND_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-BehindSymbol-Text]=${GBT_CAR_GIT_BEHIND_SYMBOL_TEXT:-$defaultBehindSymbolText}

        [model-BehindCount-Bg]=${GBT_CAR_GIT_BEHIND_COUNT_BG:-${GBT_CAR_GIT_BEHIND_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-BehindCount-Fg]=${GBT_CAR_GIT_BEHIND_COUNT_FG:-${GBT_CAR_GIT_BEHIND_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-BehindCount-Fm]=${GBT_CAR_GIT_BEHIND_COUNT_FM:-${GBT_CAR_GIT_BEHIND_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-BehindCount-Text]=${GBT_CAR_GIT_BEHIND_COUNT_TEXT:-$defaultBehindCountText}

        [model-Stash-Bg]=${GBT_CAR_GIT_STASH_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}
        [model-Stash-Fg]=${GBT_CAR_GIT_STASH_FG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}
        [model-Stash-Fm]=${GBT_CAR_GIT_STASH_FM:-${GBT_CAR_GIT_FM:-$defaultRootFm}}
        [model-Stash-Text]=${GBT_CAR_GIT_STASH_TEXT:-'{{ StashSymbol }}'}

        [model-StashSymbol-Bg]=${GBT_CAR_GIT_STASH_SYMBOL_BG:-${GBT_CAR_GIT_STASH_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-StashSymbol-Fg]=${GBT_CAR_GIT_STASH_SYMBOL_FG:-${GBT_CAR_GIT_STASH_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-StashSymbol-Fm]=${GBT_CAR_GIT_STASH_SYMBOL_FM:-${GBT_CAR_GIT_STASH_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-StashSymbol-Text]=${GBT_CAR_GIT_STASH_SYMBOL_TEXT:-$defaultStashSymbolText}

        [model-StashCount-Bg]=${GBT_CAR_GIT_STASH_COUNT_BG:-${GBT_CAR_GIT_STASH_BG:-${GBT_CAR_GIT_BG:-$defaultRootBg}}}
        [model-StashCount-Fg]=${GBT_CAR_GIT_STASH_COUNT_FG:-${GBT_CAR_GIT_STASH_BG:-${GBT_CAR_GIT_FG:-$defaultRootFg}}}
        [model-StashCount-Fm]=${GBT_CAR_GIT_STASH_COUNT_FM:-${GBT_CAR_GIT_STASH_BG:-${GBT_CAR_GIT_FM:-$defaultRootFm}}}
        [model-StashCount-Text]=${GBT_CAR_GIT_STASH_COUNT_TEXT:-$defaultStashCountText}

        [display]=${GBT_CAR_GIT_DISPLAY:-$isGitDir}
        [wrap]=${GBT_CAR_GIT_WRAP:-0}
        [sep]=${GBT_CAR_GIT_SEP:-'\x00'}
    )
}
