function gbt_gssh() {
    local GCLOUD_BIN=$(gbt__which gcloud)
    [ -z "$GCLOUD_BIN" ] && return 1

    if ( gbt__is_ssh_command "$@" ); then
        $GCLOUD_BIN compute ssh "$@"
    else
        gbt__check_md5

        $GCLOUD_BIN compute ssh "$@" --ssh-flag='-t' -- "cat /etc/motd 2>/dev/null;
echo \"$(cat $GBT__CONF | eval "$GBT__SOURCE_COMPRESS" | $GBT__SOURCE_BASE64 | tr -d '\r\n')\" | $GBT__SOURCE_BASE64 $GBT__SOURCE_BASE64_DEC | $GBT__SOURCE_DECOMPRESS > $GBT__CONF &&
chmod $GBT__CONF_MODE $GBT__CONF;
exec -a gbt.bash bash --rcfile $GBT__CONF"
    fi
}
