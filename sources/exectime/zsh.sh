# Function executed before every command run by the shell
function gbt_exectime_pre() {
    export GBT_CAR_EXECTIME_SECS=$(${GBT__SOURCE_DATE:-date} "${GBT__SOURCE_DATE_ARG:-+%s.%N}")
    GBT__EXECTIME_TMP=1
}

# Function executed after every command run by the shell
function gbt_exectime_post() {
    if [ -z $GBT__EXECTIME_TMP ]; then
        export GBT_CAR_EXECTIME_SECS=$(${GBT__SOURCE_DATE:-date} "${GBT__SOURCE_DATE_ARG:-+%s.%N}")
    else
        # This "else" part is only necessary if you want to ring the system
        # bell if the command is taking more that GBT_CAR_EXECTIME_BELL
        # seconds.
        local BELL=${GBT_CAR_EXECTIME_BELL:-0}

        if [ "$BELL" -gt 0 ]; then
            local EXECS=$(( $(${GBT__SOURCE_DATE:-date} "${GBT__SOURCE_DATE_ARG:-+%s.%N}") - $GBT_CAR_EXECTIME_SECS ))

            if [ "$EXECS" -gt "$BELL" ]; then
                echo -en '\a'
            fi
        fi
    fi

    unset GBT__EXECTIME_TMP
}

preexec_functions+=(gbt_exectime_pre)
precmd_functions+=(gbt_exectime_post)
