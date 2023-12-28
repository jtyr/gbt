# shellcheck shell=bash
function gbt__check_md5() {
    if [ -n "$GBT__CONF_MD5" ]; then
        local CAT_BIN CUT_BIN GREP_BIN MD5SUM_BIN
        CAT_BIN=$(gbt__which cat)
        CUT_BIN=$(gbt__which cut)
        GREP_BIN=$(gbt__which grep)
        MD5SUM_BIN=$(gbt__which "$GBT__SOURCE_MD5_REMOTE")

        if [ -z "$CAT_BIN" ] || [ -z "$CUT_BIN" ] || [ -z "$GREP_BIN" ] || [ -z "$MD5SUM_BIN" ]; then
            gbt__err 'WARNING: Cannot verify content of the GBT config!'
        elif [ "$($CAT_BIN "$GBT__CONF" | $GREP_BIN -v 'export GBT__CONF_MD5=[0-9a-f]' | $MD5SUM_BIN | $CUT_BIN -d' ' "-f$GBT__SOURCE_MD5_CUT_REMOTE")" != "$GBT__CONF_MD5" ]; then
            trap '' 2
            gbt__err 'SECURITY WARNING: GBT script has been changed! Exiting...'
            sleep 3
            trap 2
            exit 1
        fi
    fi
}


function gbt__finish() {
    local MY_PID=$$
    local MY_PPID
    MY_PPID=$(ps -o ppid= $MY_PID 2>/dev/null)

    if [[ ${MY_PPID// /} != '0' ]] && [[ "$(ps -o comm= "$MY_PPID" 2>/dev/null)" == 'sshd' ]]; then
        rm -f "$GBT__CONF $GBT__CONF.bash"
    fi
}


# Cleanup at the end of the 'ssh' or 'vagrant' session
trap gbt__finish EXIT

# Check Bash version
if [[ ${BASH_VERSINFO[0]} -lt 4 ]]; then
  gbt__err 'WARNING: Bash v4.x is required to run GBTS. Executing Bash without GBTS.'
  bash
  exit $?
fi

# Create executable that is used as shell in 'su'
if [ ! -e "$GBT__CONF.bash" ]; then
    echo -e "#!/bin/bash\nexec -a gbt.bash bash --rcfile $GBT__CONF \"\$@\"" > "$GBT__CONF.bash"
    chmod "$GBT__CONF_BASH_MODE" "$GBT__CONF.bash"
fi

# Add sbin paths if defined
if [ -n "$GBT__CONF_SBIN_PATH" ]; then
    export PATH="$GBT__CONF_SBIN_PATH:$PATH"
fi

# Load remote Bash profile if it exists
if [ -f ~/.bash_profile ]; then
    # shellcheck disable=SC1090
    source ~/.bash_profile
fi

# Load remote custom profile if it exists
if [ -f ~/.gbt_profile ]; then
    # shellcheck disable=SC1090
    source ~/.gbt_profile
fi
