Go Bullet Train (GBT)
=====================

Highly configurable prompt decoration for ZSH and Bash written in Go. It's
inspired by the [Oh My ZSH](https://github.com/robbyrussell/oh-my-zsh) [Bullet
Train](https://github.com/caiogondim/bullet-train.zsh) theme.

![Screenshot](https://raw.githubusercontent.com/jtyr/gbt/master/images/screenshot01.png "Screenshot")

Works well on Linux (Terminator, Konsole, Gnome Terminal) and Mac (Terminal,
iTerm). It has no other dependencies than Go and its standard libraries.


Table of contents
-----------------

- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
  - [Colors](#colors)
  - [Formatting](#formatting)
  - [Train variables](#train-variables)
  - [Cars variables](#cars-variables)
    - [`Custom` car](#custom-car)
    - [`Dir` car](#dir-car)
    - [`ExecTime` car](#exectime-car)
    - [`Git` car](#git-car)
    - [`Hostname` car](#hostname-car)
    - [`Os` car](#os-car)
    - [`PyVirtEnv` car](#pyvirtenv-car)
    - [`Sign` car](#sign-car)
    - [`Status` car](#status-car)
    - [`Time` car](#time-car)
- [Author](#author)
- [License](#license)


Installation
------------

On Arch Linux:

```shell
yaourt -S gbt
# For ZSH
PROMPT='$(gbt $?)'
# For Bash
PS1='$(gbt $?)'
```

or on Mac via [`brew`](https://brew.sh/):

```shell
brew tap jtyr/repo
brew install gbt
# For ZSH
PROMPT='$(gbt $?)'
# For Bash
PS1='$(gbt $?)'
```

or via Go:

```shell
go get -u github.com/jtyr/gbt
go build -o ~/gbt github.com/jtyr/gbt
# For ZSH
PROMPT='$(~/gbt $?)'
# For Bash
PS1='$(~/gbt $?)'
```

In order to display all colors correctly, the terminal should use 256 color
scheme:

```shell
export TERM="xterm-256color"
```

In order to display all characters of the prompt correctly, the shell should
support UTF-8 and [Nerd](https://github.com/ryanoasis/nerd-fonts) (or at least
[Powerline](https://github.com/ryanoasis/powerline-extra-symbols)) fonts should
be installed and set in the terminal application.


Usage
-----

```shell
### Test the Status car
false
true
### Test the Dir car
cd /
cd /usr/share/doc/sudo
# Display only last 3 elements of the path
export GBT_CAR_DIR_DEPTH="3"
# Display full path
export GBT_CAR_DIR_DEPTH="9999"
# Show only last element of the path
unset GBT_CAR_DIR_DEPTH
cd ~
### Test Time car
# Add the Time car into the train
export GBT_CARS="Status, Os, Time, Hostname, Dir, Sign"
# Set 12h format
export GBT_CAR_TIME_TIME_FORMAT="%I:%M:%S %p"
# Change background color of the all car
export GBT_CAR_TIME_BG="yellow"
# Change color of Date part
export GBT_CAR_TIME_DATE_FG="black"
# Reset the color of the Date part
unset GBT_CAR_TIME_DATE_FG
# Reset the background color of all Time car
unset GBT_CAR_TIME_BG
# Remove the Date part from the car
export GBT_CAR_TIME_FORMAT=" {{ Time }} "
# Reset the format of the car
unset GBT_CAR_TIME_FORMAT
# Reset the original train
unset GBT_CARS
### Themes
# Load theme
source /usr/share/gbt/themes/square_brackets_multiline
```


Configuration
-------------

The prompt (train) is assembled from several elements (cars). The look and
behavior of whole train as well as each car can be influenced by a set of
environment variables. Majority of the 


### Colors

The value of all `_BG` and `_FG` variables defines the background and
foreground color of the particular element. The value of the color can be
specified in 3 ways:

#### Color name

Only a limited number of named colors is supported:

- ![black](https://placehold.it/10/000000/000000?text=+) `black`
- ![red](https://placehold.it/10/aa0000/000000?text=+) `red`
- ![green](https://placehold.it/10/00aa00/000000?text=+) `green`
- ![yellow](https://placehold.it/10/aa5500/000000?text=+) `yellow`
- ![blue](https://placehold.it/10/0000aa/000000?text=+) `blue`
- ![magenta](https://placehold.it/10/aa00aa/000000?text=+) `magenta`
- ![cyan](https://placehold.it/10/00aaaa/000000?text=+) `cyan`
- ![light_gray](https://placehold.it/10/aaaaaa/000000?text=+) `light_gray`
- ![dark_gray](https://placehold.it/10/555555/000000?text=+) `dark_gray`
- ![light_red](https://placehold.it/10/ff5555/000000?text=+) `light_red`
- ![light_green](https://placehold.it/10/55ff55/000000?text=+) `light_green`
- ![light_green](https://placehold.it/10/ffff55/000000?text=+) `light_yellow`
- ![light_blue](https://placehold.it/10/5555ff/000000?text=+) `light_blue`
- ![light_magenta](https://placehold.it/10/ff55ff/000000?text=+) `light_magenta`
- ![light_cyan](https://placehold.it/10/55ffff/000000?text=+) `light_cyan`
- ![white](https://placehold.it/10/ffffff/000000?text=+) `white`
- `default` (default color of the terminal)

Examples:

```shell
# Set the background color of the `Dir` car to red
export GBT_CAR_DIR_BG="red"
# Set the foreground color of the `Dir` car to white
export GBT_CAR_DIR_FG="white"
```

#### Color number

Color can also by expressed by a single number in the range from `0` to
`255`. The color of each number in that range is visible in the 256-color
lookup table on
[Wikipedia](https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit). The named
colors described above are the first 16 numbers from the lookup table.

Examples:

```shell
# Set the background color of the `Dir` car to red
export GBT_CAR_DIR_BG="1"
# Set the foreground color of the `Dir` car to white
export GBT_CAR_DIR_FG="15"
```

#### RGB number

Arbitrary color can be expressed in the form of RGB triplet.

Examples:

```shell
# Set the background color of the `Dir` car to red
export GBT_CAR_DIR_BG="170;0;0"
# Set the foreground color of the `Dir` car to white
export GBT_CAR_DIR_FG="255;255;255"
```


### Formatting

Formatting is done via `_FM` variables. The possible values are:

- `bold`

  Makes the text bold. Not all font characters have variant for bold formatting.

- `underline`

  Makes the text underlined.

- `blink`

  Makes the text to blink.

- `none`

  No formatting applied.

  Example:

  ```shell
  # Set the directory name to be bold
  export GBT_CAR_DIR_FM="bold"
  ```


### Train variables

- `GBT_CARS="Status, Os, Hostname, Dir, Git, Sign"`

  List of cars used in the train.

  To add a new car into the train, the whole variable must be redefined. For
  example in order to add the `Time` car into the default set of cars between
  the `Os` and `Hostname` car, the variable should look like this:

  ```shell
  export GBT_CARS="Status, Os, Time, Hostname, Dir, Git, Sign"
  ```

- `GBT_RCARS="Time"`

  The same like `GBT_CARS` but for the right hand side prompt.

  ```shell
  # Add the Custom car into the right hand site car to have the separator visible
  export GBT_RCARS="Custom, Time"
  # Make the Custom car to be invisible (zero length text)
  export GBT_CAR_CUSTOM_BG="default"
  export GBT_CAR_CUSTOM_FORMAT=""
  # Show only time
  export GBT_CAR_TIME_FORMAT=" {{ Time }} "
  # Set the right hand side prompt (ZSH only)
  export RPROMPT='$(gbt -right)'
  ```

- `GBT_SEPARATOR=""`

  Character used to separate cars in the train.

- `GBT_RSEPARATOR=""`

  The same like `GBT_SEPARATOR` but for the righ hand side prompt.

- `GBT_CAR_BG`

  Background color inherited by the top background color variable of every car.
  That allows to set the background color of all cars via single variable.

- `GBT_CAR_FG`

  Foreground color inherited by the top foreground color variable of every car.
  That allows to set the foreground color of all cars via single variable.

- `GBT_CAR_FM`

  Formatting inherited by the top formatting variable of every car. That allows
  to set the formatting of all cars via single variable.

- `GBT_BEGINNING_BG="default"`

  Background color of the text shown at the beginning of the train.

- `GBT_BEGINNING_FG="default"`

  Foreground color of the text shown at the beginning of the train.

- `GBT_BEGINNING_FM="none"`

  Formatting of the text shown at the beginning of the train.

- `GBT_BEGINNING_TEXT=""`

  Text shown at the beginning of the train.

- `GBT_SHELL`

  Indicates which shell is used. The value can be either `zsh` or `bash`. By
  default, the value is extracted from the `$SHELL` environment variable. Set
  this variable to `bash` if your default shell is ZSH but you want to test GBT
  in Bash:

  ```shell
  export GBT_SHELL="bash"
  bash
  ```

- `GBT_DEBUG="0"`

  Shows more verbose output if some of the car modules cannot be imported.


### Cars variables

#### `Custom` car

The main purpose of this car is to provide the possibility to create car with
custom text.

- `GBT_CAR_CUSTOM_BG="yellow"`

  Background color of the car.

- `GBT_CAR_CUSTOM_FG="default"`

  Foreground color of the car.

- `GBT_CAR_CUSTOM_FM="none"`

  Formatting of the car.

- `GBT_CAR_CUSTOM_FORMAT=" {{ Text }} "`

  Format of the car.

- `GBT_CAR_CUSTOM_TEXT_BG`

  Background color of the `{{ Text }}` element.

- `GBT_CAR_CUSTOM_TEXT_FG`

  Foreground color of the `{{ Text }}` element.

- `GBT_CAR_CUSTOM_TEXT_FM`

  Formatting of the `{{ Text }}` element.

- `GBT_CAR_CUSTOM_TEXT_TEXT="?"`

  Text content of the `{{ Text }}` element.

- `GBT_CAR_CUSTOM_TEXT_CMD`

  The `{{ Text }}` element will be replaced by standard output of the command
  specified in this variable. Content of the `GBT_CAR_CUSTOM_TEXT_TEXT` variable
  takes precedence over this variable.

  ```shell
  # Show 1 minute loadavg as the content of the Text element
  export GBT_CAR_CUSTOM_CMD="uptime | sed --e 's/.*load average: //' -e 's/,.*//'"
  ```

- `GBT_CAR_CUSTOM_DISPLAY="1"`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_CUSTOM_WRAP="0"`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_CUSTOM_SEP`

  Custom separator string for this car.

Multiple `Custom` cars can be used in the `GBT_CARS` variable. Just add some
identifier behind the car name. To set properties of the new car, just add the
same identifier into the environment variable:

```shell
# Adding Custom and Custo1 car
export GBT_CARS="Status, Os, Custom, Custom1, Hostname, Dir, Git, Sign"
# The text of the default Custom car
export GBT_CAR_CUSTOM_TEXT_TEXT="default"
# The text of the Custom1 car
export GBT_CAR_CUSTOM1_TEXT_TEXT="1"
# Set different background color for the Custom1 car
export GBT_CAR_CUSTOM1_BG="magenta"
```


#### `Dir` car

Car that displays current directory name.

- `GBT_CAR_DIR_BG="blue"`

  Background color of the car.

- `GBT_CAR_DIR_FG="light_gray"`

  Foreground color of the car.

- `GBT_CAR_DIR_FM="none"`

  Formatting of the car.

- `GBT_CAR_DIR_FORMAT=" {{ Dir }} "`

  Format of the car.

- `GBT_CAR_DIR_DIR_BG`

  Background color of the `{{ Dir }}` element.

- `GBT_CAR_DIR_DIR_FG`

  Foreground color of the `{{ Dir }}` element.

- `GBT_CAR_DIR_DIR_FM`

  Formatting of the `{{ Dir }}` element.

- `GBT_CAR_DIR_DIR_TEXT`

  Text content of the `{{ Dir }}` element. The directory name.

- `GBT_CAR_DIR_DIRSEP`

  OS-default character used to separate directories.

- `GBT_CAR_DIR_HOMESIGN="~"`

  Character which represents the user's home directory. If set to empty
  string, full home directory path is used instead.

- `GBT_CAR_DIR_DEPTH="1"`

  Number of directories to show.

- `GBT_CAR_DIR_DISPLAY="1"`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_DIR_WRAP="0"`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_DIR_SEP`

  Custom separator string for this car.


#### `ExecTime` car

Car that displays how long each shell command run.

- `GBT_CAR_EXECTIME_BG="light_gray"`

  Background color of the car.

- `GBT_CAR_EXECTIME_FG="black"`

  Foreground color of the car.

- `GBT_CAR_EXECTIME_FM="none"`

  Formatting of the car.

- `GBT_CAR_EXECTIME_FORMAT=" {{ Time }} "`

  Format of the car.

- `GBT_CAR_EXECTIME_TIME_BG`

  Background color of the `{{ Time }}` element.

- `GBT_CAR_EXECTIME_TIME_FG`

  Foreground color of the `{{ Time }}` element.

- `GBT_CAR_EXECTIME_TIME_FM`

  Formatting of the `{{ Time }}` element.

- `GBT_CAR_EXECTIME_TIME_TEXT`

  Text content of the `{{ Time }}` element. The execution time.

- `GBT_CAR_EXECTIME_DIRSEP`

  OS-default character used to separate directories.

- `GBT_CAR_EXECTIME_PRECISION="0"`

  Sub-second precision to show.

- `GBT_CAR_EXECTIME_SECS`

  The number of seconds the command run in shell. This variable is defined in
  the source file as shown bellow.

- `GBT_CAR_EXECTIME_BELL="0"`

  Sound console bell if the executed command exceeds specified number of
  seconds. Value set to `0` disables the bell (default).

- `GBT_CAR_EXECTIME_DISPLAY="1"`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_EXECTIME_WRAP="0"`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_EXECTIME_SEP`

  Custom separator string for this car.

In order to allow this car to calculate the execution time, the following must
be loaded in the shell:

```shell
# For ZSH
source /usr/share/gbt/sources/ExecTime.zsh
# For Bash
source /usr/share/gbt/sources/ExecTime.bash
```


#### `Git` car

Car that displays information about a local Git repository. This car is
displayed only if the current directory is a Git repository.

- `GBT_CAR_GIT_BG="light_gray"`

  Background color of the car.

- `GBT_CAR_GIT_FG="black"`

  Foreground color of the car.

- `GBT_CAR_GIT_FM="none"`

  Formatting of the car.

- `GBT_CAR_GIT_FORMAT=" {{ Icon }} {{ Head }} {{ Status }}{{ Ahead }}{{ Behind }} "`

  Format of the car.

- `GBT_CAR_GIT_ICON_BG`

  Background color of the `{{ Icon }}` element.

- `GBT_CAR_GIT_ICON_FG`

  Foreground color of the `{{ Icon }}` element.

- `GBT_CAR_GIT_ICON_FM`

  Formatting of the `{{ Icon }}` element.

- `GBT_CAR_GIT_ICON_TEXT=""`

  Text content of the `{{ Icon }}` element.

- `GBT_CAR_GIT_HEAD_BG`

  Background color of the `{{ Head }}` element.

- `GBT_CAR_GIT_HEAD_FG`

  Foreground color of the `{{ Head }}` element.

- `GBT_CAR_GIT_HEAD_FM`

  Formatting of the `{{ Head }}` element.

- `GBT_CAR_GIT_HEAD_TEXT`

  Text content of the `{{ Head }}` element. The branch or tag name or the
  commit ID.

- `GBT_CAR_GIT_STATUS_BG`

  Background color of the `{{ Status }}` element.

- `GBT_CAR_GIT_STATUS_FG`

  Foreground color of the `{{ Status }}` element.

- `GBT_CAR_GIT_STATUS_FM`

  Formatting of the `{{ Status }}` element.

- `GBT_CAR_GIT_STATUS_FORMAT`

  Format of the `{{ Status }}` element. The content is either `{{ Dirty }}` or
  `{{ Clean }}` depending on the state of the local Git repository.

- `GBT_CAR_GIT_DIRTY_BG`

  Background color of the `{{ Dirty }}` element.

- `GBT_CAR_GIT_DIRTY_FG="red"`

  Foreground color of the `{{ Dirty }}` element.

- `GBT_CAR_GIT_DIRTY_FM`

  Formatting of the `{{ Dirty }}` element.

- `GBT_CAR_GIT_DIRTY_TEXT="✘"`

  Text content of the `{{ Dirty }}` element.

- `GBT_CAR_GIT_CLEAN_BG`

  Background color of the `{{ Clean }}` element.

- `GBT_CAR_GIT_CLEAN_FG="green"`

  Foreground color of the `{{ Clean }}` element.

- `GBT_CAR_GIT_CLEAN_FM`

  Formatting of the `{{ Clean }}` element.

- `GBT_CAR_GIT_CLEAN_TEXT="✔"`

  Text content of the `{{ Clean }}` element.

- `GBT_CAR_GIT_AHEAD_BG`

  Background color of the `{{ Ahead }}` element.

- `GBT_CAR_GIT_AHEAD_FG`

  Foreground color of the `{{ Ahead }}` element.

- `GBT_CAR_GIT_AHEAD_FM`

  Formatting of the `{{ Ahead }}` element.

- `GBT_CAR_GIT_AHEAD_TEXT=" ⬆"`

  Text content of the `{{ Ahead }}` element.

- `GBT_CAR_GIT_BEHIND_BG`

  Background color of the `{{ Behind }}` element.

- `GBT_CAR_GIT_BEHIND_FG`

  Foreground color of the `{{ Behind }}` element.

- `GBT_CAR_GIT_BEHIND_FM`

  Formatting of the `{{ Behind }}` element.

- `GBT_CAR_GIT_BEHIND_TEXT=" ⬇"`

  Text content of the `{{ Behind }}` element.

- `GBT_CAR_GIT_DISPLAY`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_GIT_WRAP="0"`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_GIT_SEP`

  Custom separator string for this car.


#### `Hostname` car

Car that displays username of the currently logged user and the hostname of the
local machine.

- `GBT_CAR_HOSTNAME_BG="dark_gray"`

  Background color of the car.

- `GBT_CAR_HOSTNAME_FG="252"`

  Foreground color of the car.

- `GBT_CAR_HOSTNAME_FM="none"`

  Formatting of the car.

- `GBT_CAR_HOSTNAME_FORMAT=" {{ UserHost }} "`

  Format of the car.

- `GBT_CAR_HOSTNAME_USERHOST_BG`

  Background color of the `{{ UserHost }}` element.

- `GBT_CAR_HOSTNAME_USERHOST_FG`

  Foreground color of the `{{ UserHost }}` element.

- `GBT_CAR_HOSTNAME_USERHOST_FM`

  Formatting of the `{{ UserHost }}` element.

- `GBT_CAR_HOSTNAME_USERHOST_FORMAT="{{ User }}@{{ Host }}"`

  Format of the `{{ UserHost }}` element.

- `GBT_CAR_HOSTNAME_USER_BG`

  Background color of the `{{ User }}` element.

- `GBT_CAR_HOSTNAME_USER_FG`

  Foreground color of the `{{ User }}` element.

- `GBT_CAR_HOSTNAME_USER_FM`

  Formatting of the `{{ User }}` element.

- `GBT_CAR_HOSTNAME_USER_TEXT`

  Text content of the `{{ User }}` element. The user name.

- `GBT_CAR_HOSTNAME_HOST_BG`

  Background color of the `{{ Host }}` element.

- `GBT_CAR_HOSTNAME_HOST_FG`

  Foreground color of the `{{ Host }}` element.

- `GBT_CAR_HOSTNAME_HOST_FM`

  Formatting of the `{{ Host }}` element.

- `GBT_CAR_HOSTNAME_HOST_TEXT`

  Text content of the `{{ Host }}` element. The host name.

- `GBT_CAR_HOSTNAME_DISPLAY="1"`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_HOSTNAME_WRAP="0"`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_HOSTNAME_SEP`

  Custom separator string for this car.


#### `Os` car

Car that displays icon of the operating system.

- `GBT_CAR_OS_BG="235"`

  Background color of the car.

- `GBT_CAR_OS_FG="white"`

  Foreground color of the car.

- `GBT_CAR_OS_FM="none"`

  Formatting of the car.

- `GBT_CAR_OS_FORMAT=" {{ Symbol }} "`

  Format of the car.

- `GBT_CAR_OS_SYMBOL_BG`

  Background color of the `{{ Symbol }}` element.

- `GBT_CAR_OS_SYMBOL_FG`

  Foreground color of the `{{ Symbol }}` element.

- `GBT_CAR_OS_SYMBOL_FM`

  Formatting of the `{{ Symbol }}` element.

- `GBT_CAR_OS_SYMBOL_TEXT`

  Text content of the `{{ Symbol }}` element.

- `GBT_CAR_OS_NAME`

  The name of the symbol to display. Default value is selected by the system
  the shell runs at. Possible names and their symbols are:

  - `amzn` 
  - `android` 
  - `arch` 
  - `archarm` 
  - `centos` 
  - `cloud` 
  - `coreos` 
  - `darwin` 
  - `debian` 
  - `docker` 
  - `elementary` 
  - `fedora` 
  - `freebsd` 
  - `gentoo` 
  - `linux` 
  - `linuxmint` 
  - `mageia` 
  - `mandriva` 
  - `opensuse` 
  - `raspbian` 
  - `redhat` 
  - `sabayon` 
  - `slackware` 
  - `ubuntu` 
  - `windows` 

  Example:

  ```shell
  export GBT_CAR_OS_NAME="arch"
  ```

- `GBT_CAR_OS_DISPLAY="1"`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_OS_WRAP="0"`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_OS_SEP`

  Custom separator string for this car.


#### `PyVirtEnv` car

Car that displays Python Virtual Environment name. This car is displayed only
if the Python Virtual Environment is activated. The activation script usually
prepends the shell prompt by the Virtual Environment name by default. In order
to disable it, the following environment variable must be set:

```shell
export VIRTUAL_ENV_DISABLE_PROMPT="1"
```

Variables used by the car:

- `GBT_CAR_PYVIRTENV_BG="222"`

  Background color of the car.

- `GBT_CAR_PYVIRTENV_FG="black"`

  Foreground color of the car.

- `GBT_CAR_PYVIRTENV_FM="none"`

  Formatting of the car.

- `GBT_CAR_PYVIRTENV_FORMAT=" {{ Icon }} {{ Name }} "`

  Format of the car.

- `GBT_CAR_PYVIRTENV_ICON_BG`

  Background color of the `{{ Icon }}` element.

- `GBT_CAR_PYVIRTENV_ICON_FG`

  Foreground color of the `{{ Icon }}` element.

- `GBT_CAR_PYVIRTENV_ICON_FM`

  Formatting of the `{{ Icon }}` element.

- `GBT_CAR_PYVIRTENV_ICON_TEXT`

  Text content of the `{{ Icon }}` element.

- `GBT_CAR_PYVIRTENV_NAME_BG`

  Background color of the `{{ Name }}` element.

- `GBT_CAR_PYVIRTENV_NAME_FG="33"`

  Foreground color of the `{{ NAME }}` element.

- `GBT_CAR_PYVIRTENV_NAME_FM`

  Formatting of the `{{ Name }}` element.

- `GBT_CAR_PYVIRTENV_NAME_TEXT`

  The name of the Python Virtual Environment deducted from the `VIRTUAL_ENV`
  environment variable.

- `GBT_CAR_PYVIRTENV_DISPLAY`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_PYVIRTENV_WRAP="0"`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_PYVIRTENV_SEP`

  Custom separator string for this car.


#### `Sign` car

Car that displays prompt character for the admin and user at the end of the
train.

- `GBT_CAR_SIGN_BG="default"`

  Background color of the car.

- `GBT_CAR_SIGN_FG="default"`

  Foreground color of the car.

- `GBT_CAR_SIGN_FM="none"`

  Formatting of the car.

- `GBT_CAR_SIGN_FORMAT=" {{ Symbol }} "`

  Format of the car.

- `GBT_CAR_SIGN_SYMBOL_BG`

  Background color of the `{{ Symbol }}` element.

- `GBT_CAR_SIGN_SYMBOL_FG`

  Foreground color of the `{{ Symbol }}` element.

- `GBT_CAR_SIGN_SYMBOL_FM="bold"`

  Formatting of the `{{ Symbol }}` element.

- `GBT_CAR_SIGN_SYMBOL_FORMAT`

  Format of the `{{ Symbol }}` element. The format is either `{{ Admin }}` if
  the UID is 0 or `{{ User }}` if the UID is not 0.

- `GBT_CAR_SIGN_ADMIN_BG`

  Background color of the `{{ Admin }}` element.

- `GBT_CAR_SIGN_ADMIN_FG="red"`

  Foreground color of the `{{ Admin }}` element.

- `GBT_CAR_SIGN_ADMIN_FM`

  Formatting of the `{{ Admin }}` element.

- `GBT_CAR_SIGN_ADMIN_TEXT="#"`

  Text content of the `{{ Admin }}` element.

- `GBT_CAR_SIGN_USER_BG`

  Background color of the `{{ User }}` element.

- `GBT_CAR_SIGN_USER_FG="light_green"`

  Foreground color of the `{{ User }}` element.

- `GBT_CAR_SIGN_USER_FM`

  Formatting of the `{{ User }}` element.

- `GBT_CAR_SIGN_USER_TEXT="$"`

  Text content of the `{{ User }}` element. The user name.

- `GBT_CAR_SIGN_DISPLAY="1"`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_SIGN_WRAP="0"`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_SIGN_SEP`

  Custom separator string for this car.


#### `Status` car

Car that visualizes return code of every command. By default, this car is
displayed only when the return code is non-zero. If you want to display it even
when if the return code is zere, set the following variable:

```shell
export GBT_CAR_STATUS_DISPLAY="1"
```

Variables used by the car:

- `GBT_CAR_STATUS_BG`

  Background color of the car. It's either `GBT_CAR_STATUS_OK_BG` if the last
  command returned `0` return code otherwise the `GBT_CAR_STATUS_ERROR_BG` is
  used.

- `GBT_CAR_STATUS_FG="default"`

  Foreground color of the car. It's either `GBT_CAR_STATUS_OK_FG` if the last
  command returned `0` return code otherwise the `GBT_CAR_STATUS_ERROR_FG` is
  used.

- `GBT_CAR_STATUS_FM="none"`

  Formatting of the car. It's either `GBT_CAR_STATUS_OK_FM` if the last command
  returned `0` return code otherwise the `GBT_CAR_STATUS_ERROR_FM` is used.

- `GBT_CAR_STATUS_FORMAT=" {{ Symbol }} "`

  Format of the car. This can be changed to contain also the value of the
  return code:

  ```shell
  export GBT_CAR_STATUS_FORMAT=" {{ Symbol }} {{ Code }} "
  ```

- `GBT_CAR_STATUS_SYMBOL_BG`

  Background color of the `{{ Symbol }}` element.

- `GBT_CAR_STATUS_SYMBOL_FG`

  Foreground color of the `{{ Symbol }}` element.

- `GBT_CAR_STATUS_SYMBOL_FM="bold"`

  Formatting of the `{{ Symbol }}` element.

- `GBT_CAR_STATUS_SYMBOL_FORMAT`

  Format of the `{{ Symbol }}` element. The format is either `{{ Error }}` if
  the last command returned non zero return code otherwise `{{ User }}` is
  used.

- `GBT_CAR_STATUS_CODE_BG="red"`

  Background color of the `{{ Code }}` element.

- `GBT_CAR_STATUS_CODE_FG="light_gray"`

  Foreground color of the `{{ Code }}` element.

- `GBT_CAR_STATUS_CODE_FM="none"`

  Formatting of the `{{ Code }}` element.

- `GBT_CAR_STATUS_CODE_TEXT`

  Text content of the `{{ Code }}` element. The return code.

- `GBT_CAR_STATUS_ERROR_BG="red"`

  Background color of the `{{ Error }}` element.

- `GBT_CAR_STATUS_ERROR_FG="light_gray"`

  Foreground color of the `{{ Error }}` element.

- `GBT_CAR_STATUS_ERROR_FM="none"`

  Formatting of the `{{ Error }}` element.

- `GBT_CAR_STATUS_ERROR_TEXT="✘"`

  Text content of the `{{ Error }}` element.

- `GBT_CAR_STATUS_OK_BG="green"`

  Background color of the `{{ Ok }}` element.

- `GBT_CAR_STATUS_OK_FG="light_gray"`

  Foreground color of the `{{ Ok }}` element.

- `GBT_CAR_STATUS_OK_FM="none"`

  Formatting of the `{{ Ok }}` element.

- `GBT_CAR_STATUS_OK_TEXT="✔"`

  Text content of the `{{ Ok }}` element.

- `GBT_CAR_STATUS_DISPLAY`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_STATUS_WRAP="0"`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_STATUS_SEP`

  Custom separator string for this car.


#### `Time` car

Car that displays current date and time.

- `GBT_CAR_TIME_BG="light_blue"`

  Background color of the car.

- `GBT_CAR_TIME_FG="light_gray"`

  Foreground color of the car.

- `GBT_CAR_TIME_FM="none"`

  Formatting of the car.

- `GBT_CAR_TIME_FORMAT=" {{ DateTime }} "`

  Format of the car.

- `GBT_CAR_TIME_DATETIME_BG`

  Background color of the `{{ DateTime }}` element.

- `GBT_CAR_TIME_DATETIME_FG`

  Foreground color of the `{{ DateTime }}` element.

- `GBT_CAR_TIME_DATETIME_FM`

  Formatting of the `{{ DateTime }}` element.

- `GBT_CAR_TIME_DATETIME_FORMAT="{{ Date }} {{ Time }}"`

  Format of the `{{ DateTime }}` element.

- `GBT_CAR_TIME_DATE_BG`

  Background color of the `{{ Date }}` element.

- `GBT_CAR_TIME_DATE_FG`

  Foreground color of the `{{ Date }}` element.

- `GBT_CAR_TIME_DATE_FM`

  Formatting of the `{{ Date }}` element.

- `GBT_CAR_TIME_DATE_FORMAT="%a %d %b"`

  Format of the `{{ Date }}` element. The format is using sequences known from
  the `date` command.

- `GBT_CAR_TIME_TIME_BG`

  Background color of the `{{ Host }}` element.

- `GBT_CAR_TIME_TIME_FG="light_yellow"`

  Foreground color of the `{{ Host }}` element.

- `GBT_CAR_TIME_TIME_FM`

  Formatting of the `{{ Host }}` element.

- `GBT_CAR_TIME_TIME_FORMAT="%H:%M:%S"`

  Text content of the `{{ Host }}` element. The format is using sequences known
  from the `date` command.

- `GBT_CAR_TIME_DISPLAY="1"`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_TIME_WRAP="0"`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_TIME_SEP`

  Custom separator string for this car.


Author
------

Jiri Tyr


License
-------

MIT
