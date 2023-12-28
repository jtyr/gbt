# shellcheck shell=bash
function gbt__is_vagrant_ssh_command() {
    # Parse through vagrant ssh to see if -c or --command is specified
    while [ $# -gt 0 ]; do
        if [[ $1 == '-c' ]] || [[ $1 == '--command' ]]; then
            return 0
        fi

        shift
    done

    return 1
}
