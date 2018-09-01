function gbt_mysql() {
    local WHICH=$(which $GBT__WHICH_OPTS which 2>/dev/null)
    [ -z $WHICH ] && gbt__err "'which' not found" && return 1
    local MYSQL_BIN=$(which $GBT__WHICH_OPTS mysql 2>/dev/null)
    [ $? -ne 0 ] && gbt__err "'mysql' not found" && return 1

    $MYSQL_BIN --prompt "$(gbt__mysql_theme; GbtMain)" "$@"
}
