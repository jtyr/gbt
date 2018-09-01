export GBT_CARS="${GBT__THEME_MYSQL_CARS:=Os, Time, Hostname, Dir, Sign}"

export GBT_CAR_OS_NAME='mysql'

export GBT_CAR_TIME_FORMAT=' {{ Time }} '
export GBT_CAR_TIME_TIME_TEXT='\R:\m:\s'

export GBT_CAR_HOSTNAME_FORMAT=' {{ User }}@{{ Host }} '
export GBT_CAR_HOSTNAME_USER_FG='cyan'
export GBT_CAR_HOSTNAME_USER_TEXT='\u'
export GBT_CAR_HOSTNAME_HOST_TEXT="$HOSTNAME"

export GBT_CAR_DIR_DIR_TEXT='\d'

export GBT_CAR_SIGN_FORMAT=' {{ User }} '
export GBT_CAR_SIGN_USER_TEXT='>'

export GBT_SHELL='bash'
