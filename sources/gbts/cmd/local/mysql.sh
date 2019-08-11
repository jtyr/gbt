function gbt_mysql() {
    local MYSQL_BIN=$(gbt__which mysql)
    [ -z "$MYSQL_BIN" ] && return 1

    [ -z "$GBT__THEME_MYSQL" ] && local GBT__THEME_MYSQL="$GBT__HOME/sources/gbts/theme/mysql/${GBT__THEME_MYSQL_NAME:-default.sh}"

    $MYSQL_BIN --prompt "$(source $GBT__THEME_MYSQL; $GBT__HOME/sources/gbts/gbts)" "$@"
}

[[ ${GBT__AUTO_ALIASES:-1} == 1 ]] && alias "${GBT__ALIASES[mysql]}"='gbt_mysql'
