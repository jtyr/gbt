function GbtCarExecTime() {
    local defaultRootBg=${GBT_CAR_BG:-light_gray}
    local defaultRootFg=${GBT_CAR_FG:-black}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local defaultSep="\x00"

    local precision=${GBT_CAR_EXECTIME_PRECISION:-0}
    local now=$(${GBT__SOURCE_DATE:-date} ${GBT__SOURCE_DATE_ARG:-'+%s.%N'})
    local execs=$(echo "$now ${GBT_CAR_EXECTIME_SECS:-$now}" | awk '{printf "%9f", $1 - $2}')
    local subsecs="0.${execs#*.}"

    local hours=$(echo "$execs" | awk '{printf "%.0f", $1 / 3600}')
    local mins=$(echo "$execs $hours" | awk '{printf "%.0f", ($1 - $2 * 3600) / 60}')
    local secs=$(echo "$execs $hours $mins" | awk '{printf "%.0f", $1 - $2 * 3600 - $3 * 60}')

    local durationtime=''
    local secondstime=''
    local exectime=''

    if [[ ${GBT_CAR_EXECTIME_FORMAT-' {{ Time }} '} == *' {{ Duration }} '* ]]; then
        # Duration
        local millis=0
        local micros=0
        local nanos=0

        if [[ $precision > 0 ]]; then
            subsecs=$(echo "$subsecs" | awk '{printf "%0.9f", $1 * 1000}')
            millis="${subsecs%.*}"

            if [[ $precision > 3 ]] || ( [[ $secs == 0 ]] && [[ $millis == 0 ]] ); then
                subsecs=$(echo "$subsecs ${subsecs%.*}" | awk '{printf "%.9f", ($1 - $2) * 1000}')
                micros="${subsecs%.*}"

                if [[ $precision > 6 ]] || ( [[ $secs == 0 ]] && [[ $millis == 0 ]] && [[ $micros == 0 ]] ); then
                    subsecs=$(echo "$subsecs ${subsecs%.*}" | awk '{printf "%.9f", ($1 - $2) * 1000}')
                    nanos="${subsecs%.*}"
                fi
            fi
        fi

        if [[ $hours > 0 ]]; then
            durationtime+="${hours}h"
        fi

        if [[ $mins > 0 ]]; then
            durationtime+="${mins}m"
        fi

        if [[ $secs > 0 ]] || [[ $precision == 0 ]]; then
            durationtime+="${secs}s"
        fi

        if [[ $millis > 0 ]]; then
            durationtime+="${millis}ms"
        fi

        if [[ $micros > 0 ]]; then
            durationtime+="${micros}µs"
        fi

        if [[ $nanos > 0 ]]; then
            durationtime+="${nanos}ns"
        fi
    elif [[ ${GBT_CAR_EXECTIME_FORMAT-' {{ Time }} '} == *' {{ Seconds }} '* ]]; then
        # Seconds
        local secondstime=$(echo "$precision $execs" | awk '{printf "%0.*f", $1, $2}')
    elif [[ ${GBT_CAR_EXECTIME_FORMAT-' {{ Time }} '} == *' {{ Time }} '* ]]; then
        # Time
        if [[ $precision == 0 ]]; then
            subsecs=''
        fi

        local exectime=$(echo "${hours%%.*} ${mins%%.*} ${secs%%.*} ${subsecs:1:$precision+1}" | awk '{printf "%.2d:%.2d:%02d%s", $1, $2, $3, $4 }')
    fi

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_EXECTIME_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_EXECTIME_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_EXECTIME_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_EXECTIME_FORMAT-' {{ Time }} '}

        [model-Duration-Bg]=${GBT_CAR_EXECTIME_DURATION_BG:-${GBT_CAR_EXECTIME_BG:-$defaultRootBg}}
        [model-Duration-Fg]=${GBT_CAR_EXECTIME_DURATION_FG:-${GBT_CAR_EXECTIME_FG:-$defaultRootFg}}
        [model-Duration-Fm]=${GBT_CAR_EXECTIME_DURATION_FM:-${GBT_CAR_EXECTIME_FM:-$defaultRootFm}}
        [model-Duration-Text]=${GBT_CAR_EXECTIME_DURATION_TEXT-$durationtime}

        [model-Seconds-Bg]=${GBT_CAR_EXECTIME_SECONDS_BG:-${GBT_CAR_EXECTIME_BG:-$defaultRootBg}}
        [model-Seconds-Fg]=${GBT_CAR_EXECTIME_SECONDS_FG:-${GBT_CAR_EXECTIME_FG:-$defaultRootFg}}
        [model-Seconds-Fm]=${GBT_CAR_EXECTIME_SECONDS_FM:-${GBT_CAR_EXECTIME_FM:-$defaultRootFm}}
        [model-Seconds-Text]=${GBT_CAR_EXECTIME_SECONDS_TEXT-$secondstime}

        [model-Time-Bg]=${GBT_CAR_EXECTIME_TIME_BG:-${GBT_CAR_EXECTIME_BG:-$defaultRootBg}}
        [model-Time-Fg]=${GBT_CAR_EXECTIME_TIME_FG:-${GBT_CAR_EXECTIME_FG:-$defaultRootFg}}
        [model-Time-Fm]=${GBT_CAR_EXECTIME_TIME_FM:-${GBT_CAR_EXECTIME_FM:-$defaultRootFm}}
        [model-Time-Text]=${GBT_CAR_EXECTIME_TIME_TEXT-$exectime}

        [model-Sep-Bg]=${GBT_CAR_EXECTIME_SEP_BG:-$defaultSep}
        [model-Sep-Fg]=${GBT_CAR_EXECTIME_SEP_FG:-$defaultSep}
        [model-Sep-Fm]=${GBT_CAR_EXECTIME_SEP_FM:-$defaultSep}
        [model-Sep-Text]=${GBT_CAR_EXECTIME_SEP_TEXT:-${GBT_CAR_EXECTIME_SEP:-$defaultSep}}

        [display]=${GBT_CAR_EXECTIME_DISPLAY:-1}
        [wrap]=${GBT_CAR_EXECTIME_WRAP:-0}
    )
}
