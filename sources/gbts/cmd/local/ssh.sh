function gbt_ssh() {
    local SSH_BIN=$(gbt__which ssh)
    [ -z "$SSH_BIN" ] && return 1

    if [[ " ${GBT__SSH_IGNORE[*]} " == *" ${@: -1} "* ]]; then
        $SSH_BIN "$@"
    else
        local RDN=$RANDOM
        local GBT__CONF="/tmp/.gbt.$RDN"

        $SSH_BIN -t "$@" "cat /etc/motd 2>/dev/null;
export GBT__CONF='$GBT__CONF';
export GBT__CONF_BASH_MODE='$GBT__CONF_BASH_MODE';
function gbt__$RDN() { echo '$((gbt__get_sources; echo 'gbt__ssh_theme') | eval "$GBT__SOURCE_COMPRESS" | $GBT__SOURCE_BASE64_LOCAL | tr -d '\r\n')' | $GBT__SOURCE_BASE64 $GBT__SOURCE_BASE64_DEC | $GBT__SOURCE_DECOMPRESS; };
if [ -z "$GBT__SOURCE_SEC_DISABLE" ]; then
  export GBT__CONF_MD5=\$(gbt__$RDN | tee $GBT__CONF | $GBT__SOURCE_MD5_REMOTE 2>/dev/null | cut -d' ' -f$GBT__SOURCE_MD5_CUT_REMOTE 2>/dev/null);
  echo '[ -z \"\$GBT__CONF_MD5\" ] && export GBT__CONF_MD5='\$GBT__CONF_MD5' || true' >> $GBT__CONF;
else
  gbt__$RDN > $GBT__CONF;
fi;
chmod $GBT__CONF_MODE \$GBT__CONF;
exec -a gbt.bash bash --rcfile \$GBT__CONF"
    fi
}
