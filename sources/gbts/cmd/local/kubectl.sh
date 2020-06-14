function gbt_kubectl() {
    local KUBECTL_BIN=$(gbt__which kubectl)
    [ -z "$KUBECTL_BIN" ] && return 1

    if [ "$1" != 'shell' ]; then
        $KUBECTL_BIN "$@"
    else
        local GBT__POD_ID="${@: -1}"
        local GBT__CONF=$(gbt__local_rcfile)

        $KUBECTL_BIN cp ${@:2:$(( $# - 2 ))} $GBT__CONF $GBT__POD_ID:$(dirname $GBT__CONF)
        $KUBECTL_BIN exec ${@:2:$(( $# - 2 ))} -it $GBT__POD_ID -- bash -c "exec -a gbt.bash bash --rcfile $GBT__CONF"
        $KUBECTL_BIN exec ${@:2:$(( $# - 2 ))} -it $GBT__POD_ID -- rm -f $GBT__CONF $GBT__CONF.bash

        rm -f $GBT__CONF $GBT__CONF.bash
        unset GBT__CONF
    fi
}

[[ ${GBT__AUTO_ALIASES:-1} == 1 ]] && alias "${GBT__ALIASES[kubectl]}"='gbt_kubectl'
