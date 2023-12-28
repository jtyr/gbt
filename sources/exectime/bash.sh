# shellcheck shell=bash

# Function executed before every command run by the shell
function gbt_exectime_pre() {
    if [ -z "$GBT__EXECTIME_TMP" ]; then
        return
    fi

    unset GBT__EXECTIME_TMP

    GBT_CAR_EXECTIME_SECS=$(${GBT__SOURCE_DATE:-date} "${GBT__SOURCE_DATE_ARG:-+%s.%N}")
    export GBT_CAR_EXECTIME_SECS
}

# Function executed after every command run by the shell
function gbt_exectime_post() {
    GBT__EXECTIME_TMP=1

    # The rest of the function is only necessary if you want to ring the system
    # bell if the command is taking more that GBT_CAR_EXECTIME_BELL seconds.
    local SECS=${GBT_CAR_EXECTIME_SECS:-0}
    local BELL=${GBT_CAR_EXECTIME_BELL:-0}

    if [ "$BELL" -gt 0 ] && [ "$SECS" -gt 0 ]; then
        local EXECS
        EXECS=$(echo "$(${GBT__SOURCE_DATE:-date} "${GBT__SOURCE_DATE_ARG:-+%s.%N}") - $GBT_CAR_EXECTIME_SECS" | bc)

        if [ "$EXECS" -gt "$BELL" ]; then
            echo -en '\a'
        fi
    fi
}

trap 'gbt_exectime_pre' DEBUG
PROMPT_COMMAND='gbt_exectime_post'
