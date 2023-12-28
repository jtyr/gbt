# shellcheck shell=bash
function GbtCarGcp() {
    if [[ $GBT_CAR_GCP_DISPLAY == 0 ]]; then
        return
    fi

    local defaultRootBg=${GBT_CAR_BG:-33}
    local defaultRootFg=${GBT_CAR_FG:-white}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local defaultRootFormat=${GBT_CAR_GCP_FORMAT-' {{ Icon }} {{ Project }} '}
    local defaultConfigText=''
    local defaultAccountText=$CLOUDSDK_CORE_ACCOUNT
    local defaultProjectText=$CLOUDSDK_CORE_PROJECT
    local defaultSep="\x00"

    if [ -n "$CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE" ]; then
        defaultAccountText=$(grep -Eo '"client_email":\s*".[^"]+"' "$CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE" | sed -e 's/.*://' -e 's/"$//' -e 's/.*"//')

        # Indicate a problem if the client_email not found in the file
        if [ -z "$defaultAccountText" ]; then
            defaultAccountText='???'
        fi
    fi

    local configDir=${CLOUDSDK_CONFIG:-$HOME/.config/gcloud}
    local defaultConfigText
    defaultConfigText=$(cat "$configDir/active_config")
    local configFile="$configDir/configurations/config_$defaultConfigText"

    if [ -z "$defaultAccountText" ]; then
        defaultAccountText=$(sed -nr "/^\[core\]/ { :l /^account[ ]*=/ { s/.*=[ ]*//; p; q;}; n; b l;}" "$configFile")
    fi

    if [ -z "$defaultProjectText" ]; then
        defaultProjectText=$(sed -nr "/^\[core\]/ { :l /^project[ ]*=/ { s/.*=[ ]*//; p; q;}; n; b l;}" "$configFile")
    fi

    if [ -n "$GBT_CAR_GCP_PROJECT_ALIASES" ]; then
        orig_IFS=$IFS
        IFS=','

        for pair in $GBT_CAR_GCP_PROJECT_ALIASES; do
            IFS='='
            breakLoop=0

            for k in $pair; do
                k=${k//[[:space:]]/}

                if [[ $k == "$defaultProjectText" ]]; then
                    breakLoop=1

                    continue
                elif [[ $breakLoop == 0 ]]; then
                    break
                fi

                defaultProjectText=$k
            done

            if [[ $breakLoop == 1 ]]; then
                break
            fi
        done

        IFS=$orig_IFS
    fi

    GbtDecorateUnicode "${GBT_CAR_GCP_ICON_TEXT-'\xee\x9e\xb2'}"
    local defaultIconText=$GBT__RETVAL

    # shellcheck disable=SC2034
    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_GCP_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_GCP_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_GCP_FM:-$defaultRootFm}
        [model-root-Text]=$defaultRootFormat

        [model-Icon-Bg]=${GBT_CAR_GCP_ICON_BG:-${GBT_CAR_GCP_BG:-$defaultRootBg}}
        [model-Icon-Fg]=${GBT_CAR_GCP_ICON_FG:-${GBT_CAR_GCP_FG:-$defaultRootFg}}
        [model-Icon-Fm]=${GBT_CAR_GCP_ICON_FM:-${GBT_CAR_GCP_FM:-$defaultRootFm}}
        [model-Icon-Text]=$defaultIconText

        [model-Config-Bg]=${GBT_CAR_GCP_CONFIG_BG:-${GBT_CAR_GCP_BG:-$defaultRootBg}}
        [model-Config-Fg]=${GBT_CAR_GCP_CONFIG_FG:-${GBT_CAR_GCP_FG:-$defaultRootFg}}
        [model-Config-Fm]=${GBT_CAR_GCP_CONFIG_FM:-${GBT_CAR_GCP_FM:-$defaultRootFm}}
        [model-Config-Text]=${GBT_CAR_GCP_CONFIG_TEXT-$defaultConfigText}

        [model-Account-Bg]=${GBT_CAR_GCP_ACCOUNT_BG:-${GBT_CAR_GCP_BG:-$defaultRootBg}}
        [model-Account-Fg]=${GBT_CAR_GCP_ACCOUNT_FG:-${GBT_CAR_GCP_FG:-$defaultRootFg}}
        [model-Account-Fm]=${GBT_CAR_GCP_ACCOUNT_FM:-${GBT_CAR_GCP_FM:-$defaultRootFm}}
        [model-Account-Text]=${GBT_CAR_GCP_ACCOUNT_TEXT-$defaultAccountText}

        [model-Project-Bg]=${GBT_CAR_GCP_PROJECT_BG:-${GBT_CAR_GCP_BG:-$defaultRootBg}}
        [model-Project-Fg]=${GBT_CAR_GCP_PROJECT_FG:-${GBT_CAR_GCP_FG:-$defaultRootFg}}
        [model-Project-Fm]=${GBT_CAR_GCP_PROJECT_FM:-${GBT_CAR_GCP_FM:-$defaultRootFm}}
        [model-Project-Text]=${GBT_CAR_GCP_PROJECT_TEXT-$defaultProjectText}

        [model-Sep-Bg]=${GBT_CAR_GCP_SEP_BG:-$defaultSep}
        [model-Sep-Fg]=${GBT_CAR_GCP_SEP_FG:-$defaultSep}
        [model-Sep-Fm]=${GBT_CAR_GCP_SEP_FM:-$defaultSep}
        [model-Sep-Text]=${GBT_CAR_GCP_SEP_TEXT:-${GBT_CAR_GCP_SEP:-${GBT_SEPARATOR:-$defaultSep}}}

        [display]=${GBT_CAR_GCP_DISPLAY:-1}
        [wrap]=${GBT_CAR_GCP_WRAP:-0}
    )
}
