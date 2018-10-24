function gbt_sudo() {
    local SU_BIN=$(gbt__which su)
    [ -z "$SU_BIN" ] && return 1
    local SUDO_BIN=$(gbt__ which sudo)
    [ -z "$SUDO_BIN" ] && return 1

    if [ "$1" != 'su' ]; then
        $SUDO_BIN "$@"
    else
        shift

        local GBT__CONF=$(gbt__local_rcfile)

        $SUDO_BIN $SU_BIN -s "$GBT__CONF.bash" "$@"
    fi

    rm -f $GBT__CONF $GBT__CONF.bash
}
