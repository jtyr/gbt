function gbt_mysql() {
    local MYSQL_BIN=$(gbt__which mysql)
    [ -z "$MYSQL_BIN" ] && return 1

    gbt__check_md5

    $MYSQL_BIN --prompt "$(gbt__mysql_theme; GbtMain)" "$@"
}
