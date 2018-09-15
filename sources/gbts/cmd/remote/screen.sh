function gbt_screen() {
    local WHICH=$(which $GBT__WHICH_OPTS which 2>/dev/null)
    [ -z $WHICH ] && gbt__err "'which' not found" && return 1
    local SCREEN_BIN=$(which $GBT__WHICH_OPTS screen 2>/dev/null)
    [ $? -ne 0 ] && gbt__err "'screen' not found" && return 1

    gbt__check_md5

    $SCREEN_BIN -s "$GBT__CONF.bash" -t bash "$@"
}
