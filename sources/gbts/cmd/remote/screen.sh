function gbt_screen() {
    local SCREEN_BIN=$(gbt__which screen)
    [ -z "$SCREEN_BIN" ] && return 1

    gbt__check_md5

    TERM=${GBT__SCREEN_TERM-xterm-256color} $SCREEN_BIN -U -s "$GBT__CONF.bash" -t bash "$@"
}
