# shellcheck shell=bash
function gbt_su() {
    local SU_BIN GBT__CONF

    SU_BIN=$(gbt__which su)
    [ -z "$SU_BIN" ] && return 1

    GBT__CONF=$(gbt__local_rcfile)

    $SU_BIN -s "$GBT__CONF.bash" "$@"

    rm -f "$GBT__CONF" "$GBT__CONF.bash"
}

# shellcheck disable=SC2139
[[ ${GBT__AUTO_ALIASES:-1} == 1 ]] && alias "${GBT__ALIASES[su]}"='gbt_su'
