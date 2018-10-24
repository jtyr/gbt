function gbt_vagrant() {
    local VAGRANT_BIN=$(gbt__which vagrant)
    [ -z "$VAGRANT_BIN" ] && return 1

    if [ "$1" != 'ssh' ]; then
        $VAGRANT_BIN "$@"
    else
        shift

        local RDN=$RANDOM
        local GBT__CONF="/tmp/.gbt.$RDN"

        $VAGRANT_BIN ssh --command "cat /etc/motd 2>/dev/null;
export GBT__CONF='$GBT__CONF' &&
function gbt__$RDN() { echo '$((gbt__get_sources; echo 'gbt__ssh_theme') | eval "$GBT__SOURCE_COMPRESS" | $GBT__SOURCE_BASE64_LOCAL | tr -d '\r\n')' | $GBT__SOURCE_BASE64 $GBT__SOURCE_BASE64_DEC | $GBT__SOURCE_DECOMPRESS; };
if [ -z "$GBT__SOURCE_SEC_DISABLE" ]; then
  export GBT__CONF_MD5=\$(gbt__$RDN | tee $GBT__CONF | $GBT__SOURCE_MD5_REMOTE 2>/dev/null | cut -d' ' -f$GBT__SOURCE_MD5_CUT_REMOTE 2>/dev/null);
  echo '[ -z \"\$GBT__CONF_MD5\" ] && export GBT__CONF_MD5='\$GBT__CONF_MD5' || true' >> $GBT__CONF;
else
  gbt__$RDN > $GBT__CONF;
fi;
exec -a gbt.bash bash --rcfile \$GBT__CONF" "$@"
    fi
}
