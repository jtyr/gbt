# shellcheck shell=bash
function gbt_su() {
    local SU_BIN
    SU_BIN=$(gbt__which su)
    [ -z "$SU_BIN" ] && return 1

    gbt__check_md5

    $SU_BIN -s "$GBT__CONF.bash" "$@"
}
