function gbt_docker() {
    local DOCKER_BIN=$(gbt__which docker)
    [ -z "$DOCKER_BIN" ] && return 1

    gbt__check_md5

    if [ "$1" != 'shell' ]; then
        $DOCKER_BIN "$@"
    else
        local GBT__CONTAINER_ID="${@: -1}"

        $DOCKER_BIN cp $GBT__CONF $GBT__CONTAINER_ID:$(dirname $GBT__CONF)
        $DOCKER_BIN exec ${@:2:$(( $# - 2 ))} -it $GBT__CONTAINER_ID bash -c "exec -a gbt.bash bash --rcfile $GBT__CONF"
        $DOCKER_BIN exec ${@:2:$(( $# - 2 ))} -it $GBT__CONTAINER_ID rm -f $GBT__CONF $GBT__CONF.bash
    fi
}
