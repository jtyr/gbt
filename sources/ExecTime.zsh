# Function executed before every command run by the shell
function gbt_exectime_pre() {
    export GBT_CAR_EXECTIME_SECS=$(date '+%s.%N')
    GBT_CAR_EXECTIME__TMP=1
}

# Function executed after every command run by the shell
function gbt_exectime_post() {
    if [ -z $GBT_CAR_EXECTIME__TMP ]; then
        export GBT_CAR_EXECTIME_SECS=$(date '+%s.%N')
    else
        # This "else" part is only necessary if you want to ring the system
        # bell if the command is taking more that GBT_CAR_EXECTIME_BELL
        # seconds.
        local BELL=${GBT_CAR_EXECTIME_BELL:-0}

        if (( $BELL > 0 )); then
            local EXECS=$(echo "$(date '+%s.%N') - $GBT_CAR_EXECTIME_SECS" | bc)

            if (( $EXECS > $BELL )); then
                echo -en '\a'
            fi
        fi
    fi

    unset GBT_CAR_EXECTIME__TMP
}

preexec_functions+=(gbt_exectime_pre)
precmd_functions+=(gbt_exectime_post)
