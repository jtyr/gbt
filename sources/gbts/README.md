GBTS - GBT written in Shell
===========================

It's possible to use `gbts` instead of `gbt` as the prompt generator. It's not
recomended as `gbts` is abou 10x slower than `gbt`. If you want to do it anyway,
here is how:

```shell
export GBT__HOME=/path/to/gbt

### Local prompt
# ZSH
PROMPT='$($GBT__HOME/sources/gbts/gbts $?)'
# Bash
PS1='$($GBT__HOME/sources/gbts/gbts $?)'

source $GBT__HOME/sources/gbts/cmd/local.sh

# Local aliases
alias docker=gbt_docker
alias mysql=gbt_mysql
alias screen=gbt_screen
alias ssh=gbt_ssh
alias vagrant=gbt_vagrant

# Remote aliases
alias gbt___sudo=gbt_sudo
alias gbt___su=gbt_su
```


Additional settings
-------------------

The following variables must be configured before sourcing the `local.sh` file:

```shell
# Don't use 'gbt_ssh' when connecting to 'myhost1' or 'myhost2'
export GBT__SSH_IGNORE=(myhost1 myhost2)

# List of cars to pack for the remote.
# Should match or exceed the list of cars from the theme or the theme variables
# (GBT__THEME_REMOTE_CARS and GBT__THEME_MYSQL_CARS).
export GBT__CARS_REMOTE='Status, Os, Time, Hostname, Dir, Custom, Git, Sign'

# Set SSH theme file
export GBT__THEME_SSH=$GBT__HOME/sources/gbts/theme/local/basic.sh

# List of plugins to pack for the remote
export GBT__PLUGINS_REMOTE='docker,mysql,screen,ssh,su,sudo,vagrant'

# List of plugins to pack for local commands
export GBT__PLUGINS_LOCAL='docker,mysql,screen,ssh,su,sudo,vagrant'

# Change default name of the alias for SSH (similarly of other plugins)
export GBT__SSH_ALIAS='sssh'

# Unalias some of the remotely created aliases
export GBT__UNALIAS_REMOTE='sudo su'

# Disable automatic aliases
export GBT__AUTO_ALIASES='0'

# Suppress code minimizing
export GBT__SOURCE_MINIMIZE='cat'

# Suppress code compressing
export GBT__SOURCE_COMPRESS='cat'

# Suppress code decompressin
export GBT__SOURCE_DECOMPRESS='cat'

# Disable verification of the GBT script content
export GBT__SOURCE_SEC_DISABLE=1

# Set md5 and cut index for MacOS (needded for the verification of the GBT script content)
export GBT__SOURCE_MD5_REMOTE=md5
export GBT__SOURCE_MD5_CUT_REMOTE=4

# When logging from MacOS to Linux (needded for the verification of the GBT script content)
export GBT__SOURCE_MD5_LOCAL=md5
export GBT__SOURCE_MD5_CUT_LOCAL=4
```


Development
-----------

Use the following command to generate the OS symbols for the `os` car script:

```shell
for N in $(grep 'nf-' pkg/cars/os/main.go | sed -r -e 's/",\s+}.*//' -e 's/",.*"/,/' -e 's/":.*"/:/' -e 's/.*"//'); do NAME=${N%%:*}; COLOR=${N##*,}; TMP=${N##*:}; ICON=${TMP%%,*}; echo "    [$NAME]='$(echo -ne $ICON | xxd -plain | sed 's/\(..\)/\\\\x\1/g')'     [${NAME}_color]=$COLOR"; done
```

Use the following command to convert Unicode character to code:

```shell
printf '\\u%02x\n' "'"
```

Use the following command to convert Unicode code to byte sequence:

```shell
echo -ne "\ue0b0" | xxd -plain | sed 's/\(..\)/\\x\1/g'
```

Use the following command to convert byte sequence to Unicode code:

```shell
printf '\\u%02x\n' "'"$'\xee\x82\xb0'
```
