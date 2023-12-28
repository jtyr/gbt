# shellcheck shell=bash
function GbtCarAzure() {
    if [[ $GBT_CAR_AZURE_DISPLAY == 0 ]]; then
        return
    fi

    local defaultRootBg=${GBT_CAR_BG:-32}
    local defaultRootFg=${GBT_CAR_FG:-white}
    local defaultRootFm=${GBT_CAR_FM:-none}

    local defaultRootFormat=${GBT_CAR_AZURE_FORMAT-' {{ Icon }} {{ Subscription }} '}
    local defaultSep="\x00"

    local cloudText=$AZURE_CLOUD_NAME
    local subscriptionText=''
    local userNameText=''
    local userTypeText=''
    local stateText=''
    local defaultsGroupText=$AZURE_DEFAULTS_GROUP

    confDir=${AZURE_CONFIG_DIR:-$HOME/.azure}

    if [ -z "$cloudText" ]; then
        configFile="$confDir/config"

        # Get the current Cloud Name
        cloudText=$(sed -nr "/^\[cloud\]/ { :l /^name[ ]*=/ { s/.*=[ ]*//; p; q;}; n; b l;}" "$configFile")

        # Get the default Resource Group
        if [ -z "$defaultsGroupText" ]; then
            defaultsGroupText=$(sed -nr "/^\[defaults\]/ { :l /^group[ ]*=/ { s/.*=[ ]*//; p; q;}; n; b l;}" "$configFile")
        fi

        if [ -z "$cloudText" ]; then
            # Default Cloud Name
            cloudText='AzureCloud'
        fi
    fi

    # Get the Subscription ID
    if [[ -n $cloudText ]]; then
        cloudsConfigFile="$confDir/clouds.config"

        subscrId=$(sed -nr "/^\[$cloudText\]/ { :l /^subscription[ ]*=/ { s/.*=[ ]*//; p; q;}; n; b l;}" "$cloudsConfigFile")

        # Get the Subscription Name, User Name, User Type and State
        if [ -n "$subscrId" ]; then
            azureProfileFile="$confDir/azureProfile.json"

            # Flatten the JSON and extract just the matching record
            record=$(sed -r -e 's/^.[^[]+\[//' -e 's/\]}$//' -e 's/"user": \{"name"/"user_name"/g' -e 's/"type"(: ".[^"]+")\}/"user_type"\1/g' -e 's/.*("id": "'"$subscrId"'".[^}]*).*/ \1/' "$azureProfileFile")

            # Check if we got valid record and extract individual values
            if [[ ${record:0:1} != '{' ]]; then
                orig_IFS=$IFS
                IFS=','

                for pair in $record; do
                    k=${pair%": "*}
                    k=${k:2:-1}
                    v=${pair#*": "}
                    v=${v:1:-1}

                    case $k in
                        name)
                            subscriptionText=$v
                            ;;
                        user_name)
                            userNameText=$v
                            ;;
                        user_type)
                            userTypeText=$v
                            ;;
                        state)
                            stateText=$v
                            ;;
                    esac
                done

                IFS=$orig_IFS
            fi
        fi
    fi

    GbtDecorateUnicode "${GBT_CAR_AZURE_ICON_TEXT-'\xef\xb4\x83'}"
    local defaultIconText=$GBT__RETVAL

    # shellcheck disable=SC2034
    GBT_CAR=(
        [model-root-Bg]=${GBT_CAR_AZURE_BG:-$defaultRootBg}
        [model-root-Fg]=${GBT_CAR_AZURE_FG:-$defaultRootFg}
        [model-root-Fm]=${GBT_CAR_AZURE_FM:-$defaultRootFm}
        [model-root-Text]=$defaultRootFormat

        [model-Icon-Bg]=${GBT_CAR_AZURE_ICON_BG:-${GBT_CAR_AZURE_BG:-$defaultRootBg}}
        [model-Icon-Fg]=${GBT_CAR_AZURE_ICON_FG:-${GBT_CAR_AZURE_FG:-$defaultRootFg}}
        [model-Icon-Fm]=${GBT_CAR_AZURE_ICON_FM:-${GBT_CAR_AZURE_FM:-$defaultRootFm}}
        [model-Icon-Text]=$defaultIconText

        [model-Cloud-Bg]=${GBT_CAR_AZURE_CLOUD_BG:-${GBT_CAR_AZURE_BG:-$defaultRootBg}}
        [model-Cloud-Fg]=${GBT_CAR_AZURE_CLOUD_FG:-${GBT_CAR_AZURE_FG:-$defaultRootFg}}
        [model-Cloud-Fm]=${GBT_CAR_AZURE_CLOUD_FM:-${GBT_CAR_AZURE_FM:-$defaultRootFm}}
        [model-Cloud-Text]=${GBT_CAR_AZURE_CLOUD_TEXT-$cloudText}

        [model-Subscription-Bg]=${GBT_CAR_AZURE_SUBSCRIPTION_BG:-${GBT_CAR_AZURE_BG:-$defaultRootBg}}
        [model-Subscription-Fg]=${GBT_CAR_AZURE_SUBSCRIPTION_FG:-${GBT_CAR_AZURE_FG:-$defaultRootFg}}
        [model-Subscription-Fm]=${GBT_CAR_AZURE_SUBSCRIPTION_FM:-${GBT_CAR_AZURE_FM:-$defaultRootFm}}
        [model-Subscription-Text]=${GBT_CAR_AZURE_SUBSCRIPTION_TEXT-$subscriptionText}

        [model-UserName-Bg]=${GBT_CAR_AZURE_USERNAME_BG:-${GBT_CAR_AZURE_BG:-$defaultRootBg}}
        [model-UserName-Fg]=${GBT_CAR_AZURE_USERNAME_FG:-${GBT_CAR_AZURE_FG:-$defaultRootFg}}
        [model-UserName-Fm]=${GBT_CAR_AZURE_USERNAME_FM:-${GBT_CAR_AZURE_FM:-$defaultRootFm}}
        [model-UserName-Text]=${GBT_CAR_AZURE_USERNAME_TEXT-$userNameText}

        [model-UserType-Bg]=${GBT_CAR_AZURE_USERTYPE_BG:-${GBT_CAR_AZURE_BG:-$defaultRootBg}}
        [model-UserType-Fg]=${GBT_CAR_AZURE_USERTYPE_FG:-${GBT_CAR_AZURE_FG:-$defaultRootFg}}
        [model-UserType-Fm]=${GBT_CAR_AZURE_USERTYPE_FM:-${GBT_CAR_AZURE_FM:-$defaultRootFm}}
        [model-UserType-Text]=${GBT_CAR_AZURE_USERTYPE_TEXT-$userTypeText}

        [model-State-Bg]=${GBT_CAR_AZURE_STATE_BG:-${GBT_CAR_AZURE_BG:-$defaultRootBg}}
        [model-State-Fg]=${GBT_CAR_AZURE_STATE_FG:-${GBT_CAR_AZURE_FG:-$defaultRootFg}}
        [model-State-Fm]=${GBT_CAR_AZURE_STATE_FM:-${GBT_CAR_AZURE_FM:-$defaultRootFm}}
        [model-State-Text]=${GBT_CAR_AZURE_STATE_TEXT-$stateText}

        [model-DefaultsGroup-Bg]=${GBT_CAR_AZURE_DEFAULTS_GROUP_BG:-${GBT_CAR_AZURE_BG:-$defaultRootBg}}
        [model-DefaultsGroup-Fg]=${GBT_CAR_AZURE_DEFAULTS_GROUP_FG:-${GBT_CAR_AZURE_FG:-$defaultRootFg}}
        [model-DefaultsGroup-Fm]=${GBT_CAR_AZURE_DEFAULTS_GROUP_FM:-${GBT_CAR_AZURE_FM:-$defaultRootFm}}
        [model-DefaultsGroup-Text]=${GBT_CAR_AZURE_DEFAULTS_GROUP_TEXT-$defaultsGroupText}

        [model-Sep-Bg]=${GBT_CAR_AZURE_SEP_BG:-$defaultSep}
        [model-Sep-Fg]=${GBT_CAR_AZURE_SEP_FG:-$defaultSep}
        [model-Sep-Fm]=${GBT_CAR_AZURE_SEP_FM:-$defaultSep}
        [model-Sep-Text]=${GBT_CAR_AZURE_SEP_TEXT:-${GBT_CAR_AZURE_SEP:-${GBT_SEPARATOR:-$defaultSep}}}

        [display]=${GBT_CAR_AZURE_DISPLAY:-1}
        [wrap]=${GBT_CAR_AZURE_WRAP:-0}
    )
}
