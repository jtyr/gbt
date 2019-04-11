[ -z "$GBT__SOURCE_COMPRESS" ] && GBT__SOURCE_COMPRESS='gzip -qc9'
[ -z "$GBT__SOURCE_DECOMPRESS" ] && GBT__SOURCE_DECOMPRESS='gzip -qd'
[ -z "$GBT__SOURCE_BASE64" ] && GBT__SOURCE_BASE64='base64'
[ -z "$GBT__SOURCE_BASE64_DEC" ] && GBT__SOURCE_BASE64_DEC='-d'
[ -z "$GBT__SOURCE_MD5_CUT_LOCAL" ] && GBT__SOURCE_MD5_CUT_LOCAL='1'
[ -z "$GBT__SOURCE_MD5_CUT_REMOTE" ] && GBT__SOURCE_MD5_CUT_REMOTE='1'
[ -z "$GBT__SOURCE_MD5_LOCAL" ] && GBT__SOURCE_MD5_LOCAL='md5sum'
[ -z "$GBT__SOURCE_MD5_REMOTE" ] && GBT__SOURCE_MD5_REMOTE='md5sum'
[ -z "$GBT__CONF_MODE" ] && GBT__CONF_MODE='0644'
[ -z "$GBT__CONF_BASH_MODE" ] && GBT__CONF_BASH_MODE='0755'


function gbt__which() {
    local PROG=$1

    if [ -z "$GBT__WHICH" ]; then
        # Ignore aliases when using 'which'
        if [ "$(ps -p $$ 2>/dev/null | awk '$1 != "PID" {print $4}' | sed -e 's,.*/,,' -e 's/^-//')" = 'zsh' ]; then
            GBT__WHICH_OPTS='-p'

            if [ -z "$GBT__WHICH" ]; then
                # ZSH has built in which command
                GBT__WHICH='which'
            fi
        else
            # Run which
            GBT__WHICH=$(which which 2>/dev/null)

            if [ $? -ne 0 ]; then
                # Fail if there is no which
                GBT__WHICH=''
            else
                if [ ! -e "$GBT__WHICH" ]; then
                    if [ -z "$GBT__WHICH_OPTS" ]; then
                        # If it's not a path, try to get a path by excluding aliases
                        GBT__WHICH_OPTS='--skip-alias'
                    fi

                    GBT__WHICH=$(which $GBT__WHICH_OPTS which 2>/dev/null)

                    if [ $? -ne 0 ] || [ ! -e "$GBT__WHICH" ]; then
                        # Fail if that didn't work or if the returned string isn't a path
                        GBT__WHICH=''
                    fi
                fi
            fi
        fi
    fi

    if [ -z "$GBT__WHICH" ]; then
        gbt__err "ERROR: 'which' not found"
        return 1
    fi

    GBT__WHICH_PROG_PATH=$($GBT__WHICH $GBT__WHICH_OPTS $PROG 2>/dev/null)

    if [ $? -ne 0 ]; then
        gbt__err "ERROR: '$PROG' not found"
        return 1
    fi

    echo $GBT__WHICH_PROG_PATH
}


function gbt__err() {
    echo "$@" >&2
}


function gbt__finish() {
    local MY_PPID=$$

    if [[ "$(ps -o comm= $(ps -o ppid= $MY_PPID))" == 'sshd' ]]; then
        # Cleanup at the end of the 'ssh' or 'vagrant' session
        rm -f $GBT__CONF $GBT__CONF.bash
    fi
}


function gbt__is_ssh_command() {
    # Parse through ssh command options and determine
    # If there is a remote command to be executed
    local SSH_DUAL_OPTIONS="BbcDEeFIiJLlmOopQRSWw"

    while [[ $# -gt 0 ]]; do
        # Check if it's an option and start with dash
        if [[ "${1:0:1}" == "-" ]]; then
            # Check $1 is a option with argument, then do an extra shift
            if [[ "$SSH_DUAL_OPTIONS" =~ "${1:1}" ]]; then
                shift
            fi

            shift
        else
            # Shift over ssh destination
            shift

            if [[ -z "$@" ]];then
                # No command specified to be executed on remote host
                return 1
            else
                # Command specified to be exexuted
                return 0
            fi

            break
        fi
    done
}


trap gbt__finish EXIT
