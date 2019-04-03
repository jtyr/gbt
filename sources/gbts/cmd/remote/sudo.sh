function gbt_sudo() {
    local SU_BIN=$(gbt__which su)
    [ -z "$SU_BIN" ] && return 1
    local SUDO_BIN=$(gbt__which sudo)
    [ -z "$SUDO_BIN" ] && return 1

    gbt__check_md5

    if [ "$1" != 'su' ] && [[ " $@ " != *" -i "* ]]; then
        $SUDO_BIN "$@"
    else
        shift

        $SUDO_BIN $SU_BIN -s "$GBT__CONF.bash" "$@"
    fi
}
