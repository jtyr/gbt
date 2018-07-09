Use the following command to generate the OS symbols for the `remote` script:

```shell
for N in $(grep 'nf-' pkg/cars/os/main.go | sed -r -e 's/",\s+}.*//' -e 's/",.*"/,/' -e 's/":.*"/:/' -e 's/.*"//'); do NAME=${N%%:*}; COLOR=${N##*,}; TMP=${N##*:}; ICON=${TMP%%,*}; echo "GBT__SYMBOLS[$NAME]='\\\\001\\\\e[38;5;${COLOR}m\\\\002$(echo -ne $ICON | xxd -plain | sed 's/\(..\)/\\\\x\1/g')'"; done
```

Use the following command to convert Unicode character to code:

```shell
printf '\\u%02x\n' "'"
```
