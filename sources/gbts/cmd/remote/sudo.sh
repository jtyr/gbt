function gbt_sudo() {
    local WHICH=$(which $GBT__WHICH_OPTS which 2>/dev/null)
    [ -z $WHICH ] && gbt__err "'which' not found" && return 1
    local SU_BIN=$(which $GBT__WHICH_OPTS su 2>/dev/null)
    [ $? -ne 0 ] && gbt__err "'su' not found" && return 1
    local SUDO_BIN=$(which $GBT__WHICH_OPTS sudo 2>/dev/null)
    [ $? -ne 0 ] && gbt__err "'sudo' not found" && return 1

    if [ "$1" != 'su' ]; then
        $SUDO_BIN "$@"
    else
        shift
        $SUDO_BIN $SU_BIN -s "$GBT__CONF.bash" "$@"
    fi
}
