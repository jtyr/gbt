function gbt_screen() {
    local SCREEN_BIN=$(gbt__which screen)
    [ -z "$SCREEN_BIN" ] && return 1

    local GBT__CONF=$(gbt__local_rcfile)

    $SCREEN_BIN -U -s "$GBT__CONF.bash" -t bash "$@"

    rm -f $GBT__CONF $GBT__CONF.bash
}

[[ ${GBT__AUTO_ALIASES:-1} == 1 ]] && alias "${GBT__ALIASES[screen]}"='gbt_screen'
