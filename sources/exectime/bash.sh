# Allow to override the date command (e.g. by 'gdate' on Mac)
if [ -z "$GBT__SOURCE_DATE" ]; then
    export GBT__SOURCE_DATE='date'
fi

# Allow to override the date argument.
# See https://github.com/jtyr/gbt/issues/14
if [ -z "$GBT__SOURCE_DATE_ARG" ]; then
    export GBT__SOURCE_DATE_ARG='+%s.%N'
fi

# Function executed before every command run by the shell
function gbt_exectime_pre() {
    if [ -z $GBT__EXECTIME_TMP ]; then
        return
    fi

    unset GBT__EXECTIME_TMP

    export GBT_CAR_EXECTIME_SECS=$($GBT__SOURCE_DATE "$GBT__SOURCE_DATE_ARG")
}

# Function executed after every command run by the shell
function gbt_exectime_post() {
    GBT__EXECTIME_TMP=1

    # The rest of the function is only necessary if you want to ring the system
    # bell if the command is taking more that GBT_CAR_EXECTIME_BELL seconds.
    local SECS=${GBT_CAR_EXECTIME_SECS:-0}
    local BELL=${GBT_CAR_EXECTIME_BELL:-0}

    if [ "$BELL" -gt 0 ] && [ "$SECS" -gt 0 ]; then
        local EXECS=$(echo "$(GBT_CAR_EXECTIME__DATE "$GBT_CAR_EXECTIME__DATE_ARG") - $GBT_CAR_EXECTIME_SECS" | bc)

        if [ "$EXECS" -gt "$BELL" ]; then
            echo -en '\a'
        fi
    fi
}

trap 'gbt_exectime_pre' DEBUG
PROMPT_COMMAND='gbt_exectime_post'
