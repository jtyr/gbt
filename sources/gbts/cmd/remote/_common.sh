function gbt__check_md5() {
    if [ -n "$GBT__CONF_MD5" ]; then
        local WHICH=$(which $GBT__WHICH_OPTS which 2>/dev/null)
        [ -z $WHICH ] && gbt__err "'which' not found" && return 1
        local CAT=$(which $GBT__WHICH_OPTS cat 2>/dev/null)
        local CUT=$(which $GBT__WHICH_OPTS cut 2>/dev/null)
        local GREP=$(which $GBT__WHICH_OPTS grep 2>/dev/null)
        local MD5SUM=$(which $GBT__WHICH_OPTS $GBT__SOURCE_SEC_SUM_REMOTE 2>/dev/null)

        if [ -z "$CAT" ] || [ -z "$CUT" ] || [ -z "$GREP" ] || [ -z "$MD5SUM" ]; then
            gbt__err 'WARNING: Cannot verify content of the GBT config!'
        elif [ "$($CAT $GBT__CONF | $GREP -v 'export GBT__CONF_MD5=[0-9a-f]' | $MD5SUM | $CUT -d' ' -f$GBT__SOURCE_SEC_CUT_REMOTE)" != "$GBT__CONF_MD5" ]; then
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

# Load remote custom profile if it exists
if [ -e ~/.gbt_profile ]; then
    source ~/.gbt_profile
fi
