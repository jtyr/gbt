# shellcheck shell=bash
# shellcheck disable=SC1091
source "$GBT__HOME/sources/gbts/cmd/_common.sh"
source "$GBT__HOME/sources/gbts/cmd/local/_common.sh"

if echo "${GBT__PLUGINS_LOCAL__HASH[@]}" | grep -qFw 'docker'; then
    source "$GBT__HOME/sources/gbts/cmd/local/docker.sh"
fi
if echo "${GBT__PLUGINS_LOCAL__HASH[@]}" | grep -qFw 'gssh'; then
    source "$GBT__HOME/sources/gbts/cmd/local/gssh.sh"
    source "$GBT__HOME/sources/gbts/cmd/_common_ssh.sh"
fi
if echo "${GBT__PLUGINS_LOCAL__HASH[@]}" | grep -qFw 'kubectl'; then
    source "$GBT__HOME/sources/gbts/cmd/local/kubectl.sh"
fi
if echo "${GBT__PLUGINS_LOCAL__HASH[@]}" | grep -qFw 'mysql'; then
    source "$GBT__HOME/sources/gbts/cmd/local/mysql.sh"
fi
if echo "${GBT__PLUGINS_LOCAL__HASH[@]}" | grep -qFw 'screen'; then
    source "$GBT__HOME/sources/gbts/cmd/local/screen.sh"
fi
if echo "${GBT__PLUGINS_LOCAL__HASH[@]}" | grep -qFw 'ssh'; then
    source "$GBT__HOME/sources/gbts/cmd/local/ssh.sh"
    source "$GBT__HOME/sources/gbts/cmd/_common_ssh.sh"
fi
if echo "${GBT__PLUGINS_LOCAL__HASH[@]}" | grep -qFw 'su'; then
    source "$GBT__HOME/sources/gbts/cmd/local/su.sh"
fi
if echo "${GBT__PLUGINS_LOCAL__HASH[@]}" | grep -qFw 'sudo'; then
    source "$GBT__HOME/sources/gbts/cmd/local/sudo.sh"
fi
if echo "${GBT__PLUGINS_LOCAL__HASH[@]}" | grep -qFw 'vagrant'; then
    source "$GBT__HOME/sources/gbts/cmd/local/vagrant.sh"
    source "$GBT__HOME/sources/gbts/cmd/_common_vagrant.sh"
fi
