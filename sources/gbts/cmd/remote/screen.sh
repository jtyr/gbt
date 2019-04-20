function gbt_screen() {
    local SCREEN_BIN=$(gbt__which screen)
    [ -z "$SCREEN_BIN" ] && return 1

    gbt__check_md5

    $SCREEN_BIN -U -s "$GBT__CONF.bash" -t bash "$@"
}
