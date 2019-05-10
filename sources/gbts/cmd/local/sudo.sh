function gbt_sudo() {
    local SU_BIN=$(gbt__which su)
    [ -z "$SU_BIN" ] && return 1
    local SUDO_BIN=$(gbt__which sudo)
    [ -z "$SUDO_BIN" ] && return 1

    local rv
    if [ "$1" != 'su' ] && [[ " $@ " != *" -i "* ]]; then
        $SUDO_BIN "$@"

        rv=$?
    else
        shift

        local GBT__CONF=$(gbt__local_rcfile)

        $SUDO_BIN $SU_BIN -s "$GBT__CONF.bash" "$@"

        rv=$?
        
        rm -f $GBT__CONF $GBT__CONF.bash
    fi

    return $rv
}
