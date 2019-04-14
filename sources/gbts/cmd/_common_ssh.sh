function gbt__is_ssh_command() {
    # Parse through ssh command options and determine if there is a remote
    # command to be executed
    while [ $# -gt 0 ]; do
        # Check if it's an option and start with dash
        if [[ ${1:0:1} == '-' ]]; then
            # Check $1 is a option with argument, then do an extra shift
            if [[ 'BbcDEeFIiJLlmOopQRSWw' =~ ${1:1} ]]; then
                shift
            fi

            shift
        else
            # Shift over ssh destination
            shift

            if [[ -z "$@" ]]; then
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
