function GbtCarExecTime() {
    local defaultRootBg=${GBT_CAR_BG:-light_gray}
    local defaultRootFg=${GBT_CAR_FG:-black}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local precision=${GBT_CAR_EXECTIME_PRECISION:-0}
    local now=$(${GBT__SOURCE_DATE:-date} ${GBT__SOURCE_DATE_ARG:-'+%s.%N'})
    local execs=$(echo "$precision $now ${GBT_CAR_EXECTIME_SECS:-$now}" | awk '{printf "%0.*f", $1, $2 - $3}')

    if (( $precision > 0 )); then
        local subsecs=".${execs#*.}"
    fi

    execs=${execs%%.*}

    local hours=$(( execs / 3600 ))
    local mins=$(( (execs - hours * 3600) / 60 ))
    local secs=$(( execs - hours * 3600 - mins * 60 ))
    local exectime=$(echo "${hours%%.*} ${mins%%.*} ${secs%%.*} ${subsecs}" | awk '{printf "%.2d:%.2d:%02d%s", $1, $2, $3, $4 }' )

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_EXECTIME_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_EXECTIME_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_EXECTIME_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_EXECTIME_FORMAT:-' {{ Time }} '}
        [model-Time-Bg]=${GBT_CAR_EXECTIME_TIME_BG:-${GBT_CAR_EXECTIME_BG:-$defaultRootBg}}
        [model-Time-Fg]=${GBT_CAR_EXECTIME_TIME_FG:-${GBT_CAR_EXECTIME_FG:-$defaultRootFg}}
        [model-Time-Fm]=${GBT_CAR_EXECTIME_TIME_FM:-${GBT_CAR_EXECTIME_FM:-$defaultRootFm}}
        [model-Time-Text]=${GBT_CAR_EXECTIME_TIME_TEXT:-$exectime}

        [display]=${GBT_CAR_EXECTIME_DISPLAY:-1}
        [wrap]=${GBT_CAR_EXECTIME_WRAP:-0}
        [sep]=${GBT_CAR_EXECTIME_SEP:-'\x00'}
    )
}
