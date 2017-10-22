# Function executed before every command run by the shell
function gbt_exectime_pre() {
    if [ -z $GBT_CAR_EXECTIME__TMP ]; then
      return
    fi

    unset GBT_CAR_EXECTIME__TMP

    export GBT_CAR_EXECTIME_SECS=$(date '+%s.%N')
}

# Function executed after every command run by the shell
function gbt_exectime_post() {
    GBT_CAR_EXECTIME__TMP=1

    # The rest of the function is only necessary if you want to ring the system
    # bell if the command is taking more that GBT_CAR_EXECTIME_BELL seconds.
    local SECS=${GBT_CAR_EXECTIME_SECS:-0}
    local BELL=${GBT_CAR_EXECTIME_BELL:-0}

    if (( $(echo "$SECS > 0" | bc) )) && (( $BELL > 0 )); then
        local EXECS=$(echo "$(date '+%s.%N') - $GBT_CAR_EXECTIME_SECS" | bc)

        if (( $(echo "$EXECS > $BELL" | bc) )); then
            echo -en '\a'
        fi
    fi
}

trap 'gbt_exectime_pre' DEBUG
PROMPT_COMMAND='gbt_exectime_post'
