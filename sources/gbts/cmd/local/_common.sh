GBT__PLUGINS_LOCAL__HASH=". $(echo ${GBT__PLUGINS_LOCAL:-docker,mysql,screen,ssh,su,sudo,vagrant} | sed -E 's/,\ */ /g' | tr '[:upper:]' '[:lower:]') ."
GBT__PLUGINS_REMOTE__HASH=". $(echo ${GBT__PLUGINS_REMOTE:-docker,mysql,screen,ssh,su,sudo,vagrant} | sed -E 's/,\ */ /g' | tr '[:upper:]' '[:lower:]') ."
GBT__CARS_REMOTE__HASH=". $(echo ${GBT__CARS_REMOTE:-dir,git,hostname,os,sign,status,time} | sed -E 's/,\ */ /g' | tr '[:upper:]' '[:lower:]') ."

GBT__SOURCE_BASE64_LOCAL=${GBT__SOURCE_BASE64_LOCAL:-base64}

declare -A GBT__ALIASES
GBT__ALIASES[docker]=${GBT__DOCKER_ALIAS:-docker}
GBT__ALIASES[mysql]=${GBT__MYSQL_ALIAS:-mysql}
GBT__ALIASES[screen]=${GBT__SCREEN_ALIAS:-screen}
GBT__ALIASES[ssh]=${GBT__SSH_ALIAS:-ssh}
GBT__ALIASES[su]=${GBT__SU_ALIAS:-su}
GBT__ALIASES[sudo]=${GBT__SUDO_ALIAS:-sudo}
GBT__ALIASES[vagrant]=${GBT__VAGRANT_ALIAS:-vagrant}


function gbt__local_rcfile() {
    local GBT__CONF="/tmp/.gbt.$RANDOM"
    local MD5SUM=$(gbt__get_sources | tee $GBT__CONF | $GBT__SOURCE_MD5_LOCAL 2>/dev/null | cut -d' ' -f$GBT__SOURCE_MD5_CUT_LOCAL 2>/dev/null)

    if [ -z "$GBT__SOURCE_SEC_DISABLE" ]; then
        echo "export GBT__CONF_MD5=${GBT__CONF_MD5:-$MD5SUM}" >> $GBT__CONF
    else
        echo 'export GBT__SOURCE_SEC_DISABLE=1' >> $GBT__CONF
    fi

    chmod $GBT__CONF_MODE $GBT__CONF

    echo -e "#!/bin/bash\nexec -a gbt.bash bash --rcfile $GBT__CONF \"\$@\"" > $GBT__CONF.bash
    chmod $GBT__CONF_BASH_MODE $GBT__CONF.bash

    echo $GBT__CONF
}


function gbt__get_sources_cars() {
    local C=$1

    [ "$C" = 'exectime' ] && cat $GBT__HOME/sources/exectime/bash.sh
    [ "${C:0:6}" = 'custom' ] && C=${C:0:6}

    if [ -f $GBT__HOME/sources/gbts/car/$C.sh ]; then
        cat $GBT__HOME/sources/gbts/car/$C.sh
    fi
}


function gbt__get_sources() {
    [ -z "$GBT__HOME" ] && gbt__err "'GBT__HOME' not defined" && return 1

    GBT__SOURCE_MINIMIZE=${GBT__SOURCE_MINIMIZE:-"sed -E -e '/^\\ *#.*/d' -e '/^\\ *$/d' -e 's/^\\ +//g' -e 's/default([A-Z])/d\\1/g' -e 's/model-/m-/g' -e 's/\\ {2,}/ /g'"}

    # Conditional for remote only (GBT__PLUGINS_REMOTE)
    [[ ${GBT__PLUGINS_REMOTE__HASH[@]} == *' ssh '* ]] && local GBT__THEME_SSH=${GBT__THEME_SSH:-$GBT__HOME/sources/gbts/theme/ssh/${GBT__THEME_SSH_NAME:-default}.sh}
    [[ ${GBT__PLUGINS_REMOTE__HASH[@]} == *' mysql '* ]] && local GBT__THEME_MYSQL=${GBT__THEME_MYSQL:-$GBT__HOME/sources/gbts/theme/mysql/${GBT__THEME_MYSQL_NAME:-default}.sh}

    (
        echo "export GBT__CONF='$GBT__CONF'"
        echo "export GBT__CONF_SBIN_PATH='$GBT__CONF_SBIN_PATH'"
        cat $GBT__HOME/sources/gbts/{cmd{,/remote},car}/_common.sh

        # Automatic aliases
        if [[ ${GBT__AUTO_ALIASES:-1} == 1 ]]; then
            for plugin in $(echo $GBT__PLUGINS_REMOTE__HASH | sed 's/\ /\n/g'); do
                if [[ $plugin != '.' ]]; then
                    echo "alias ${GBT__ALIASES[$plugin]}='gbt_$plugin'"
                fi
            done

            if [[ -n $GBT__UNALIAS_REMOTE ]]; then
                echo "unalias $GBT__UNALIAS_REMOTE 2>/dev/null"
            fi
        fi

        # Include SSH common function if car is present
        if [[ ${GBT__PLUGINS_REMOTE__HASH[@]} == *' ssh '* ]]; then
            cat $GBT__HOME/sources/gbts/cmd/_common_ssh.sh
        fi

        # Include Vagrant common function if car is present
        if [[ ${GBT__PLUGINS_REMOTE__HASH[@]} == *' vagrant '* ]]; then
            cat $GBT__HOME/sources/gbts/cmd/_common_vagrant.sh
        fi

        # Preserver modes
        [ "$GBT__CONF_MODE" != '0600' ] && echo "export GBT__CONF_MODE='$GBT__CONF_MODE'"
        [ "$GBT__CONF_BASH_MODE" != '0755' ] && echo "export GBT__CONF_BASH_MODE='$GBT__CONF_BASH_MODE'"

        # Allow to override default list of cars defined in the theme
        [ -n "$GBT__THEME_REMOTE_CARS" ] && echo "export GBT__THEME_REMOTE_CARS='$GBT__THEME_REMOTE_CARS'"
        [ -n "$GBT__THEME_MYSQL_CARS" ] && echo "export GBT__THEME_MYSQL_CARS='$GBT__THEME_MYSQL_CARS'"

        # Security on the remote site
        if [ -z "$GBT__SOURCE_SEC_DISABLE" ]; then
            [ -n "$GBT__SOURCE_MD5_CUT_REMOTE" ] && echo "export GBT__SOURCE_MD5_CUT_REMOTE=\"\${GBT__SOURCE_MD5_CUT_LOCAL:-$GBT__SOURCE_MD5_CUT_REMOTE}\""
            [ -n "$GBT__SOURCE_MD5_REMOTE" ] && echo "export GBT__SOURCE_MD5_REMOTE=\"\${GBT__SOURCE_MD5_LOCAL:-$GBT__SOURCE_MD5_REMOTE}\""
        else
            echo 'export GBT__SOURCE_SEC_DISABLE=1'
        fi

        for P in $(echo $GBT__PLUGINS_REMOTE__HASH); do
            if [ -f $GBT__HOME/sources/gbts/cmd/remote/$P.sh ]; then
                cat $GBT__HOME/sources/gbts/cmd/remote/$P.sh
            fi
        done

        for C in $(echo $GBT__CARS_REMOTE__HASH); do
            gbt__get_sources_cars $C
        done

        if [[ ${GBT__PLUGINS_REMOTE__HASH[@]} == *' ssh '* ]]; then
            echo 'function gbt__ssh_theme() {'
            cat $GBT__THEME_SSH
            echo '}'
        fi

        if [[ ${GBT__PLUGINS_REMOTE__HASH[@]} == *' mysql '* ]]; then
            echo 'function gbt__mysql_theme() {'
            cat $GBT__THEME_MYSQL
            echo '}'
        fi

        [[ ${GBT__CARS_REMOTE__HASH[@]} == *' ssh '* ]] && [ -n "$GBT__THEME_SSH_CARS" ] && echo "export GBT__THEME_SSH_CARS='$GBT__THEME_SSH_CARS'"
        alias | awk '/gbt___/ {sub(/^(alias )?(gbt___)?/, "", $0); print "alias "$0}'
        echo "PS1='\$(GbtMain \$?)'"
    ) | eval "$GBT__SOURCE_MINIMIZE"
}
