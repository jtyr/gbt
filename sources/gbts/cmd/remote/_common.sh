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
