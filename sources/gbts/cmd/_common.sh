# Customize 'which' option for ZSH
[ "$(ps -p $$ 2>/dev/null | awk '$1 != "PID" {print $(NF)}' | sed 's/-//g')" = 'zsh' ] && GBT__WHICH_OPTS='-p'

[ -z "$GBT__SOURCE_COMPRESS" ] && GBT__SOURCE_COMPRESS='gzip -qc9'
[ -z "$GBT__SOURCE_DECOMPRESS" ] && GBT__SOURCE_DECOMPRESS='gzip -qd'
[ -z "$GBT__SOURCE_BASE64" ] && GBT__SOURCE_BASE64='base64'
[ -z "$GBT__SOURCE_BASE64_DEC" ] && GBT__SOURCE_BASE64_DEC='-d'
[ -z "$GBT__SOURCE_SEC_CUT_LOCAL" ] && GBT__SOURCE_SEC_CUT_LOCAL='1'
[ -z "$GBT__SOURCE_SEC_SUM_LOCAL" ] && GBT__SOURCE_SEC_SUM_LOCAL='md5sum'
[ -z "$GBT__SOURCE_SEC_CUT_REMOTE" ] && GBT__SOURCE_SEC_CUT_REMOTE='1'
[ -z "$GBT__SOURCE_SEC_SUM_REMOTE" ] && GBT__SOURCE_SEC_SUM_REMOTE='md5sum'


function gbt__err() {
    echo "$@" >&2
}
