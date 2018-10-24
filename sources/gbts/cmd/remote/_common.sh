function gbt__check_md5() {
    if [ -n "$GBT__CONF_MD5" ]; then
        local CAT_BIN=$(gbt__which cat)
        local CUT_BIN=$(gbt__which cut)
        local GREP_BIN=$(gbt__which grep)
        local MD5SUM_BIN=$(gbt__which $GBT__SOURCE_MD5_REMOTE)

        if [ -z "$CAT_BIN" ] || [ -z "$CUT_BIN" ] || [ -z "$GREP_BIN" ] || [ -z "$MD5SUM_BIN" ]; then
            gbt__err 'WARNING: Cannot verify content of the GBT config!'
        elif [ "$($CAT_BIN $GBT__CONF | $GREP_BIN -v 'export GBT__CONF_MD5=[0-9a-f]' | $MD5SUM_BIN | $CUT_BIN -d' ' -f$GBT__SOURCE_MD5_CUT_REMOTE)" != "$GBT__CONF_MD5" ]; then
            gbt__err 'SECURITY WARNING: GBT script has been changed! Exiting...'
            sleep 3
            exit 1
        fi
    fi
}

# Check Bash version
if [[ ${BASH_VERSINFO[0]} -lt 4 ]]; then
  gbt__err 'ERROR: Bash v4.x is required to run GBTS.'
  sleep 3
  exit 1
fi

# Create executable that is used as shell in 'su'
if [ ! -e "$GBT__CONF.bash" ]; then
    echo -e "#!/bin/bash\nexec -a gbt.bash bash --rcfile $GBT__CONF \"\$@\"" > $GBT__CONF.bash
    chmod +x $GBT__CONF.bash
fi

# Load remote Bash profile if it exists
if [ -f ~/.bash_profile ]; then
    source ~/.bash_profile
fi

# Load remote custom profile if it exists
if [ -f ~/.gbt_profile ]; then
    source ~/.gbt_profile
fi
