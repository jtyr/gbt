declare -A GBT__STATUS_SIGNALS
GBT__STATUS_SIGNALS=(
    # Usual exit codes
    [-1]='FATAL'
    [0]='OK'
    [1]='FAIL'
    [2]='BLTINMUSE'
    [6]='UNKADDR'

    # Issue with the actual command being invoked
    [126]='NOEXEC'
    [127]='NOTFOUND'

    # Signal errors (128 + signal)
    [129]='SIGHUP'
    [130]='SIGINT'
    [131]='SIGQUIT'
    [132]='SIGILL'
    [133]='SIGTRAP'
    [134]='SIGABRT'
    [135]='SIGBUS'
    [136]='SIGFPE'
    [137]='SIGKILL'
    [138]='SIGUSR1'
    [139]='SIGSEGV'
    [140]='SIGUSR2'
    [141]='SIGPIPE'
    [142]='SIGALRM'
    [143]='SIGTERM'
    [145]='SIGCHLD'
    [146]='SIGCONT'
    [147]='SIGSTOP'
    [148]='SIGTSTP'
    [149]='SIGTTIN'
    [150]='SIGTTOU'
)

function GbtCarStatus() {
    if [[ ${GBT_CAR_STATUS_DISPLAY:-0} == 0 ]] && [[ $1 == 0 ]]; then
        return
    fi

    local defaultErrorBg=${GBT_CAR_BG:-red}
    local defaultErrorFg=${GBT_CAR_BG:-light_gray}
    local defaultErrorFm=${GBT_CAR_BG:-none}
    local defaultOkBg=${GBT_CAR_BG:-green}
    local defaultOkFg=${GBT_CAR_BG:-light_gray}
    local defaultOkFm=${GBT_CAR_BG:-none}

    local defaultRootBg=$defaultErrorBg
    local defaultRootFg=$defaultErrorFg
    local defaultRootFm=$defaultErrorFm

    GbtDecorateUnicode ${GBT_CAR_STATUS_ERROR_TEXT-'\xe2\x9c\x98'}
    local defaultErrorText=$GBT__RETVAL
    GbtDecorateUnicode ${GBT_CAR_STATUS_OK_TEXT-'\xe2\x9c\x94'}
    local defaultOkText=$GBT__RETVAL

    local defaultDetailsFormat=' {{ Signal }}'
    local defaultSymbolFormat='{{ Error }}'
    local defaultCodeText='?'

    if (( ${#@} > 0 )); then
        defaultCodeText=$1
    fi

    if [[ $defaultCodeText == 0 ]]; then
        defaultRootBg=$defaultOkBg
        defaultRootFg=$defaultOkFg
        defaultRootFm=$defaultOkFm
        defaultDetailsFormat=''
        defaultSymbolFormat='{{ Ok }}'
    else
        defaultDetailsFormat=${GBT_CAR_STATUS_DETAILS_FORMAT-$defaultDetailsFormat}
    fi

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_STATUS_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_STATUS_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_STATUS_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_STATUS_FORMAT-' {{ Symbol }} '}

        [model-Error-Bg]=${GBT_CAR_STATUS_ERROR_BG:-${GBT_CAR_STATUS_SYMBOL_BG:-${GBT_CAR_STATUS_BG:-$defaultErrorBg}}}
        [model-Error-Fg]=${GBT_CAR_STATUS_ERROR_FG:-${GBT_CAR_STATUS_SYMBOL_FG:-${GBT_CAR_STATUS_FG:-$defaultErrorFg}}}
        [model-Error-Fm]=${GBT_CAR_STATUS_ERROR_FM:-${GBT_CAR_STATUS_SYMBOL_FM:-${GBT_CAR_STATUS_FM:-$defaultErrorFm}}}
        [model-Error-Text]=$defaultErrorText

        [model-Ok-Bg]=${GBT_CAR_STATUS_OK_BG:-${GBT_CAR_STATUS_SYMBOL_BG:-${GBT_CAR_STATUS_BG:-$defaultOkBg}}}
        [model-Ok-Fg]=${GBT_CAR_STATUS_OK_FG:-${GBT_CAR_STATUS_SYMBOL_FG:-${GBT_CAR_STATUS_FG:-$defaultOkFg}}}
        [model-Ok-Fm]=${GBT_CAR_STATUS_OK_FM:-${GBT_CAR_STATUS_SYMBOL_FM:-${GBT_CAR_STATUS_FM:-$defaultOkFm}}}
        [model-Ok-Text]=$defaultOkText

        [model-Symbol-Bg]=${GBT_CAR_STATUS_SYMBOL_BG:-${GBT_CAR_STATUS_BG:-$defaultRootBg}}
        [model-Symbol-Fg]=${GBT_CAR_STATUS_SYMBOL_FG-${GBT_CAR_STATUS_FG:-$defaultRootFg}}
        [model-Symbol-Fm]=${GBT_CAR_STATUS_SYMBOL_FM:-${GBT_CAR_STATUS_FM:-$defaultRootFm}}
        [model-Symbol-Text]=${GBT_CAR_STATUS_SYMBOL_FORMAT-$defaultSymbolFormat}

        [model-Details-Bg]=${GBT_CAR_STATUS_DETAILS_BG:-${GBT_CAR_STATUS_BG:-$defaultDetailsBg}}
        [model-Details-Fg]=${GBT_CAR_STATUS_DETAILS_FG:-${GBT_CAR_STATUS_FG:-$defaultDetailsFg}}
        [model-Details-Fm]=${GBT_CAR_STATUS_DETAILS_FM:-${GBT_CAR_STATUS_FM:-$defaultDetailsFm}}
        [model-Details-Text]=$defaultDetailsFormat

        [model-Code-Bg]=${GBT_CAR_STATUS_CODE_BG:-${GBT_CAR_STATUS_BG:-$defaultRootBg}}
        [model-Code-Fg]=${GBT_CAR_STATUS_CODE_FG:-${GBT_CAR_STATUS_FG:-$defaultRootFg}}
        [model-Code-Fm]=${GBT_CAR_STATUS_CODE_FM:-${GBT_CAR_STATUS_FM:-$defaultRootFm}}
        [model-Code-Text]=${GBT_CAR_STATUS_CODE_TEXT-$defaultCodeText}

        [model-Signal-Bg]=${GBT_CAR_STATUS_SIGNAL_BG:-${GBT_CAR_STATUS_BG:-$defaultRootBg}}
        [model-Signal-Fg]=${GBT_CAR_STATUS_SIGNAL_FG:-${GBT_CAR_STATUS_FG:-$defaultRootFg}}
        [model-Signal-Fm]=${GBT_CAR_STATUS_SIGNAL_FM:-${GBT_CAR_STATUS_FM:-$defaultRootFm}}
        [model-Signal-Text]=${GBT_CAR_STATUS_SIGNAL_TEXT-${GBT__STATUS_SIGNALS[$defaultCodeText]:-UNK}}

        [wrap]=${GBT_CAR_STATUS_WRAP:-0}
        [sep]=${GBT_CAR_STATUS_SEP-'\x00'}
    )

    if [[ $1 == 0 ]]; then
        GBT_CAR[display]=${GBT_CAR_STATUS_DISPLAY:-0}
    else
        GBT_CAR[display]=${GBT_CAR_STATUS_DISPLAY:-1}
    fi
}
