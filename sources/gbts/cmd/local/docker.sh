# shellcheck shell=bash
function gbt_docker() {
    local DOCKER_BIN
    DOCKER_BIN=$(gbt__which docker)
    [ -z "$DOCKER_BIN" ] && return 1

    if [ "$1" != 'shell' ]; then
        $DOCKER_BIN "$@"
    else
        local GBT__CONTAINER_ID="${*: -1}"
        local GBT__CONF
        GBT__CONF=$(gbt__local_rcfile)

        $DOCKER_BIN cp "$GBT__CONF" "$GBT__CONTAINER_ID:$(dirname "$GBT__CONF")"
        $DOCKER_BIN exec "${@:2:$(( $# - 2 ))}" -it "$GBT__CONTAINER_ID" bash -c "exec -a gbt.bash bash --rcfile $GBT__CONF"
        $DOCKER_BIN exec "${@:2:$(( $# - 2 ))}" -it -u root "$GBT__CONTAINER_ID" rm -f "$GBT__CONF $GBT__CONF.bash"

        rm -f "$GBT__CONF $GBT__CONF.bash"
        unset GBT__CONF
    fi
}

# shellcheck disable=SC2139
[[ ${GBT__AUTO_ALIASES:-1} == 1 ]] && alias "${GBT__ALIASES[docker]}"='gbt_docker'
