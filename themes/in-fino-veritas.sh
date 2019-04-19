export GBT_CARS='Time, Hostname, Dir, Git, Status'

# Reset all colors and formatting
export GBT_CAR_BG='default'
export GBT_CAR_FG='default'
export GBT_CAR_FM='default'

export GBT_SEPARATOR=''

export GBT_CAR_TIME_FORMAT='╭─[{{ Time }}]'
export GBT_CAR_TIME_FG='default'

export GBT_CAR_HOSTNAME_FORMAT=' {{ UserHost }}: '

export GBT_CAR_DIR_FORMAT='{{ Dir }}'
export GBT_CAR_DIR_DEPTH=9999
export GBT_CAR_DIR_FM='bold'

#export GBT_CAR_GIT_FM='bold'
export GBT_CAR_GIT_HEAD_FG='light_blue'
export GBT_CAR_GIT_HEAD_FM='bold'
export GBT_CAR_GIT_STATUS_DIRTY_FG='light_red'
export GBT_CAR_GIT_STATUS_CLEAN_FG='light_green'
export GBT_CAR_GIT_STATUS_MODIFIED_FORMAT='  {{ StatusModifiedCount }} {{ StatusModifiedSymbol }}'
export GBT_CAR_GIT_STATUS_MODIFIED_SYMBOL_TEXT='changed'
export GBT_CAR_GIT_STATUS_MODIFIED_COUNT_FM='bold'
export GBT_CAR_GIT_STATUS_MODIFIED_COUNT_FG='white'
export GBT_CAR_GIT_STATUS_DELETED_FORMAT='  {{ StatusDeletedCount }} {{ StatusDeletedSymbol }}'
export GBT_CAR_GIT_STATUS_DELETED_SYMBOL_TEXT='deleted'
export GBT_CAR_GIT_STATUS_DELETER_COUNT_FM='bold'
export GBT_CAR_GIT_STATUS_DELETER_COUNT_FG='white'
export GBT_CAR_GIT_STATUS_UNTRACKED_FORMAT='  {{ StatusUntrackedCount }} {{ StatusUntrackedSymbol }}'
export GBT_CAR_GIT_STATUS_UNTRACKED_SYMBOL_TEXT='untracked'
export GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_FM='bold'
export GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_FG='white'
export GBT_CAR_GIT_STATUS_STAGED_FORMAT='  {{ StatusStagedCount }} {{ StatusStagedSymbol }}'
export GBT_CAR_GIT_STATUS_STAGED_SYMBOL_TEXT='staged'
export GBT_CAR_GIT_FORMAT='┆ {{ Head }} {{ Status }}{{ StatusModified }}{{ StatusDeleted }}{{ StatusUntracked }}{{ StatusStaged }}'
export GBT_CAR_GIT_WRAP='yes'

export GBT_CAR_STATUS_FORMAT='╰─{{ Symbol }} '
export GBT_CAR_STATUS_OK_TEXT='○'
export GBT_CAR_STATUS_OK_FG='green'
export GBT_CAR_STATUS_ERROR_FG='red'
export GBT_CAR_STATUS_DISPLAY='yes'
export GBT_CAR_STATUS_WRAP='yes'
