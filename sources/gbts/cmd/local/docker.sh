function gbt_docker() {
    local WHICH=$(which $GBT__WHICH_OPTS which 2>/dev/null)
    [ -z $WHICH ] && gbt__err "'which' not found" && return 1
    local DOCKER_BIN=$(which $GBT__WHICH_OPTS docker 2>/dev/null)
    [ $? -ne 0 ] && gbt__err "'docker' not found" && return 1

    if [ "$1" != 'shell' ]; then
        $DOCKER_BIN "$@"
    else
        local GBT__CONTAINER_ID="${@: -1}"
        local GBT__CONF=$(gbt__local_rcfile)

        $DOCKER_BIN cp $GBT__CONF $GBT__CONTAINER_ID:$(dirname $GBT__CONF)
        $DOCKER_BIN exec ${@:2:$(( $# - 2 ))} -it $GBT__CONTAINER_ID bash -c "exec -a gbt.bash bash --rcfile $GBT__CONF"
        $DOCKER_BIN exec ${@:2:$(( $# - 2 ))} -it $GBT__CONTAINER_ID rm -f $GBT__CONF $GBT__CONF.bash

        rm -f $GBT__CONF $GBT__CONF.bash
        unset GBT__CONF
    fi
}
