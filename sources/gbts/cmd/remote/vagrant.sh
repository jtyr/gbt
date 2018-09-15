function gbt_vagrant() {
    local WHICH=$(which $GBT__WHICH_OPTS which 2>/dev/null)
    [ -z $WHICH ] && gbt__err "'which' not found" && return 1
    local VAGRANT_BIN=$(which $GBT__WHICH_OPTS vagrant 2>/dev/null)
    [ $? -ne 0 ] && gbt__err "'vagrant' not found" && return 1

    gbt__check_md5

    if [ "$1" != 'ssh' ]; then
        $VAGRANT_BIN "$@"
    else
        shift

        $VAGRANT_BIN ssh --command "cat /etc/motd 2>/dev/null;
echo \"$(cat $GBT__CONF | eval "$GBT__SOURCE_COMPRESS" | $GBT__SOURCE_BASE64 | tr -d '\r\n')\" | $GBT__SOURCE_BASE64 $GBT__SOURCE_BASE64_DEC | $GBT__SOURCE_DECOMPRESS > $GBT__CONF &&
exec -a gbt.bash bash --rcfile $GBT__CONF;
rm -f $GBT__CONF $GBT__CONF.bash" "$@"
    fi
}
