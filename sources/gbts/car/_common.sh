# shellcheck shell=bash
declare -A GBT_COLORS
GBT_COLORS=(
    [black]=0
    [red]=1
    [green]=2
    [yellow]=3
    [blue]=4
    [magenta]=5
    [cyan]=6
    [light_gray]=7
    [dark_gray]=8
    [light_red]=9
    [light_green]=10
    [light_yellow]=11
    [light_blue]=12
    [light_magenta]=13
    [light_cyan]=14
    [white]=15
)

declare -A GBT_HIGHER_COLORS
GBT_HIGHER_COLORS=(
    [black]=16
    [red]=124
    [green]=34
    [yellow]=100
    [blue]=19
    [magenta]=90
    [cyan]=30
    [light_gray]=248
    [dark_gray]=240
    [light_red]=196
    [light_green]=46
    [light_yellow]=226
    [light_blue]=63
    [light_magenta]=201
    [light_cyan]=51
    [white]=231
)

function GbtGetColor() {
    local name=$1
    local isFg=$2

    local kind=4
    local seq=''
    local esc='\x1b'

    if [[ $isFg == 1 ]]; then
        kind=3
    fi

    if [[ $GBT_SHELL == '_bash' ]]; then
        esc='\e'
    fi

    if [[ $name == 'RESETALL' ]]; then
        seq="${esc}[0m"
    elif [[ $name == 'default' ]]; then
        # Default
        seq="${esc}[${kind}9m"
    else
        if [ ${GBT_COLORS[$name]+1} ]; then
            local val=${GBT_COLORS[$name]}

            if [[ ${GBT_FORCE_HIGHER_COLORS:-1} == 1 ]]; then
                val=${GBT_HIGHER_COLORS[$name]}
            fi

            # Named color
            seq="${esc}[${kind}8;5;${val}m"
        elif [[ $name =~ ^[0-9]{1,3}$ ]]; then
            local val=$name

            if [[ ${GBT_FORCE_HIGHER_COLORS:-1} == 1 ]]; then
                for k in "${!GBT_COLORS[@]}"; do
                    if [[ ${GBT_COLORS[$k]} == "$name" ]]; then
                        val=${GBT_HIGHER_COLORS[$k]}
                    fi
                done
            fi

            # Color number
            seq="${esc}[${kind}8;5;${val}m"
        elif [[ $name =~ ^[0-9]{1,3}\;[0-9]{1,3}\;[0-9]{1,3}$ ]]; then
            # RGB color
            seq="${esc}[${kind}8;2;${name}m"
        else
            # If anything else, use default
            seq="${esc}[${kind}9m"
        fi
    fi

    GbtDecorateShell "$seq"
}


function GbtGetFormat() {
    local name=$1
    local end=$2

    local seq=''
    local kind=''
    local esc='\x1b'

    if [[ $end == 1 ]]; then
        kind=2
    fi

    if [[ $GBT_SHELL == '_bash' ]]; then
        esc='\e'
    fi

    if [[ $name != "${name//normal/}" ]]; then
        seq+="${esc}[0m"
    fi

    if [[ $name != "${name//bold/}" ]]; then
        if [[ $end == 1 ]]; then
            seq+="${esc}[22m"
        else
            seq+="${esc}[${kind}1m"
        fi
    fi

    if [[ $name != "${name//dim/}" ]]; then
        seq+="${esc}[${kind}2m"
    fi

    if [[ $name != "${name//underline/}" ]]; then
        seq+="${esc}[${kind}4m"
    fi

    if [[ $name != "${name//blink/}" ]]; then
        seq+="${esc}[${kind}5m"
    fi

    if [[ $name != "${name//invert/}" ]]; then
        seq+="${esc}[${kind}7m"
    fi

    if [[ $name != "${name//hide/}" ]]; then
        seq+="${esc}[${kind}8m"
    fi

    if [[ $name != "${name//strikeout/}" ]]; then
        seq+="${esc}[${kind}9m"
    fi

    GbtDecorateShell "$seq"
}


function GbtDecorateUnicode() {
    local unicode=$1

    # Shell decorate all characters but the last four
    if [[ ${unicode} =~ ^(\\x[0-9a-f]{2}){5}$ ]]; then
        GbtDecorateShell "${unicode:0:${#unicode}-4}"
        GBT__RETVAL="$GBT__RETVAL${unicode:16}"
    elif [[ ${unicode} =~ ^(\\x[0-9a-f]{2}){3}$ ]]; then
        GbtDecorateShell "${unicode:0:${#unicode}-4}"
        GBT__RETVAL="$GBT__RETVAL${unicode:8}"
    else
        GBT__RETVAL=$unicode
    fi
}


function GbtDecorateShell() {
    local seq=$1

    if [ -z "$seq" ]; then
        GBT__RETVAL=''
    elif [[ $GBT_SHELL == 'zsh' ]]; then
        GBT__RETVAL="%{${seq}%}"
    elif [[ $GBT_SHELL == '_bash' ]]; then
        GBT__RETVAL="\\[${seq}\\]"
    elif [[ $GBT_SHELL == 'plain' ]]; then
        GBT__RETVAL="$seq"
    else
        # bash
        GBT__RETVAL="\x01${seq}\x02"
    fi
}


function GbtDecorateElement() {
    local element=$1
    local text=$2
    local bg=$3
    local fg=$4
    local fm=$5

    local fmEnd=''
    local root=''

    if [[ $element != '' ]]; then
        GbtGetColor "${GBT_CAR["model-${element}-Bg"]}" 0
        bg=$GBT__RETVAL
        GbtGetColor "${GBT_CAR["model-${element}-Fg"]}" 1
        fg=$GBT__RETVAL
        GbtGetFormat "${GBT_CAR["model-${element}-Fm"]}" 0
        fm=$GBT__RETVAL

        if [[ $element == 'root' ]]; then
            text=''
        else
            GbtDecorateElement 'root'
            root=$GBT__RETVAL
            text="${GBT_CAR["model-${element}-Text"]}"
        fi

        GbtGetFormat 'empty' 0

        if [[ $fm != "$GBT__RETVAL" ]]; then
            GbtGetFormat "${GBT_CAR["model-${element}-Fm"]}" 1
            fmEnd=$GBT__RETVAL
        fi
    fi

    GBT__RETVAL="$bg$fg$fm$text$fmEnd$root"
}

function GbtFormatCar() {
    GbtDecorateElement 'root' "${GBT_CAR[model-root-Text]}"
    local text="$GBT__RETVAL${GBT_CAR[model-root-Text]}"
    local placeholder=',,,,'

    # shellcheck disable=SC2034
    for n in {0..10}; do
        local new_text
        new_text=$(echo "$text" | sed -E 's/\{\{\ *[a-zA-Z0-9]+\ *\}\}/'$placeholder'/')

        if [[ ${#new_text} == "${#text}" ]]; then
            break
        fi

        local before="${new_text%%"$placeholder"*}"
        local after="${new_text#*"$placeholder"}"
        local element="${text:$(( ${#before} + 2 )):$((${#text} - ${#after} - ${#before} - 4))}"
        element=${element// }

        GbtDecorateElement "$element" "${new_text//$placeholder/${GBT_CAR["model-${element}-Text"]}}"
        local replacement=$GBT__RETVAL

        if [ ${GBT_CAR["model-${element}-Text"]+1} ]; then
            text="${new_text//$placeholder/$replacement}"
        else
            text="${new_text//$placeholder/"\x7b\x7b ${element} \x7d\x7d"}"
        fi
    done

    echo -ne "$text"
}


function GbtMsg() {
    local type=$1
    local text=$2

    ( >&2 echo "$type: $text" )

    if [[ $type == 'E' ]]; then
        exit 1
    fi
}


function GbtMain() {
    local first=1
    local right=${GBT_RIGHT:-0}

    local prevBg="\x00"
    local prevDisplay=1
    local defaultSeparator='\xee\x82\xb0'

    if [[ $right == 1 ]]; then
        defaultSeparator='\xee\x82\xb2'
    fi

    if [[ $right != 1 ]] && [ "$GBT_BEGINNING_TEXT" != '' ]; then
        GbtGetColor "${GBT_BEGINNING_BG:-default}" 0
        local beginning_bg=$GBT__RETVAL
        GbtGetColor "${GBT_BEGINNING_FG:-default}" 1
        local beginning_fg=$GBT__RETVAL
        GbtGetFormat "${GBT_BEGINNING_FM:-default}" 0
        local beginning_fm=$GBT__RETVAL

        GbtDecorateElement '' "$beginning_bg" "$beginning_fg" "$beginning_fm" "$GBT_BEGINNING_TEXT"
        echo -en "$GBT__RETVAL"
    fi

    declare -A GBT_CAR

    for car in $(echo "${GBT_CARS:-status,os,hostname,dir,git,sign}" | sed -E 's/,\ */ /g' | tr '[:upper:]' '[:lower:]'); do
        GBT_CAR=()

        local unknown=0

        # Fill in the model
        if [ "$car" = 'aws' ]; then
            GbtCarAws
        elif [ "$car" = 'azure' ]; then
            GbtCarAzure
        elif [ "${car:0:6}" = 'custom' ]; then
            GbtCarCustom "${car:6}"
        elif [ "$car" = 'dir' ]; then
            GbtCarDir
        elif [ "$car" = 'exectime' ]; then
            GbtCarExecTime
        elif [ "$car" = 'gcp' ]; then
            GbtCarGcp
        elif [ "$car" = 'git' ]; then
            GbtCarGit
        elif [ "$car" = 'hostname' ]; then
            GbtCarHostname
        elif [ "$car" = 'kubectl' ]; then
            GbtCarKubectl
        elif [ "$car" = 'os' ]; then
            GbtCarOs
        elif [ "$car" = 'pyvirtenv' ]; then
            GbtCarPyVirtEnv
        elif [ "$car" = 'sign' ]; then
            GbtCarSign
        elif [ "$car" = 'status' ]; then
            GbtCarStatus "$@"
        elif [ "$car" = 'time' ]; then
            GbtCarTime
        else
            unknown=1
        fi

        if [[ $unknown == 0 ]] && [[ ${GBT_CAR[display]} == 1 ]]; then
            local separator=$defaultSeparator

            if [[ ${GBT_CAR[model-Sep-Text]} != "\x00" ]]; then
                separator=${GBT_CAR[model-Sep-Text]}
            fi

            GbtGetColor 'RESETALL' 0
            echo -en "$GBT__RETVAL"

            if [[ $prevBg != "\x00" ]] && [[ $prevDisplay == 1 ]]; then
                GbtGetColor "${GBT_CAR[model-root-Bg]}" 0
                local bg=$GBT__RETVAL
                GbtGetColor "${GBT_CAR[model-root-Bg]}" 1
                local fg=$GBT__RETVAL
                local fm=''

                if [[ ${GBT_CAR[wrap]} == 1 ]]; then
                    GbtGetColor 'default' 0
                    bg=$GBT__RETVAL
                    GbtGetColor 'default' 1
                    fg=$GBT__RETVAL
                fi

                if [[ $right == 1 ]]; then
                    GbtGetColor "$prevBg" 0
                    bg=$GBT__RETVAL
                else
                    GbtGetColor "$prevBg" 1
                    fg=$GBT__RETVAL
                fi

                if [[ ${GBT_CAR[model-Sep-Bg]} != "\x00" ]]; then
                    GbtGetColor "${GBT_CAR[model-Sep-Bg]}" 0
                    bg=$GBT__RETVAL
                fi

                if [[ ${GBT_CAR[model-Sep-Fg]} != "\x00" ]]; then
                    GbtGetColor "${GBT_CAR[model-Sep-Fg]}" 1
                    fg=$GBT__RETVAL
                fi

                if [[ ${GBT_CAR[model-Sep-Fm]} != "\x00" ]]; then
                    GbtGetFormat "${GBT_CAR[model-Sep-Fm]}" 0
                    fm=$GBT__RETVAL
                fi

                GbtDecorateUnicode "$separator"
                separator=$GBT__RETVAL

                GbtDecorateElement '' "$separator" "$bg" "$fg" "$fm"
                echo -en "$GBT__RETVAL"

                if [[ ${GBT_CAR[wrap]} == 1 ]]; then
                    echo
                fi
            fi

            prevBg=${GBT_CAR["model-root-Bg"]}
            prevDisplay=${GBT_CAR[display]}

            # Print the car
            GbtFormatCar

            # shellcheck disable=SC2034
            first=0
        fi
    done

    GbtGetColor 'RESETALL' 0
    echo -en "$GBT__RETVAL"
}
