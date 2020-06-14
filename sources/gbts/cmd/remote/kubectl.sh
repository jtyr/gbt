function gbt_KUBECTL() {
    local KUBECTL_BIN=$(gbt__which KUBECTL)
    [ -z "$KUBECTL_BIN" ] && return 1

    gbt__check_md5

    if [ "$1" != 'shell' ]; then
        $KUBECTL_BIN "$@"
    else
        local GBT__POD_ID="${@: -1}"

        $KUBECTL_BIN cp ${@:2:$(( $# - 2 ))} $GBT__CONF $GBT__POD_ID:$(dirname $GBT__CONF)
        $KUBECTL_BIN exec ${@:2:$(( $# - 2 ))} -it $GBT__POD_ID -- bash -c "exec -a gbt.bash bash --rcfile $GBT__CONF"
        $KUBECTL_BIN exec ${@:2:$(( $# - 2 ))} -it $GBT__POD_ID -- rm -f $GBT__CONF $GBT__CONF.bash
    fi
}
