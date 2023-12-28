# shellcheck shell=bash
function gbt_ssh() {
    local SSH_BIN
    SSH_BIN=$(gbt__which ssh)
    [ -z "$SSH_BIN" ] && return 1

    if ( gbt__is_ssh_command "$@" ); then
        $SSH_BIN "$@"
    else
        gbt__check_md5

        $SSH_BIN -t "$@" "cat /etc/motd 2>/dev/null;
echo \"$(eval "$GBT__SOURCE_COMPRESS" < '$GBT__CONF' | $GBT__SOURCE_BASE64 | tr -d '\r\n')\" | $GBT__SOURCE_BASE64 $GBT__SOURCE_BASE64_DEC | $GBT__SOURCE_DECOMPRESS > '$GBT__CONF' &&
chmod '$GBT__CONF_MODE' '$GBT__CONF';
exec -a gbt.bash bash --rcfile '$GBT__CONF'"
    fi
}
