declare -A GBT__OS_SYMBOLS
GBT__OS_SYMBOLS=(
    [amzn]='\xef\x94\xac'                [amzn_color]=208
    [android]='\xef\x85\xbb'             [android_color]=113
    [arch]='\xef\x8c\x83'                [arch_color]=25
    [archarm]='\xef\x8c\x83'             [archarm_color]=125
    [alpine]='\xef\x8c\x80'              [alpine_color]=24
    [aosc]='\xef\x8c\x81'                [aosc_color]=172
    [centos]='\xef\x8c\x84'              [centos_color]=27
    [cloud]='\xef\x99\x9e'               [cloud_color]=39
    [coreos]='\xef\x8c\x85'              [coreos_color]=32
    [darwin]='\xef\x94\xb4'              [darwin_color]=15
    [debian]='\xee\x9d\xbd'              [debian_color]=88
    [devuan]='\xef\x8c\x87'              [devuan_color]=16
    [docker]='\xee\x9e\xb0'              [docker_color]=26
    [elementary]='\xef\x8c\x89'          [elementary_color]=33
    [fedora]='\xef\x8c\x8a'              [fedora_color]=32
    [freebsd]='\xef\x8c\x8c'             [freebsd_color]=1
    [gentoo]='\xef\x8c\x8d'              [gentoo_color]=62
    [linux]='\xef\x85\xbc'               [linux_color]=15
    [linuxmint]='\xef\x8c\x8e'           [linuxmint_color]=47
    [mageia]='\xef\x8c\x90'              [mageia_color]=24
    [mandriva]='\xef\x8c\x91'            [mandriva_color]=208
    [manjaro]='\xef\x8c\x92'             [manjaro_color]=34
    [mysql]='\xee\x9c\x84'               [mysql_color]=30
    [nixos]='\xef\x8c\x93'               [nixos_color]=88
    [opensuse]='\xef\x8c\x94'            [opensuse_color]=113
    [opensuse-leap]='\xef\x8c\x94'       [opensuse-leap_color]=113
    [opensuse-tumbleweed]='\xef\x8c\x94' [opensuse-tumbleweed_color]=113
    [raspbian]='\xef\x8c\x95'            [raspbian_color]=125
    [rhel]='\xee\x9e\xbb'                [rhel_color]=1
    [sabayon]='\xef\x8c\x97'             [sabayon_color]=255
    [slackware]='\xef\x8c\x98'           [slackware_color]=63
    [sles]='\xef\x8c\x94'                [sles_color]=113
    [ubuntu]='\xef\x8c\x9b'              [ubuntu_color]=166
    [windows]='\xee\x98\xaa'             [windows_color]=6
)

function GbtCarOs() {
    local os

    if [ -n "$GBT_CAR_OS_NAME" ]; then
        os=$GBT_CAR_OS_NAME
    elif [ -e /proc/1/sched ] && [[ ! $(cat /proc/1/sched | head -n 1 | egrep '(init|systemd)') ]]; then
        os='docker'
    else
        os="$(uname -s)"

        if [ "$os" = 'Darwin' ]; then
            os='darwin'
        elif [ "$os" = 'Linux' ] && [ -e /etc/os-release ]; then
            os=$(source /etc/os-release; echo "$ID")

            if [ ! ${GBT__OS_SYMBOLS[$os]+1} ]; then
                os='linux'
            fi
        fi
    fi

    os=${os,,}

    local defaultRootBg=${GBT_CAR_BG:-235}
    local defaultRootFg=${GBT_CAR_FG:-white}
    local defaultRootFm=${GBT_CAR_FM:-none}

    GbtDecorateUnicode ${GBT_CAR_OS_SYMBOL_TEXT-${GBT__OS_SYMBOLS[$os]:-?}}
    local defaultSymbolText=$GBT__RETVAL

    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_OS_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_OS_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_OS_FM:-$defaultRootFm}
        [model-root-Text]=${GBT_CAR_OS_FORMAT-' {{ Symbol }} '}

        [model-Symbol-Bg]=${GBT_CAR_OS_SYMBOL_BG:-${GBT_CAR_OS_BG:-$defaultRootBg}}
        [model-Symbol-Fg]=${GBT_CAR_OS_SYMBOL_FG:-${GBT_CAR_OS_FG:-${GBT__OS_SYMBOLS["${os}_color"]:-none}}}
        [model-Symbol-Fm]=${GBT_CAR_OS_SYMBOL_FM:-${GBT_CAR_OS_FM:-$defaultRootFm}}
        [model-Symbol-Text]=$defaultSymbolText

        [display]=${GBT_CAR_OS_DISPLAY:-1}
        [wrap]=${GBT_CAR_OS_WRAP:-0}
        [sep]=${GBT_CAR_OS_SEP-'\x00'}
    )
}
