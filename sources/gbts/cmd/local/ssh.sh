function remote_ssh_command {
    # Parse through ssh command options and determine
    # if there is a remote command to be executed
    local SSH_DUAL_OPTIONS="BbcDEeFIiJLlmOopQRSWw"
    while (( "$#" )); do
        #check if it's an option and start with dash
        if [[ "${1:0:1}" == "-" ]]; then
            #check $1 is a option with argument, then do an extra shift
            if [[ "$SSH_DUAL_OPTIONS" =~ "${1:1}" ]]; then
                shift
            fi
            shift
        else
            #shift over ssh destination
            shift
            if [[ -z "$@" ]];then
                # no command specified to be executed on remote host
                return 1
            else
                # command specified to be exexuted
                return 0
            fi
            break
        fi
    done
}

function gbt_ssh() {
    local SSH_BIN=$(gbt__which ssh)
    [ -z "$SSH_BIN" ] && return 1

    if [[ " ${GBT__SSH_IGNORE[*]} " == *" ${@: -1} "* ]]; then
        $SSH_BIN "$@"
    elif remote_ssh_command "$@"; then
        $SSH_BIN "$@"
    else
        local RND=$RANDOM
        local GBT__CONF="/tmp/.gbt.$RND"

        $SSH_BIN -t "$@" "cat /etc/motd 2>/dev/null;
export GBT__CONF='$GBT__CONF';
export GBT__CONF_BASH_MODE='$GBT__CONF_BASH_MODE';
function gbt__$RND() { echo '$((gbt__get_sources; echo 'gbt__ssh_theme') | eval "$GBT__SOURCE_COMPRESS" | $GBT__SOURCE_BASE64_LOCAL | tr -d '\r\n')' | $GBT__SOURCE_BASE64 $GBT__SOURCE_BASE64_DEC | $GBT__SOURCE_DECOMPRESS; };
if [ -z "$GBT__SOURCE_SEC_DISABLE" ]; then
  export GBT__CONF_MD5=\$(gbt__$RND | tee $GBT__CONF | $GBT__SOURCE_MD5_REMOTE 2>/dev/null | cut -d' ' -f$GBT__SOURCE_MD5_CUT_REMOTE 2>/dev/null);
  echo '[ -z \"\$GBT__CONF_MD5\" ] && export GBT__CONF_MD5='\$GBT__CONF_MD5' || true' >> $GBT__CONF;
else
  gbt__$RND > $GBT__CONF;
fi;
chmod $GBT__CONF_MODE \$GBT__CONF;
exec -a gbt.bash bash --rcfile \$GBT__CONF"
    fi
}
