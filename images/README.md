First we need to get the window ID:

```shell
WIN_ID=$(printf '%d' $(xwininfo | grep -Po '(?<=xwininfo: Window id: )(0x[a-f0-9]{7})'))
```

Then we can type the file `demo.input` into the selected window:

```shell
xdotool windowactivate $WIN_ID type --file demo.input --delay 100
```

Set terminal window to display 118x14 characters, sect curosr color to black
and disable KDE effects. Then record the demo with
[Peek](https://github.com/phw/peek).
