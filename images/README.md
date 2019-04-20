Set terminal window to the size of 104x14 characters (834x240px). Disable KDE
effects and make sure [Peek](https://github.com/phw/peek) is installed and
configured to start recording after 1 second by pressing `CTRL+ALT+E` and using
20 FPS.

Then get the window ID where the typing will happen:

```shell
WIN_ID=$(printf '%d' $(xwininfo | grep -Po '(?<=xwininfo: Window id: )(0x[a-f0-9]+)'))
```

Automatically start recording and typing and then stop recording:


```shell
xdotool \
  windowactivate $WIN_ID \
  key ctrl+alt+e \
  sleep 2 \
  type --file demo.input --delay 100 && \
xdotool \
  sleep 2 \
  key ctrl+alt+e
```

Optimize final GIF by using [ezgif.com](https://ezgif.com/optimize),
compression level 35.
