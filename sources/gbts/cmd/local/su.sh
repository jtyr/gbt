function gbt_su() {
    local SU_BIN=$(gbt__which su)
    [ -z "$SU_BIN" ] && return 1

    local GBT__CONF=$(gbt__local_rcfile)

    $SU_BIN -s "$GBT__CONF.bash" "$@"

    rm -f $GBT__CONF $GBT__CONF.bash
}
