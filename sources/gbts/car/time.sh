function GbtCarTimeFormat() {
    local s=$1

    # Matching of `date` and `go` time formatting strings
    # (https://golang.org/src/time/format.go#L87)
    s=${s//January/%B}
    s=${s//Jan/%b}
    s=${s//01/%m}
    s=${s//Monday/%A}
    s=${s//Mon/%a}
    s=${s//02/%d}
    s=${s//15/%H}
    s=${s//03/%I}
    s=${s//04/%M}
    s=${s//05/%S}
    s=${s//2006/%Y}
    s=${s//06/%y}
    s=${s//PM/%p}
    s=${s//MST/%Z}

    GBT__RETVAL=$s
}


function GbtCarTime() {
    local defaultRootBg=${GBT_CAR_BG:-light_blue}
    local defaultRootFg=${GBT_CAR_FG:-light_gray}
    local defaultRootFm=${GBT_CAR_FM:-none}

    GbtCarTimeFormat "${GBT_CAR_TIME_DATE_FORMAT-Mon 02 Jan}"
    local defaultDateText=$(${GBT__SOURCE_DATE:-date} "+$GBT__RETVAL")
    GbtCarTimeFormat "${GBT_CAR_TIME_TIME_FORMAT-15:04:05}"
    local defaultTimeText=$(${GBT__SOURCE_DATE:-date} "+$GBT__RETVAL")

    local defaultSep="\x00"

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_TIME_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_TIME_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_TIME_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_TIME_FORMAT-' {{ DateTime }} '}

        [model-DateTime-Bg]=${GBT_CAR_TIME_DATETIME_BG:-${GBT_CAR_TIME_BG:-$defaultRootBg}}
        [model-DateTime-Fg]=${GBT_CAR_TIME_DATETIME_FG:-${GBT_CAR_TIME_FG:-$defaultRootFg}}
        [model-DateTime-Fm]=${GBT_CAR_TIME_DATETIME_FM:-${GBT_CAR_TIME_FM:-$defaultRootFm}}
        [model-DateTime-Text]=${GBT_CAR_TIME_DATETIME_FORMAT-'{{ Date }} {{ Time }}'}

        [model-Date-Bg]=${GBT_CAR_TIME_DATE_BG:-${GBT_CAR_TIME_DATETIME_FG:-${GBT_CAR_TIME_BG:-$defaultRootBg}}}
        [model-Date-Fg]=${GBT_CAR_TIME_DATE_FG:-${GBT_CAR_TIME_DATETIME_BG:-${GBT_CAR_TIME_FG:-$defaultRootFg}}}
        [model-Date-Fm]=${GBT_CAR_TIME_DATE_FM:-${GBT_CAR_TIME_DATETIME_FM:-${GBT_CAR_TIME_FM:-$defaultRootFm}}}
        [model-Date-Text]=$defaultDateText

        [model-Time-Bg]=${GBT_CAR_TIME_TIME_BG:-${GBT_CAR_TIME_DATETIME_FG:-${GBT_CAR_TIME_BG:-$defaultRootBg}}}
        [model-Time-Fg]=${GBT_CAR_TIME_TIME_FG:-${GBT_CAR_TIME_DATETIME_BG:-${GBT_CAR_TIME_FG:-light_yellow}}}
        [model-Time-Fm]=${GBT_CAR_TIME_TIME_FM:-${GBT_CAR_TIME_DATETIME_FM:-${GBT_CAR_TIME_FM:-$defaultRootFm}}}
        [model-Time-Text]=$defaultTimeText

        [model-Sep-Bg]=${GBT_CAR_TIME_SEP_BG:-$defaultSep}
        [model-Sep-Fg]=${GBT_CAR_TIME_SEP_FG:-$defaultSep}
        [model-Sep-Fm]=${GBT_CAR_TIME_SEP_FM:-$defaultSep}
        [model-Sep-Text]=${GBT_CAR_TIME_SEP_TEXT:-${GBT_CAR_TIME_SEP:-$defaultSep}}

        [display]=${GBT_CAR_TIME_DISPLAY:-1}
        [wrap]=${GBT_CAR_TIME_WRAP:-0}
    )
}
