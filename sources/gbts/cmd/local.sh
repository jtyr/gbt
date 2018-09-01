source $GBT__HOME/sources/gbts/cmd/_common.sh
source $GBT__HOME/sources/gbts/cmd/local/_common.sh

if [[ ${GBT__PLUGINS_LOCAL__HASH[@]} == *' docker '* ]]; then
    source $GBT__HOME/sources/gbts/cmd/local/docker.sh
fi
if [[ ${GBT__PLUGINS_LOCAL__HASH[@]} == *' mysql '* ]]; then
    source $GBT__HOME/sources/gbts/cmd/local/mysql.sh
fi
if [[ ${GBT__PLUGINS_LOCAL__HASH[@]} == *' screen '* ]]; then
    source $GBT__HOME/sources/gbts/cmd/local/screen.sh
fi
if [[ ${GBT__PLUGINS_LOCAL__HASH[@]} == *' ssh '* ]]; then
    source $GBT__HOME/sources/gbts/cmd/local/ssh.sh
fi
if [[ ${GBT__PLUGINS_LOCAL__HASH[@]} == *' su '* ]]; then
    source $GBT__HOME/sources/gbts/cmd/local/su.sh
fi
if [[ ${GBT__PLUGINS_LOCAL__HASH[@]} == *' sudo '* ]]; then
    source $GBT__HOME/sources/gbts/cmd/local/sudo.sh
fi
if [[ ${GBT__PLUGINS_LOCAL__HASH[@]} == *' vagrant '* ]]; then
    source $GBT__HOME/sources/gbts/cmd/local/vagrant.sh
fi
