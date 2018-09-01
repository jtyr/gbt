function gbt_su() {
    local WHICH=$(which $GBT__WHICH_OPTS which 2>/dev/null)
    [ -z $WHICH ] && gbt__err "'which' not found" && return 1
    local SU_BIN=$(which $GBT__WHICH_OPTS su 2>/dev/null)
    [ $? -ne 0 ] && gbt__err "'su' not found" && return 1

    local GBT__CONF=$(gbt__local_rcfile)

    $SU_BIN -s "$GBT__CONF.bash" "$@"

    rm -f $GBT__CONF $GBT__CONF.bash
}
