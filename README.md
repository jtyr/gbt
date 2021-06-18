Go Bullet Train (GBT)
=====================

Highly configurable prompt builder for Bash, ZSH and PowerShell written in Go.
It's inspired by the [Oh My ZSH](https://github.com/robbyrussell/oh-my-zsh)
[Bullet Train](https://github.com/caiogondim/bullet-train.zsh) theme but runs
significantly faster.

![Demo](https://raw.githubusercontent.com/jtyr/gbt/master/images/demo.gif "Demo")

GBT comes with an interesting feature called
[prompt forwarding](#prompt-forwarding) which allows to forward user-defined
prompt to a remote machine and have the same-looking prompt across all machines
via SSH but also in Docker, Kubectl, Vagrant, MySQL or in Screen without the
need to install anything remotely.

![Prompt forwarding demo](https://raw.githubusercontent.com/jtyr/gbt/master/images/prompt_forwarding.gif "Prompt forwarding demo")

All the above works well on Linux (Terminator, Konsole, Gnome Terminal), Mac
(Terminal, iTerm), Android (Termux) and Windows (PowerShell, Windows Terminal).

[![Release](https://img.shields.io/github/release/jtyr/gbt.svg)](https://github.com/jtyr/gbt/releases)
[![Build status](https://travis-ci.org/jtyr/gbt.svg?branch=master)](https://travis-ci.org/jtyr/gbt)
[![Coverage Status](https://coveralls.io/repos/github/jtyr/gbt/badge.svg?branch=master)](https://coveralls.io/github/jtyr/gbt?branch=master)
[![Packagecloud](https://img.shields.io/badge/%E2%98%81-Packagecloud-707aed.svg)](https://packagecloud.io/gbt/release)


Table of contents
-----------------

- [Setup](#setup)
  - [Installation](#installation)
    - [Arch Linux](#arch-linux)
    - [CentOS/RHEL](#centosrhel)
    - [Ubuntu/Debian](#ubuntudebian)
    - [Mac](#mac)
    - [Windows](#windows)
    - [Android](#android)
    - [From the source code](#from-the-source-code)
  - [Activation](#activation)
  - [Fonts and colors](#fonts-and-colors)
- [Configuration](#configuration)
  - [Colors](#colors)
  - [Formatting](#formatting)
  - [Train variables](#train-variables)
  - [Cars variables](#cars-variables)
    - [`Aws` car](#aws-car)
    - [`Azure` car](#azure-car)
    - [`Custom` car](#custom-car)
    - [`Dir` car](#dir-car)
    - [`ExecTime` car](#exectime-car)
    - [`Gcp` car](#gcp-car)
    - [`Git` car](#git-car)
    - [`Hostname` car](#hostname-car)
    - [`Kubectl` car](#kubectl-car)
    - [`Os` car](#os-car)
    - [`PyVirtEnv` car](#pyvirtenv-car)
    - [`Sign` car](#sign-car)
    - [`Status` car](#status-car)
    - [`Time` car](#time-car)
- [Benchmark](#benchmark)
- [Prompt forwarding](#prompt-forwarding)
  - [Principle](#principle)
  - [Additional settings](#additional-settings)
  - [MacOS users](#macos-users)
  - [Limitations](#limitations)
- [TODO](#todo)
- [Author](#author)
- [License](#license)


Setup
-----

In order to setup GBT on your machine, you have to [install](#installation) it,
[activate](#activation) it and setup a special [font](#fonts-and-colors) in your
terminal (optional).

### Installation

#### Arch Linux

```shell
yaourt -S gbt
```

Or install `gbt-git` if you would like to run the latest greatest from the
`master` branch.

#### CentOS/RHEL

Packages hosted by [Packagecloud](https://packagecloud.io/gbt/release):

```shell
echo '[gbt]
name=GBT YUM repo
baseurl=https://packagecloud.io/gbt/release/el/7/$basearch
gpgkey=https://packagecloud.io/gbt/release/gpgkey
       https://packagecloud.io/gbt/release/gpgkey/gbt-release-4C6E79EFF45439B6.pub.gpg
gpgcheck=1
repo_gpgcheck=1' | sudo tee /etc/yum.repos.d/gbt.repo >/dev/null
sudo yum install gbt
```

Use the exact repository definition from above for all RedHat-based
distribution regardless its version.

#### Ubuntu/Debian

Packages hosted by [Packagecloud](https://packagecloud.io/gbt/release):

```shell
curl -L https://packagecloud.io/gbt/release/gpgkey | sudo apt-key add -
echo 'deb https://packagecloud.io/gbt/release/ubuntu/ xenial main' | sudo tee /etc/apt/sources.list.d/gbt.list >/dev/null
sudo apt-get update
sudo apt-get install gbt
```

Use the exact repository definition from above for all Debian-based
distribution regardless its version.

#### Mac

Using [`brew`](https://brew.sh):

```shell
brew tap jtyr/repo
brew install gbt
```
Or install `gbt-git` if you would like to run the latest greatest from the
`master` branch:

```shell
brew tap jtyr/repo
brew install --HEAD gbt-git
```

#### Windows

Using [`choco`](https://chocolatey.org):

```powershell
choco install gbt
```

Using [`scoop`](https://scoop.sh):

```powershell
scoop install gbt
```

Or manually by copying the `gbt.exe` file into a directory listed in the `PATH`
environment variable (e.g. `C:\Windows\system32`).

#### Android

Install [Termux](https://termux.com) from [Google Play Store](https://play.google.com/store/apps/details?id=com.termux)
and then type this in the Termux app:

```shell
apt update
apt install gbt
```

#### From the source code

Make sure [Go](https://golang.org) is installed and then run the following on
Linux and Mac:

```shell
mkdir ~/go
export GOPATH=~/go
export PATH="$PATH:$GOPATH/bin"
go get github.com/jtyr/gbt/cmd/gbt
```

Or the following on Windows using PowerShell:

```powershell
mkdir ~/go
$Env:GOPATH = '~/go'
$Env:PATH = "~/go/bin;$Env:PATH"
go get github.com/jtyr/gbt/cmd/gbt
```

---

### Activation

After GBT is installed, it can be activated by calling it from the shell prompt
variable:

```shell
# For Bash
PS1='$(gbt $?)'

# For ZSH
PROMPT='$(gbt $?)'
```

If you are using ZSH together with some shell framework (e.g. [Oh My
ZSH](https://github.com/robbyrussell/oh-my-zsh)), your shell is processing a
fair amount of shell scripts upon ever prompt appearance. You can speed up your
shell by removing the framework dependency from your configuration and replacing
it with GBT and a [simple ZSH
configuration](https://gist.github.com/jtyr/be0e6007bd22c9d51e8702a70430d116#file-zshrc-L1-L43).
Combining pure ZSH configuration with GBT will provide the best possible
performance for your shell.

To activate GBT in PowerShell, run the following in the console or store it to
the PowerShell profile file (`echo $profile`):

```powershell
function prompt {
    $rc = [int]$(-Not $?)
    $Env:GBT_SHELL = 'plain'
    $Env:PWD = get-location
    $Env:GBT_CAR_CUSTOM_EXECUTOR='powershell.exe'
    $Env:GBT_CAR_CUSTOM_EXECUTOR_PARAM='-Command'
    $gbt_output = & @({gbt $rc},{gbt.exe $rc})[$PSVersionTable.PSVersion.Major -lt 6 -or $IsWindows] | Out-String
    $gbt_output = $gbt_output -replace ([Environment]::NewLine + '$'), ''
    Write-Host -NoNewline $gbt_output
    return [char]0
}
# Needed only on Windows
[console]::InputEncoding = [console]::OutputEncoding = New-Object System.Text.UTF8Encoding
```

---

### Fonts and colors

Although GBT can be configured to use only ASCII characters (see
[`basic`](blob/master/themes/basic.sh) theme), the default configuration uses
some UTF-8 characters which require special font. In order to display all
characters of the default prompt correctly, the shell should support UTF-8 and
[Nerd](https://github.com/ryanoasis/nerd-fonts) fonts (or at least the
[DejaVuSansMono
Nerd](https://github.com/ryanoasis/nerd-fonts/tree/master/patched-fonts/DejaVuSansMono/Regular/complete)
font) should be installed. On Linux, you can install it like this:

```shell
mkdir ~/.fonts
curl -L -o ~/.fonts/DejaVuSansMonoNerdFontCompleteMono.ttf https://github.com/ryanoasis/nerd-fonts/raw/master/patched-fonts/DejaVuSansMono/Regular/complete/DejaVu%20Sans%20Mono%20Nerd%20Font%20Complete%20Mono.ttf
fc-cache
```

On Mac, it can be installed via `brew`:

```shell
brew tap homebrew/cask-fonts
brew install --cask font-dejavu-sans-mono-nerd-font
```

On Windows, it can be installed via `choco`:

```powershell
choco install font-nerd-DejaVuSansMono
```

Or via `scoop`:

```powershell
scoop bucket add nerd-fonts
scoop install DejaVuSansMono-NF
```

Or just [download](https://github.com/ryanoasis/nerd-fonts/raw/master/patched-fonts/DejaVuSansMono/Regular/complete/DejaVu%20Sans%20Mono%20Nerd%20Font%20Complete%20Mono%20Windows%20Compatible.ttf)
the font, open it and then install it.

Once the font is installed, it has to be set in the terminal application to
render all prompt characters correctly. Search for the font name `DejaVuSansMono
Nerd Font Mono` on Linux and Mac and `DejaVuSansMono NF` on Windows.

In order to have the Nerd font in Termux on Android, you have to install
[Termux:Styling](https://play.google.com/store/apps/details?id=com.termux.styling)
application. Then longpress the terminal screen and select `MORE...` → `Style`
→ `CHOOSE FONT` and there choose the `DejaVu` font.

Some Unix terminals might not use 256 color palette by default. In such case try
to set the following:

```shell
export TERM='xterm-256color'
```


Configuration
-------------

The prompt (train) is assembled from several elements (cars). The look and
behavior of whole train as well as each car can be influenced by a set of
environment variables. To set the environment variable, use `export` in the
Linux and Mac shell and `$Env:` on Windows.


### Colors

The value of all `_BG` and `_FG` variables defines the background and
foreground color of the particular element. The value of the color can be
specified in 3 ways:

#### Color name

Only a limited number of named colors is supported:

- ![black](https://via.placeholder.com/10/000000/000000?text=+) `black`
- ![red](https://via.placeholder.com/10/800000/000000?text=+) `red`
- ![green](https://via.placeholder.com/10/008000/000000?text=+) `green`
- ![yellow](https://via.placeholder.com/10/808000/000000?text=+) `yellow`
- ![blue](https://via.placeholder.com/10/000080/000000?text=+) `blue`
- ![magenta](https://via.placeholder.com/10/800080/000000?text=+) `magenta`
- ![cyan](https://via.placeholder.com/10/008080/000000?text=+) `cyan`
- ![light_gray](https://via.placeholder.com/10/c0c0c0/000000?text=+) `light_gray`
- ![dark_gray](https://via.placeholder.com/10/808080/000000?text=+) `dark_gray`
- ![light_red](https://via.placeholder.com/10/ff0000/000000?text=+) `light_red`
- ![light_green](https://via.placeholder.com/10/00ff00/000000?text=+) `light_green`
- ![light_green](https://via.placeholder.com/10/ffff00/000000?text=+) `light_yellow`
- ![light_blue](https://via.placeholder.com/10/0000ff/000000?text=+) `light_blue`
- ![light_magenta](https://via.placeholder.com/10/ff00ff/000000?text=+) `light_magenta`
- ![light_cyan](https://via.placeholder.com/10/00ffff/000000?text=+) `light_cyan`
- ![white](https://via.placeholder.com/10/ffffff/000000?text=+) `white`
- `default` (default color of the terminal)

Examples:

```shell
# Set the background color of the `Dir` car to red
export GBT_CAR_DIR_BG='red'
# Set the foreground color of the `Dir` car to white
export GBT_CAR_DIR_FG='white'
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
export GBT_CAR_DIR_BG='1'
# Set the foreground color of the `Dir` car to white
export GBT_CAR_DIR_FG='15'
```

#### RGB color

Arbitrary color can be expressed in the form of RGB triplet.

Examples:

```shell
# Set the background color of the `Dir` car to red
export GBT_CAR_DIR_BG='170;0;0'
# Set the foreground color of the `Dir` car to white
export GBT_CAR_DIR_FG='255;255;255'
```

#### Color scheme resistance

GBT is using [8-bit color
palette](https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit) to color
individual cars of the train. First 16 colors (Standart and High-intensity
colors) of the palette are prone to a change if the terminal is using some color
scheme (e.g.
[Solarized](https://en.wikipedia.org/wiki/Solarized_(color_scheme))). That means
that if one GBT train uses mixture of the first 16 and the remaining 240 colors,
the look might be inconsistent because some of the colors might change
(depending on the color scheme) and some not. Luckily the first 16 colors can be
found in the remaining 240 colors and therefore GBT can automatically convert
the first 16 colors into higher colors which provides consistent look regardless
the color scheme. This works automatically for [color names](#color-name) as
well as for [color numbers](#color-number). If needed, the automatic conversion
can be disabled with the following variable:

```shell
export GBT_FORCE_HIGHER_COLORS='0'
```


### Formatting

Formatting is done via `_FM` variables. The possible values are:

- `normal`

  Makes the text normal.

- `dim`

  Makes the text dim.

- `bold`

  Makes the text bold. Not all font characters have variant for bold formatting.

- `underline`

  Makes the text underlined.

- `blink`

  Makes the text to blink.

- `invert`

  Makes the text color inverted.

- `hide`

  Makes the text hidden.

- `none`

  No formatting applied.

  Multiple formattings can be combined into comma-separated list.

  Examples:

  ```shell
  # Set the directory name to be bold
  export GBT_CAR_DIR_FM='bold'
  # Set the directory name to be bold and underlined
  export GBT_CAR_DIR_FM='bold,underline'
  ```


### Train variables

- `GBT_CARS='Status, Os, Hostname, Dir, Git, Sign'`

  List of cars used in the train.

  To add a new car into the train, the whole variable must be redefined. For
  example in order to add the `Time` car into the default set of cars between
  the `Os` and `Hostname` car, the variable should look like this:

  ```shell
  export GBT_CARS='Status, Os, Time, Hostname, Dir, Git, Sign'
  ```

- `GBT_RCARS='Time'`

  The same like `GBT_CARS` but for the right hand side prompt.

  ```shell
  # Add the Custom car into the right hand site car to have the separator visible
  export GBT_RCARS='Custom, Time'
  # Make the Custom car to be invisible (zero length text)
  export GBT_CAR_CUSTOM_BG='default'
  export GBT_CAR_CUSTOM_FORMAT=''
  # Show only time
  export GBT_CAR_TIME_FORMAT=' {{ Time }} '
  # Set the right hand side prompt (ZSH only)
  RPROMPT='$(gbt -right)'
  ```

- `GBT_SEPARATOR=''`

  Character used to separate cars in the train.

- `GBT_RSEPARATOR=''`

  The same like `GBT_SEPARATOR` but for the right hand side prompt.

- `GBT_CAR_BG`

  Background color inherited by the top background color variable of every car.
  That allows to set the background color of all cars via single variable.

- `GBT_CAR_FG`

  Foreground color inherited by the top foreground color variable of every car.
  That allows to set the foreground color of all cars via single variable.

- `GBT_CAR_FM`

  Formatting inherited by the top formatting variable of every car. That allows
  to set the formatting of all cars via single variable.

- `GBT_BEGINNING_BG='default'`

  Background color of the text shown at the beginning of the train.

- `GBT_BEGINNING_FG='default'`

  Foreground color of the text shown at the beginning of the train.

- `GBT_BEGINNING_FM='none'`

  Formatting of the text shown at the beginning of the train.

- `GBT_BEGINNING_TEXT=''`

  Text shown at the beginning of the train.

- `GBT_SHELL`

  Indicates which shell is used. The value can be either `zsh`, `bash` or
  `plain`. By default, the value is extracted from the `$SHELL` environment
  variable. Set this variable to `bash` if your default shell is ZSH but you
  want to test GBT in Bash:

  ```shell
  export GBT_SHELL='bash'
  bash
  ```
  If set to `plain`, no shell-specific decoration is included in the output
  text. That's suitable for displaying the GBT-generated string in the console
  output.

- `GBT_DEBUG='0'`

  Shows more verbose output if some of the car modules cannot be imported.


### Cars variables

#### `Aws` car

Car that displays information about the local [AWS](https://aws.amazon.com/)
configuration.

- `GBT_CAR_AWS_BG='166'`

  Background color of the car.

- `GBT_CAR_AWS_FG='white'`

  Foreground color of the car.

- `GBT_CAR_AWS_FM='none'`

  Formatting of the car.

- `GBT_CAR_AWS_FORMAT=' {{ Icon }} {{ Project }} '`

  Format of the car.

- `GBT_CAR_AWS_ICON_BG`

  Background color of the `{{ Icon }}` element.

- `GBT_CAR_AWS_ICON_FG`

  Foreground color of the `{{ Icon }}` element.

- `GBT_CAR_AWS_ICON_FM`

  Formatting of the `{{ Icon }}` element.

- `GBT_CAR_AWS_ICON_TEXT=''`

  Text content of the `{{ Icon }}` element.

- `GBT_CAR_AWS_PROFILE_BG`

  Background color of the `{{ Profile }}` element.

- `GBT_CAR_AWS_PROFILE_FG`

  Foreground color of the `{{ Profile }}` element.

- `GBT_CAR_AWS_PROFILE_FM`

  Formatting of the `{{ Profile }}` element.

- `GBT_CAR_AWS_PROFILE_TEXT`

  Text content of the `{{ Profile }}` element specifying the configured profile.

- `GBT_CAR_AWS_REGION_BG`

  Background color of the `{{ Region }}` element.

- `GBT_CAR_AWS_REGION_FG`

  Foreground color of the `{{ Region }}` element.

- `GBT_CAR_AWS_REGION_FM`

  Formatting of the `{{ Region }}` element.

- `GBT_CAR_AWS_REGION_TEXT`

  Text content of the `{{ Region }}` element specifying the configured region.

- `GBT_CAR_AWS_DISPLAY`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_AWS_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_AWS_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_AWS_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_AWS_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_AWS_SEP_FM`

  Formatting of the separator for this car.


#### `Azure` car

Car that displays information about the local [Azure](https://azure.microsoft.com/)
configuration.

- `GBT_CAR_AZURE_BG='32'`

  Background color of the car.

- `GBT_CAR_AZURE_FG='white'`

  Foreground color of the car.

- `GBT_CAR_AZURE_FM='none'`

  Formatting of the car.

- `GBT_CAR_AZURE_FORMAT=' {{ Icon }} {{ Subscription }} '`

  Format of the car.

- `GBT_CAR_AZURE_ICON_BG`

  Background color of the `{{ Icon }}` element.

- `GBT_CAR_AZURE_ICON_FG`

  Foreground color of the `{{ Icon }}` element.

- `GBT_CAR_AZURE_ICON_FM`

  Formatting of the `{{ Icon }}` element.

- `GBT_CAR_AZURE_ICON_TEXT='ﴃ'`

  Text content of the `{{ Icon }}` element.

- `GBT_CAR_AZURE_CLOUD_BG`

  Background color of the `{{ Cloud }}` element.

- `GBT_CAR_AZURE_CLOUD_FG`

  Foreground color of the `{{ Cloud }}` element.

- `GBT_CAR_AZURE_CLOUD_FM`

  Formatting of the `{{ Cloud }}` element.

- `GBT_CAR_AZURE_CLOUD_TEXT`

  Text content of the `{{ Cloud }}` element specifying the configured cloud.

- `GBT_CAR_AZURE_SUBSCRIPTION_BG`

  Background color of the `{{ Subscription }}` element.

- `GBT_CAR_AZURE_SUBSCRIPTION_FG`

  Foreground color of the `{{ Subscription }}` element.

- `GBT_CAR_AZURE_SUBSCRIPTION_FM`

  Formatting of the `{{ Subscription }}` element.

- `GBT_CAR_AZURE_SUBSCRIPTION_TEXT`

  Text content of the `{{ Subscription }}` element specifying the configured
  subscription.

- `GBT_CAR_AZURE_USERNAME_BG`

  Background color of the `{{ UserName }}` element.

- `GBT_CAR_AZURE_USERNAME_FG`

  Foreground color of the `{{ UserName }}` element.

- `GBT_CAR_AZURE_USERNAME_FM`

  Formatting of the `{{ UserName }}` element.

- `GBT_CAR_AZURE_USERNAME_TEXT`

  Text content of the `{{ UserName }}` element specifying the configured user
  name.

- `GBT_CAR_AZURE_USERTYPE_BG`

  Background color of the `{{ UserType }}` element.

- `GBT_CAR_AZURE_USERTYPE_FG`

  Foreground color of the `{{ UserType }}` element.

- `GBT_CAR_AZURE_USERTYPE_FM`

  Formatting of the `{{ UserType }}` element.

- `GBT_CAR_AZURE_USERTYPE_TEXT`

  Text content of the `{{ UserType }}` element specifying the configured user
  type.

- `GBT_CAR_AZURE_STATE_BG`

  Background color of the `{{ State }}` element.

- `GBT_CAR_AZURE_STATE_FG`

  Foreground color of the `{{ State }}` element.

- `GBT_CAR_AZURE_STATE_FM`

  Formatting of the `{{ State }}` element.

- `GBT_CAR_AZURE_STATE_TEXT`

  Text content of the `{{ State }}` element specifying the configured
  subscription state.

- `GBT_CAR_AZURE_DEFAULTS_GROUP_BG`

  Background color of the `{{ DefaultsGroup }}` element.

- `GBT_CAR_AZURE_DEFAULTS_GROUP_FG`

  Foreground color of the `{{ DefaultsGroup }}` element.

- `GBT_CAR_AZURE_DEFAULTS_GROUP_FM`

  Formatting of the `{{ DefaultsGroup }}` element.

- `GBT_CAR_AZURE_DEFAULTS_GROUP_TEXT`

  Text content of the `{{ DefaultsGroup }}` element specifying the configured
  default resource group.

- `GBT_CAR_AZURE_DISPLAY`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_AZURE_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_AZURE_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_AZURE_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_AZURE_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_AZURE_SEP_FM`

  Formatting of the separator for this car.


#### `Custom` car

The main purpose of this car is to provide the possibility to create car with
custom text.

- `GBT_CAR_CUSTOM_BG='yellow'`

  Background color of the car.

- `GBT_CAR_CUSTOM_FG='default'`

  Foreground color of the car.

- `GBT_CAR_CUSTOM_FM='none'`

  Formatting of the car.

- `GBT_CAR_CUSTOM_FORMAT=' {{ Text }} '`

  Format of the car.

- `GBT_CAR_CUSTOM_TEXT_BG`

  Background color of the `{{ Text }}` element.

- `GBT_CAR_CUSTOM_TEXT_FG`

  Foreground color of the `{{ Text }}` element.

- `GBT_CAR_CUSTOM_TEXT_FM`

  Formatting of the `{{ Text }}` element.

- `GBT_CAR_CUSTOM_TEXT_TEXT='?'`

  Text content of the `{{ Text }}` element.

- `GBT_CAR_CUSTOM_TEXT_CMD`

  The `{{ Text }}` element will be replaced by standard output of the command
  specified in this variable. Content of the `GBT_CAR_CUSTOM_TEXT_TEXT` variable
  takes precedence over this variable.

  ```shell
  # Show 1 minute loadavg as the content of the Text element
  export GBT_CAR_CUSTOM_TEXT_CMD="uptime | sed -e 's/.*load average: //' -e 's/,.*//'"
  ```

- `GBT_CAR_CUSTOM_TEXT_EXECUTOR='sh'`

  Executor used to execute all text command (`_TEXT_CMD`).

- `GBT_CAR_CUSTOM_TEXT_EXECUTOR='-c'`

  Parameter for the executor used to execute text command (`_TEXT_CMD`).

- `GBT_CAR_CUSTOM_DISPLAY='1'`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_CUSTOM_DISPLAY_CMD`

  Command which gets executed in order to evaluate whether the car should be
  displayed or not. Content of the `GBT_CAR_CUSTOM_DISPLAY` variable takes
  precedence over this variable.

- `GBT_CAR_CUSTOM_DISPLAY_EXECUTOR='sh'`

  Executor used to execute all display command (`_TEXT_CMD`).

- `GBT_CAR_CUSTOM_DISPLAY_EXECUTOR='-c'`

  Parameter for the executor used to execute display command (`_TEXT_CMD`).

  ```shell
  # Show percentage of used disk space of the root partition
  export GBT_CAR_CUSTOM_TEXT_CMD="df -h --output=pcent / | tail -n1 | sed -re 's/\s//g' -e 's/%/%%/'"
  # Display the car only if the percentage is above 90%
  export GBT_CAR_CUSTOM_DISPLAY_CMD="[[ $(df -h --output=pcent / | tail -n1 | sed -re 's/\s//g' -e 's/%//') -gt 70 ]] && echo YES"
  ```

- `GBT_CAR_CUSTOM_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_CUSTOM_EXECUTOR='sh'`

  Executor used to execute all custom commands (`_TEXT_CMD` and `_DISPLAY_CMD`).

- `GBT_CAR_CUSTOM_EXECUTOR='-c'`

  Parameter for the executor used to execute all custom commands (`_TEXT_CMD`
  and `_DISPLAY_CMD`).

- `GBT_CAR_CUSTOM_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_CUSTOM_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_CUSTOM_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_CUSTOM_SEP_FM`

  Formatting of the separator for this car.

Multiple `Custom` cars can be used in the `GBT_CARS` variable. Just add some
identifier behind the car name. To set properties of the new car, just add the
same identifier into the environment variable:

```shell
# Adding Custom and Custom1 car
export GBT_CARS='Status, Os, Custom, Custom1, Hostname, Dir, Git, Sign'
# The text of the default Custom car
export GBT_CAR_CUSTOM_TEXT_TEXT='default'
# The text of the Custom1 car
export GBT_CAR_CUSTOM1_TEXT_TEXT='1'
# Set different background color for the Custom1 car
export GBT_CAR_CUSTOM1_BG='magenta'
```


#### `Dir` car

Car that displays current directory name.

- `GBT_CAR_DIR_BG='blue'`

  Background color of the car.

- `GBT_CAR_DIR_FG='light_gray'`

  Foreground color of the car.

- `GBT_CAR_DIR_FM='none'`

  Formatting of the car.

- `GBT_CAR_DIR_FORMAT=' {{ Dir }} '`

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

- `GBT_CAR_DIR_HOMESIGN='~'`

  Character which represents the user's home directory. If set to empty
  string, full home directory path is used instead.

- `GBT_CAR_DIR_DEPTH='1'`

  Number of directories to show.

- `GBT_CAR_DIR_NONCURLEN='255'`

  Indicates how many characters of the non-current directory name should be
  displayed. This can be set to `1` to display only the first character of the
  directory name when using `GBT_CAR_DIR_DEPTH` with value grater than one.

- `GBT_CAR_DIR_DISPLAY='1'`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_DIR_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_DIR_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_DIR_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_DIR_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_DIR_SEP_FM`

  Formatting of the separator for this car.


#### `ExecTime` car

Car that displays how long each shell command run.

- `GBT_CAR_EXECTIME_BG='light_gray'`

  Background color of the car.

- `GBT_CAR_EXECTIME_FG='black'`

  Foreground color of the car.

- `GBT_CAR_EXECTIME_FM='none'`

  Formatting of the car.

- `GBT_CAR_EXECTIME_FORMAT=' {{ Time }} '`

  Format of the car.

- `GBT_CAR_EXECTIME_DURATION_BG`

  Background color of the `{{ Duration }}` element.

- `GBT_CAR_EXECTIME_DURATION_FG`

  Foreground color of the `{{ Duration }}` element.

- `GBT_CAR_EXECTIME_DURATION_FM`

  Formatting of the `{{ Duration }}` element.

- `GBT_CAR_EXECTIME_DURATION_TEXT`

  Text content of the `{{ Duration }}` element. The duration of the execution
  time (e.g `1h8m19s135ms` for precision set to `3`).

- `GBT_CAR_EXECTIME_SECONDS_BG`

  Background color of the `{{ Seconds }}` element.

- `GBT_CAR_EXECTIME_SECONDS_FG`

  Foreground color of the `{{ Seconds }}` element.

- `GBT_CAR_EXECTIME_SECONDS_FM`

  Formatting of the `{{ Seconds }}` element.

- `GBT_CAR_EXECTIME_SECONDS_TEXT`

  Text content of the `{{ Seconds }}` element. The execution time in seconds
  (e.g. `4099.1358` for precision set to `4`).

- `GBT_CAR_EXECTIME_TIME_BG`

  Background color of the `{{ Time }}` element.

- `GBT_CAR_EXECTIME_TIME_FG`

  Foreground color of the `{{ Time }}` element.

- `GBT_CAR_EXECTIME_TIME_FM`

  Formatting of the `{{ Time }}` element.

- `GBT_CAR_EXECTIME_TIME_TEXT`

  Text content of the `{{ Time }}` element. The execution time (e.g.
  `01:08:19.1358` for precision set to `4`).

- `GBT_CAR_EXECTIME_PRECISION='0'`

  Sub-second precision to show.

- `GBT_CAR_EXECTIME_SECS`

  The number of seconds the command run in shell. This variable is defined in
  the source file as shown bellow.

- `GBT_CAR_EXECTIME_BELL='0'`

  Sound console bell if the executed command exceeds specified number of
  seconds. Value set to `0` disables the bell (default).

- `GBT_CAR_EXECTIME_DISPLAY='1'`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_EXECTIME_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_EXECTIME_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_EXECTIME_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_EXECTIME_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_EXECTIME_SEP_FM`

  Formatting of the separator for this car.

In order to allow this car to calculate the execution time, the following must
be loaded in the shell:

```shell
# For Bash
source /usr/share/gbt/sources/exectime/bash.sh
# For ZSH
source /usr/share/gbt/sources/exectime/zsh.sh
```

On macOS the `date` command does not support `%N` format for milliseconds and
you need to override the environment variable `GBT__SOURCE_DATE_ARG='+%s`.


#### `Gcp` car

Car that displays information about the local [GCP](https://cloud.google.com/)
configuration.

- `GBT_CAR_GCP_BG='33'`

  Background color of the car.

- `GBT_CAR_GCP_FG='white'`

  Foreground color of the car.

- `GBT_CAR_GCP_FM='none'`

  Formatting of the car.

- `GBT_CAR_GCP_FORMAT=' {{ Icon }} {{ Project }} '`

  Format of the car.

- `GBT_CAR_GCP_ICON_BG`

  Background color of the `{{ Icon }}` element.

- `GBT_CAR_GCP_ICON_FG`

  Foreground color of the `{{ Icon }}` element.

- `GBT_CAR_GCP_ICON_FM`

  Formatting of the `{{ Icon }}` element.

- `GBT_CAR_GCP_ICON_TEXT=''`

  Text content of the `{{ Icon }}` element.

- `GBT_CAR_GCP_ACCOUNT_BG`

  Background color of the `{{ Account }}` element.

- `GBT_CAR_GCP_ACCOUNT_FG`

  Foreground color of the `{{ Account }}` element.

- `GBT_CAR_GCP_ACCOUNT_FM`

  Formatting of the `{{ Account }}` element.

- `GBT_CAR_GCP_ACCOUNT_TEXT`

  Text content of the `{{ Account }}` element specifying the configured account.

- `GBT_CAR_GCP_CONFIG_BG`

  Background color of the `{{ Config }}` element.

- `GBT_CAR_GCP_CONFIG_FG`

  Foreground color of the `{{ Config }}` element.

- `GBT_CAR_GCP_CONFIG_FM`

  Formatting of the `{{ Config }}` element.

- `GBT_CAR_GCP_CONFIG_TEXT`

  Text content of the `{{ Config }}` element specifying the active
  configuration.

- `GBT_CAR_GCP_PROJECT_BG`

  Background color of the `{{ Project }}` element.

- `GBT_CAR_GCP_PROJECT_FG`

  Foreground color of the `{{ Project }}` element.

- `GBT_CAR_GCP_PROJECT_FM`

  Formatting of the `{{ Project }}` element.

- `GBT_CAR_GCP_PROJECT_TEXT`

  Text content of the `{{ Project }}` element specifying the configured project.

- `GBT_CAR_GCP_PROJECT_ALIASES`

  List of aliases that allow to display different project name based on the
  original name. The following example shows how to change the project
  `my-dev-project-123456` to `dev` and the project `my-prod-project-654321` to
  `prod`.

  ```shell
  export GBT_CAR_GCP_PROJECT_ALIASES='
    my-dev-project-123456=dev,
    my-prod-project-654321=prod,
  '
  ```

- `GBT_CAR_GCP_DISPLAY`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_GCP_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_GCP_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_GCP_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_GCP_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_GCP_SEP_FM`

  Formatting of the separator for this car.


#### `Git` car

Car that displays information about a local Git repository. This car is
displayed only if the current directory is a Git repository.

- `GBT_CAR_GIT_BG='light_gray'`

  Background color of the car.

- `GBT_CAR_GIT_FG='black'`

  Foreground color of the car.

- `GBT_CAR_GIT_FM='none'`

  Formatting of the car.

- `GBT_CAR_GIT_FORMAT=' {{ Icon }} {{ Head }} {{ Status }}{{ Ahead }}{{ Behind }} '`

  Format of the car.

- `GBT_CAR_GIT_ICON_BG`

  Background color of the `{{ Icon }}` element.

- `GBT_CAR_GIT_ICON_FG`

  Foreground color of the `{{ Icon }}` element.

- `GBT_CAR_GIT_ICON_FM`

  Formatting of the `{{ Icon }}` element.

- `GBT_CAR_GIT_ICON_TEXT=''`

  Text content of the `{{ Icon }}` element.

- `GBT_CAR_GIT_HEAD_BG`

  Background color of the `{{ Head }}` element.

- `GBT_CAR_GIT_HEAD_FG`

  Foreground color of the `{{ Head }}` element.

- `GBT_CAR_GIT_HEAD_FM`

  Formatting of the `{{ Head }}` element.

- `GBT_CAR_GIT_HEAD_TEXT`

  Text content of the `{{ Head }}` element - branch, tag name or the
  commit ID.

- `GBT_CAR_GIT_STATUS_BG`

  Background color of the `{{ Status }}` element.

- `GBT_CAR_GIT_STATUS_FG`

  Foreground color of the `{{ Status }}` element.

- `GBT_CAR_GIT_STATUS_FM`

  Formatting of the `{{ Status }}` element.

- `GBT_CAR_GIT_STATUS_FORMAT`

  Format of the `{{ Status }}` element. The content is either
  `{{ StatusDirty }}` or `{{ StatusClean }}` depending on the state of the
  local Git repository.

- `GBT_CAR_GIT_STATUS_DIRTY_BG`

  Background color of the `{{ StatusDirty }}` element.

- `GBT_CAR_GIT_STATUS_DIRTY_FG='red'`

  Foreground color of the `{{ StatusDirty }}` element.

- `GBT_CAR_GIT_STATUS_DIRTY_FM`

  Formatting of the `{{ StatusDirty }}` element.

- `GBT_CAR_GIT_STATUS_DIRTY_TEXT='✘'`

  Text content of the `{{ StatusDirty }}` element.

- `GBT_CAR_GIT_STATUS_CLEAN_BG`

  Background color of the `{{ StatusClean }}` element.

- `GBT_CAR_GIT_STATUS_CLEAN_FG='green'`

  Foreground color of the `{{ StatusClean }}` element.

- `GBT_CAR_GIT_STATUS_CLEAN_FM`

  Formatting of the `{{ StatusClean }}` element.

- `GBT_CAR_GIT_STATUS_CLEAN_TEXT='✔'`

  Text content of the `{{ StatusClean }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_BG`

  Background color of the `{{ StatusAdded }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_FG`

  Foreground color of the `{{ StatusAdded }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_FM`

  Formatting of the `{{ StatusAdded }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_FORMAT='{{ StatusAddedSymbol }}'`

  Format of the the `{{ StatusAdded }}` element. It can be
  `{{ StatusAddedSymbol }}` or `{{ StatusAddedCount }}`.

- `GBT_CAR_GIT_STATUS_ADDED_SYMBOL_BG`

  Background color of the `{{ StatusAddedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_SYMBOL_FG`

  Foreground color of the `{{ StatusAddedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_SYMBOL_FM`

  Formatting of the `{{ StatusAddedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_SYMBOL_TEXT=' ⟴'`

  Text content of the `{{ StatusAddedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_COUNT_BG`

  Background color of the `{{ StatusAddedCount }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_COUNT_FG`

  Foreground color of the `{{ StatusAddedCount }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_COUNT_FM`

  Formatting of the `{{ StatusAddedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_ADDED_COUNT_TEXT`

  Text content of the `{{ StatusAddedCount }}` element. By default it contains
  a number of added files.

- `GBT_CAR_GIT_STATUS_COPIED_BG`

  Background color of the `{{ StatusCopied }}` element.

- `GBT_CAR_GIT_STATUS_COPIED_FG`

  Foreground color of the `{{ StatusCopied }}` element.

- `GBT_CAR_GIT_STATUS_COPIED_FM`

  Formatting of the `{{ StatusCopied }}` element.

- `GBT_CAR_GIT_STATUS_COPIED_FORMAT='{{ StatusCopiedSymbol }}'`

  Format of the the `{{ StatusCopied }}` element. It can be
  `{{ StatusCopiedSymbol }}` or `{{ StatusCopiedCount }}`.

- `GBT_CAR_GIT_STATUS_COPIED_SYMBOL_BG`

  Background color of the `{{ StatusCopiedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_COPIED_SYMBOL_FG`

  Foreground color of the `{{ StatusCopiedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_COPIED_SYMBOL_FM`

  Formatting of the `{{ StatusCopiedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_COPIED_SYMBOL_TEXT=' ⥈'`

  Text content of the `{{ StatusCopiedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_COPIED_COUNT_BG`

  Background color of the `{{ StatusCopiedCount }}` element.

- `GBT_CAR_GIT_STATUS_COPIED_COUNT_FG`

  Foreground color of the `{{ StatusCopiedCount }}` element.

- `GBT_CAR_GIT_STATUS_COPIED_COUNT_FM`

  Formatting of the `{{ StatusCopiedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_COPIED_COUNT_TEXT`

  Text content of the `{{ StatusCopiedCount }}` element. By default it contains
  a number of files copied.

- `GBT_CAR_GIT_STATUS_DELETED_BG`

  Background color of the `{{ StatusDeleted }}` element.

- `GBT_CAR_GIT_STATUS_DELETED_FG`

  Foreground color of the `{{ StatusDeleted }}` element.

- `GBT_CAR_GIT_STATUS_DELETED_FM`

  Formatting of the `{{ StatusDeleted }}` element.

- `GBT_CAR_GIT_STATUS_DELETED_FORMAT='{{ StatusDeletedSymbol }}'`

  Format of the the `{{ StatusDeleted }}` element. It can be
  `{{ StatusDeletedSymbol }}` or `{{ StatusDeletedCount }}`.

- `GBT_CAR_GIT_STATUS_DELETED_SYMBOL_BG`

  Background color of the `{{ StatusDeletedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_DELETED_SYMBOL_FG`

  Foreground color of the `{{ StatusDeletedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_DELETED_SYMBOL_FM`

  Formatting of the `{{ StatusDeletedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_DELETED_SYMBOL_TEXT=' ➖'`

  Text content of the `{{ StatusDeletedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_DELETED_COUNT_BG`

  Background color of the `{{ StatusDeletedCount }}` element.

- `GBT_CAR_GIT_STATUS_DELETED_COUNT_FG`

  Foreground color of the `{{ StatusDeletedCount }}` element.

- `GBT_CAR_GIT_STATUS_DELETED_COUNT_FM`

  Formatting of the `{{ StatusDeletedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_DELETED_COUNT_TEXT`

  Text content of the `{{ StatusDeletedCount }}` element. By default it contains
  a number of deleted files.

- `GBT_CAR_GIT_STATUS_IGNORED_BG`

  Background color of the `{{ StatusIgnored }}` element.

- `GBT_CAR_GIT_STATUS_IGNORED_FG`

  Foreground color of the `{{ StatusIgnored }}` element.

- `GBT_CAR_GIT_STATUS_IGNORED_FM`

  Formatting of the `{{ StatusIgnored }}` element.

- `GBT_CAR_GIT_STATUS_IGNORED_FORMAT='{{ StatusIgnoredSymbol }}'`

  Format of the the `{{ StatusIgnored }}` element. It can be
  `{{ StatusIgnoredSymbol }}` or `{{ StatusIgnoredCount }}`.

- `GBT_CAR_GIT_STATUS_IGNORED_SYMBOL_BG`

  Background color of the `{{ StatusIgnoredSymbol }}` element.

- `GBT_CAR_GIT_STATUS_IGNORED_SYMBOL_FG`

  Foreground color of the `{{ StatusIgnoredSymbol }}` element.

- `GBT_CAR_GIT_STATUS_IGNORED_SYMBOL_FM`

  Formatting of the `{{ StatusIgnoredSymbol }}` element.

- `GBT_CAR_GIT_STATUS_IGNORED_SYMBOL_TEXT=' ⬆'`

  Text content of the `{{ StatusIgnoredSymbol }}` element.

- `GBT_CAR_GIT_STATUS_IGNORED_COUNT_BG`

  Background color of the `{{ StatusIgnoredCount }}` element.

- `GBT_CAR_GIT_STATUS_IGNORED_COUNT_FG`

  Foreground color of the `{{ StatusIgnoredCount }}` element.

- `GBT_CAR_GIT_STATUS_IGNORED_COUNT_FM`

  Formatting of the `{{ StatusIgnoredSymbol }}` element.

- `GBT_CAR_GIT_STATUS_IGNORED_COUNT_TEXT`

  Text content of the `{{ StatusIgnoredCount }}` element. By default it contains
  a number of ignored files.

- `GBT_CAR_GIT_STATUS_MODIFIED_BG`

  Background color of the `{{ StatusModified }}` element.

- `GBT_CAR_GIT_STATUS_MODIFIED_FG`

  Foreground color of the `{{ StatusModified }}` element.

- `GBT_CAR_GIT_STATUS_MODIFIED_FM`

  Formatting of the `{{ StatusModified }}` element.

- `GBT_CAR_GIT_STATUS_MODIFIED_FORMAT='{{ StatusModifiedSymbol }}'`

  Format of the the `{{ StatusModified }}` element. It can be
  `{{ StatusModifiedSymbol }}` or `{{ StatusModifiedCount }}`.

- `GBT_CAR_GIT_STATUS_MODIFIED_SYMBOL_BG`

  Background color of the `{{ StatusModifiedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_MODIFIED_SYMBOL_FG`

  Foreground color of the `{{ StatusModifiedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_MODIFIED_SYMBOL_FM`

  Formatting of the `{{ StatusModifiedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_MODIFIED_SYMBOL_TEXT=' ⬆'`

  Text content of the `{{ StatusModifiedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_MODIFIED_COUNT_BG`

  Background color of the `{{ StatusModifiedCount }}` element.

- `GBT_CAR_GIT_STATUS_MODIFIED_COUNT_FG`

  Foreground color of the `{{ StatusModifiedCount }}` element.

- `GBT_CAR_GIT_STATUS_MODIFIED_COUNT_FM`

  Formatting of the `{{ StatusModifiedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_MODIFIED_COUNT_TEXT`

  Text content of the `{{ StatusModifiedCount }}` element. By default it
  contains a number of modified files.

- `GBT_CAR_GIT_STATUS_RENAMED_BG`

  Background color of the `{{ StatusRenamed }}` element.

- `GBT_CAR_GIT_STATUS_RENAMED_FG`

  Foreground color of the `{{ StatusRenamed }}` element.

- `GBT_CAR_GIT_STATUS_RENAMED_FM`

  Formatting of the `{{ StatusRenamed }}` element.

- `GBT_CAR_GIT_STATUS_RENAMED_FORMAT='{{ StatusRenamedSymbol }}'`

  Format of the the `{{ StatusRenamed }}` element. It can be
  `{{ StatusRenamedSymbol }}` or `{{ StatusRenamedCount }}`.

- `GBT_CAR_GIT_STATUS_RENAMED_SYMBOL_BG`

  Background color of the `{{ StatusRenamedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_RENAMED_SYMBOL_FG`

  Foreground color of the `{{ StatusRenamedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_RENAMED_SYMBOL_FM`

  Formatting of the `{{ StatusRenamedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_RENAMED_SYMBOL_TEXT=' ⬆'`

  Text content of the `{{ StatusRenamedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_RENAMED_COUNT_BG`

  Background color of the `{{ StatusRenamedCount }}` element.

- `GBT_CAR_GIT_STATUS_RENAMED_COUNT_FG`

  Foreground color of the `{{ StatusRenamedCount }}` element.

- `GBT_CAR_GIT_STATUS_RENAMED_COUNT_FM`

  Formatting of the `{{ StatusRenamedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_RENAMED_COUNT_TEXT`

  Text content of the `{{ StatusRenamedCount }}` element. By default it contains
  a number of renamed files.

- `GBT_CAR_GIT_STATUS_STAGED_BG`

  Background color of the `{{ StatusStaged }}` element.

- `GBT_CAR_GIT_STATUS_STAGED_FG`

  Foreground color of the `{{ StatusStaged }}` element.

- `GBT_CAR_GIT_STATUS_STAGED_FM`

  Formatting of the `{{ StatusStaged }}` element.

- `GBT_CAR_GIT_STATUS_STAGED_FORMAT='{{ StatusStagedSymbol }}'`

  Format of the the `{{ StatusStaged }}` element. It can be
  `{{ StatusStagedSymbol }}` or `{{ StatusStagedCount }}`.

- `GBT_CAR_GIT_STATUS_STAGED_SYMBOL_BG`

  Background color of the `{{ StatusStagedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_STAGED_SYMBOL_FG`

  Foreground color of the `{{ StatusStagedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_STAGED_SYMBOL_FM`

  Formatting of the `{{ StatusStagedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_STAGED_SYMBOL_TEXT=' ⬆'`

  Text content of the `{{ StatusStagedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_STAGED_COUNT_BG`

  Background color of the `{{ StatusStagedCount }}` element.

- `GBT_CAR_GIT_STATUS_STAGED_COUNT_FG`

  Foreground color of the `{{ StatusStagedCount }}` element.

- `GBT_CAR_GIT_STATUS_STAGED_COUNT_FM`

  Formatting of the `{{ StatusStagedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_STAGED_COUNT_TEXT`

  Text content of the `{{ StatusStagedCount }}` element. By default it contains
  a number of staged files.

- `GBT_CAR_GIT_STATUS_UNMERGED_BG`

  Background color of the `{{ StatusUnmerged }}` element.

- `GBT_CAR_GIT_STATUS_UNMERGED_FG`

  Foreground color of the `{{ StatusUnmerged }}` element.

- `GBT_CAR_GIT_STATUS_UNMERGED_FM`

  Formatting of the `{{ StatusUnmerged }}` element.

- `GBT_CAR_GIT_STATUS_UNMERGED_FORMAT='{{ StatusUnmergedSymbol }}'`

  Format of the the `{{ StatusUnmerged }}` element. It can be
  `{{ StatusUnmergedSymbol }}` or `{{ StatusUnmergedCount }}`.

- `GBT_CAR_GIT_STATUS_UNMERGED_SYMBOL_BG`

  Background color of the `{{ StatusUnmergedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_UNMERGED_SYMBOL_FG`

  Foreground color of the `{{ StatusUnmergedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_UNMERGED_SYMBOL_FM`

  Formatting of the `{{ StatusUnmergedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_UNMERGED_SYMBOL_TEXT=' ⬆'`

  Text content of the `{{ StatusUnmergedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_UNMERGED_COUNT_BG`

  Background color of the `{{ StatusUnmergedCount }}` element.

- `GBT_CAR_GIT_STATUS_UNMERGED_COUNT_FG`

  Foreground color of the `{{ StatusUnmergedCount }}` element.

- `GBT_CAR_GIT_STATUS_UNMERGED_COUNT_FM`

  Formatting of the `{{ StatusUnmergedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_UNMERGED_COUNT_TEXT`

  Text content of the `{{ StatusUnmergedCount }}` element. By default it
  contains a number of unmerged files.

- `GBT_CAR_GIT_STATUS_UNTRACKED_BG`

  Background color of the `{{ StatusUntracked }}` element.

- `GBT_CAR_GIT_STATUS_UNTRACKED_FG`

  Foreground color of the `{{ StatusUntracked }}` element.

- `GBT_CAR_GIT_STATUS_UNTRACKED_FM`

  Formatting of the `{{ StatusUntracked }}` element.

- `GBT_CAR_GIT_STATUS_UNTRACKED_FORMAT='{{ StatusUntrackedSymbol }}'`

  Format of the the `{{ StatusUntracked }}` element. It can be
  `{{ StatusUntrackedSymbol }}` or `{{ StatusUntrackedCount }}`.

- `GBT_CAR_GIT_STATUS_UNTRACKED_SYMBOL_BG`

  Background color of the `{{ StatusUntrackedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_UNTRACKED_SYMBOL_FG`

  Foreground color of the `{{ StatusUntrackedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_UNTRACKED_SYMBOL_FM`

  Formatting of the `{{ StatusUntrackedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_UNTRACKED_SYMBOL_TEXT=' ⬆'`

  Text content of the `{{ StatusUntrackedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_BG`

  Background color of the `{{ StatusUntrackedCount }}` element.

- `GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_FG`

  Foreground color of the `{{ StatusUntrackedCount }}` element.

- `GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_FM`

  Formatting of the `{{ StatusUntrackedSymbol }}` element.

- `GBT_CAR_GIT_STATUS_UNTRACKED_COUNT_TEXT`

  Text content of the `{{ StatusUntrackedCount }}` element. By default it
  contains a number of untracked files.

- `GBT_CAR_GIT_AHEAD_BG`

  Background color of the `{{ Ahead }}` element.

- `GBT_CAR_GIT_AHEAD_FG`

  Foreground color of the `{{ Ahead }}` element.

- `GBT_CAR_GIT_AHEAD_FM`

  Formatting of the `{{ Ahead }}` element.

- `GBT_CAR_GIT_AHEAD_FORMAT='{{ AheadSymbol }}'`

  Format of the the `{{ Ahead }}` element. It can be `{{ AheadSymbol }}` or
  `{{ AheadCount }}`.

- `GBT_CAR_GIT_AHEAD_SYMBOL_BG`

  Background color of the `{{ AheadSymbol }}` element.

- `GBT_CAR_GIT_AHEAD_SYMBOL_FG`

  Foreground color of the `{{ AheadSymbol }}` element.

- `GBT_CAR_GIT_AHEAD_SYMBOL_FM`

  Formatting of the `{{ AheadSymbol }}` element.

- `GBT_CAR_GIT_AHEAD_SYMBOL_TEXT=' ⬆'`

  Text content of the `{{ AheadSymbol }}` element.

- `GBT_CAR_GIT_AHEAD_COUNT_BG`

  Background color of the `{{ AheadCount }}` element.

- `GBT_CAR_GIT_AHEAD_COUNT_FG`

  Foreground color of the `{{ AheadCount }}` element.

- `GBT_CAR_GIT_AHEAD_COUNT_FM`

  Formatting of the `{{ AheadSymbol }}` element.

- `GBT_CAR_GIT_AHEAD_COUNT_TEXT`

  Text content of the `{{ AheadCount }}` element. By default it contains
  a number of commits ahead of the remote branch.

- `GBT_CAR_GIT_BEHIND_BG`

  Background color of the `{{ Behind }}` element.

- `GBT_CAR_GIT_BEHIND_FG`

  Foreground color of the `{{ Behind }}` element.

- `GBT_CAR_GIT_BEHIND_FM`

  Formatting of the `{{ Behind }}` element.

- `GBT_CAR_GIT_BEHIND_FORMAT='{{ BehindSymbol }}'`

  Format of the the `{{ Behind }}` element. It can be `{{ BehindSymbol }}` or
  `{{ BehindCount }}`.

- `GBT_CAR_GIT_BEHIND_SYMBOL_BG`

  Background color of the `{{ BehindSymbol }}` element.

- `GBT_CAR_GIT_BEHIND_SYMBOL_FG`

  Foreground color of the `{{ BehindSymbol }}` element.

- `GBT_CAR_GIT_BEHIND_SYMBOL_FM`

  Formatting of the `{{ BehindSymbol }}` element.

- `GBT_CAR_GIT_BEHIND_SYMBOL_TEXT=' ⬇'`

  Text content of the `{{ BehindSymbol }}` element.

- `GBT_CAR_GIT_BEHIND_COUNT_BG`

  Background color of the `{{ BehindCount }}` element.

- `GBT_CAR_GIT_BEHIND_COUNT_FG`

  Foreground color of the `{{ BehindCount }}` element.

- `GBT_CAR_GIT_BEHIND_COUNT_FM`

  Formatting of the `{{ BehindSymbol }}` element.

- `GBT_CAR_GIT_BEHIND_COUNT_TEXT`

  Text content of the `{{ BehindCount }}` element. By default it contains
  a number of commits ahead of the remote branch.

- `GBT_CAR_GIT_STASH_BG`

  Background color of the `{{ Stash }}` element.

- `GBT_CAR_GIT_STASH_FG`

  Foreground color of the `{{ Stash }}` element.

- `GBT_CAR_GIT_STASH_FM`

  Formatting of the `{{ Stash }}` element.

- `GBT_CAR_GIT_STASH_FORMAT='{{ StashSymbol }}'`

  Format of the the `{{ Stash }}` element. It can be `{{ StashSymbol }}` or
  `{{ StashCount }}`.

- `GBT_CAR_GIT_STASH_SYMBOL_BG`

  Background color of the `{{ StashSymbol }}` element.

- `GBT_CAR_GIT_STASH_SYMBOL_FG`

  Foreground color of the `{{ StashSymbol }}` element.

- `GBT_CAR_GIT_STASH_SYMBOL_FM`

  Formatting of the `{{ StashSymbol }}` element.

- `GBT_CAR_GIT_STASH_SYMBOL_TEXT=' ⚑'`

  Text content of the `{{ StashSymbol }}` element.

- `GBT_CAR_GIT_STASH_COUNT_BG`

  Background color of the `{{ StashCount }}` element.

- `GBT_CAR_GIT_STASH_COUNT_FG`

  Foreground color of the `{{ StashCount }}` element.

- `GBT_CAR_GIT_STASH_COUNT_FM`

  Formatting of the `{{ StashSymbol }}` element.

- `GBT_CAR_GIT_STASH_COUNT_TEXT`

  Text content of the `{{ StashCount }}` element. By default it contains
  a number of stashes.

- `GBT_CAR_GIT_DISPLAY`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_GIT_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_GIT_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_GIT_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_GIT_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_GIT_SEP_FM`

  Formatting of the separator for this car.


#### `Hostname` car

Car that displays username of the currently logged user and the hostname of the
local machine.

- `GBT_CAR_HOSTNAME_BG='dark_gray'`

  Background color of the car.

- `GBT_CAR_HOSTNAME_FG='252'`

  Foreground color of the car.

- `GBT_CAR_HOSTNAME_FM='none'`

  Formatting of the car.

- `GBT_CAR_HOSTNAME_FORMAT=' {{ UserHost }} '`

  Format of the car.

- `GBT_CAR_HOSTNAME_USERHOST_BG`

  Background color of the `{{ UserHost }}` element.

- `GBT_CAR_HOSTNAME_USERHOST_FG`

  Foreground color of the `{{ UserHost }}` element.

- `GBT_CAR_HOSTNAME_USERHOST_FM`

  Formatting of the `{{ UserHost }}` element.

- `GBT_CAR_HOSTNAME_USERHOST_FORMAT`

  Format of the `{{ UserHost }}` element. The value is either
  `{{ Admin }}@{{ Host }}` if the user is `root` or `{{ User }}@{{ Host }}`
  if the user is a normal user.

- `GBT_CAR_HOSTNAME_ADMIN_BG`

  Background color of the `{{ Admin }}` element.

- `GBT_CAR_HOSTNAME_ADMIN_FG`

  Foreground color of the `{{ Admin }}` element.

- `GBT_CAR_HOSTNAME_ADMIN_FM`

  Formatting of the `{{ Admin }}` element.

- `GBT_CAR_HOSTNAME_ADMIN_TEXT`

  Text content of the `{{ Admin }}` element. The user name.

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

- `GBT_CAR_HOSTNAME_DISPLAY='1'`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_HOSTNAME_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_HOSTNAME_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_HOSTNAME_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_HOSTNAME_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_HOSTNAME_SEP_FM`

  Formatting of the separator for this car.


#### `Kubectl` car

Car that displays `kubectl` information.

- `GBT_CAR_KUBECTL_BG='26'`

  Background color of the car.

- `GBT_CAR_KUBECTL_FG='white'`

  Foreground color of the car.

- `GBT_CAR_KUBECTL_FM='none'`

  Formatting of the car.

- `GBT_CAR_KUBECTL_FORMAT=' {{ Icon }} {{ Context }} '`

  Format of the car. `{{ Cluster }}`, `{{ AuthInfo }}` and `{{ Namespace }}`
  can be used here as well.

- `GBT_CAR_KUBECTL_ICON_BG`

  Background color of the `{{ Icon }}` element.

- `GBT_CAR_KUBECTL_ICON_FG`

  Foreground color of the `{{ Icon }}` element.

- `GBT_CAR_KUBECTL_ICON_FM`

  Formatting of the `{{ Icon }}` element.

- `GBT_CAR_KUBECTL_ICON_TEXT='⎈'`

  Text content of the `{{ Icon }}` element.

- `GBT_CAR_KUBECTL_CONTEXT_BG`

  Background color of the `{{ Context }}` element.

- `GBT_CAR_KUBECTL_CONTEXT_FG`

  Foreground color of the `{{ Context }}` element.

- `GBT_CAR_KUBECTL_CONTEXT_FM`

  Formatting of the `{{ Context }}` element.

- `GBT_CAR_KUBECTL_CONTEXT_TEXT`

  Text content of the `{{ Context }}` element.

- `GBT_CAR_KUBECTL_CLUSTER_BG`

  Background color of the `{{ Cluster }}` element.

- `GBT_CAR_KUBECTL_CLUSTER_FG`

  Foreground color of the `{{ Cluster }}` element.

- `GBT_CAR_KUBECTL_CLUSTER_FM`

  Formatting of the `{{ Cluster }}` element.

- `GBT_CAR_KUBECTL_CLUSTER_TEXT`

  Text content of the `{{ Cluster }}` element.

- `GBT_CAR_KUBECTL_AUTHINFO_BG`

  Background color of the `{{ AuthInfo }}` element.

- `GBT_CAR_KUBECTL_AUTHINFO_FG`

  Foreground color of the `{{ AuthInfo }}` element.

- `GBT_CAR_KUBECTL_AUTHINFO_FM`

  Formatting of the `{{ AuthInfo }}` element.

- `GBT_CAR_KUBECTL_AUTHINFO_TEXT`

  Text content of the `{{ AuthInfo }}` element.

- `GBT_CAR_KUBECTL_NAMESPACE_BG`

  Background color of the `{{ Namespace }}` element.

- `GBT_CAR_KUBECTL_NAMESPACE_FG`

  Foreground color of the `{{ Namespace }}` element.

- `GBT_CAR_KUBECTL_NAMESPACE_FM`

  Formatting of the `{{ Namespace }}` element.

- `GBT_CAR_KUBECTL_NAMESPACE_TEXT`

  Text content of the `{{ Namespace }}` element.

- `GBT_CAR_KUBECTL_DISPLAY='1'`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_KUBECTL_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_KUBECTL_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_KUBECTL_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_KUBECTL_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_KUBECTL_SEP_FM`

  Formatting of the separator for this car.


#### `Os` car

Car that displays icon of the operating system.

- `GBT_CAR_OS_BG='235'`

  Background color of the car.

- `GBT_CAR_OS_FG='white'`

  Foreground color of the car.

- `GBT_CAR_OS_FM='none'`

  Formatting of the car.

- `GBT_CAR_OS_FORMAT=' {{ Symbol }} '`

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
  export GBT_CAR_OS_NAME='arch'
  ```

- `GBT_CAR_OS_DISPLAY='1'`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_OS_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_OS_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_OS_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_OS_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_OS_SEP_FM`

  Formatting of the separator for this car.


#### `PyVirtEnv` car

Car that displays Python Virtual Environment name. This car is displayed only
if the Python Virtual Environment is activated. The activation script usually
prepends the shell prompt by the Virtual Environment name by default. In order
to disable it, the following environment variable must be set:

```shell
export VIRTUAL_ENV_DISABLE_PROMPT='1'
```

Variables used by the car:

- `GBT_CAR_PYVIRTENV_BG='222'`

  Background color of the car.

- `GBT_CAR_PYVIRTENV_FG='black'`

  Foreground color of the car.

- `GBT_CAR_PYVIRTENV_FM='none'`

  Formatting of the car.

- `GBT_CAR_PYVIRTENV_FORMAT=' {{ Icon }} {{ Name }} '`

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

- `GBT_CAR_PYVIRTENV_NAME_FG='33'`

  Foreground color of the `{{ NAME }}` element.

- `GBT_CAR_PYVIRTENV_NAME_FM`

  Formatting of the `{{ Name }}` element.

- `GBT_CAR_PYVIRTENV_NAME_TEXT`

  The name of the Python Virtual Environment deducted from the `VIRTUAL_ENV`
  environment variable.

- `GBT_CAR_PYVIRTENV_DISPLAY`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_PYVIRTENV_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_PYVIRTENV_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_PYVIRTENV_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_PYVIRTENV_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_PYVIRTENV_SEP_FM`

  Formatting of the separator for this car.


#### `Sign` car

Car that displays prompt character for the admin and user at the end of the
train.

- `GBT_CAR_SIGN_BG='default'`

  Background color of the car.

- `GBT_CAR_SIGN_FG='default'`

  Foreground color of the car.

- `GBT_CAR_SIGN_FM='none'`

  Formatting of the car.

- `GBT_CAR_SIGN_FORMAT=' {{ Symbol }} '`

  Format of the car.

- `GBT_CAR_SIGN_SYMBOL_BG`

  Background color of the `{{ Symbol }}` element.

- `GBT_CAR_SIGN_SYMBOL_FG`

  Foreground color of the `{{ Symbol }}` element.

- `GBT_CAR_SIGN_SYMBOL_FM='bold'`

  Formatting of the `{{ Symbol }}` element.

- `GBT_CAR_SIGN_SYMBOL_FORMAT`

  Format of the `{{ Symbol }}` element. The format is either `{{ Admin }}` if
  the UID is 0 or `{{ User }}` if the UID is not 0.

- `GBT_CAR_SIGN_ADMIN_BG`

  Background color of the `{{ Admin }}` element.

- `GBT_CAR_SIGN_ADMIN_FG='red'`

  Foreground color of the `{{ Admin }}` element.

- `GBT_CAR_SIGN_ADMIN_FM`

  Formatting of the `{{ Admin }}` element.

- `GBT_CAR_SIGN_ADMIN_TEXT='#'`

  Text content of the `{{ Admin }}` element.

- `GBT_CAR_SIGN_USER_BG`

  Background color of the `{{ User }}` element.

- `GBT_CAR_SIGN_USER_FG='light_green'`

  Foreground color of the `{{ User }}` element.

- `GBT_CAR_SIGN_USER_FM`

  Formatting of the `{{ User }}` element.

- `GBT_CAR_SIGN_USER_TEXT='$'`

  Text content of the `{{ User }}` element. The user name.

- `GBT_CAR_SIGN_DISPLAY='1'`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_SIGN_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_SIGN_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_SIGN_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_SIGN_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_SIGN_SEP_FM`

  Formatting of the separator for this car.


#### `Status` car

Car that visualizes return code of every command. By default, this car is
displayed only when the return code is non-zero. If you want to display it even
if the return code is zero, set the following variable:

```shell
export GBT_CAR_STATUS_DISPLAY='1'
```

Variables used by the car:

- `GBT_CAR_STATUS_BG`

  Background color of the car. It's either `GBT_CAR_STATUS_OK_BG` if the last
  command returned `0` return code otherwise the `GBT_CAR_STATUS_ERROR_BG` is
  used.

- `GBT_CAR_STATUS_FG='default'`

  Foreground color of the car. It's either `GBT_CAR_STATUS_OK_FG` if the last
  command returned `0` return code otherwise the `GBT_CAR_STATUS_ERROR_FG` is
  used.

- `GBT_CAR_STATUS_FM='none'`

  Formatting of the car. It's either `GBT_CAR_STATUS_OK_FM` if the last command
  returned `0` return code otherwise the `GBT_CAR_STATUS_ERROR_FM` is used.

- `GBT_CAR_STATUS_FORMAT=' {{ Symbol }} '`

  Format of the car. This can be changed to contain also the value of the
  return code:

  ```shell
  export GBT_CAR_STATUS_FORMAT=' {{ Symbol }} {{ Code }} '
  ```

  or the signal name of the return code:

  ```shell
  export GBT_CAR_STATUS_FORMAT=' {{ Symbol }} {{ Signal }} '
  ```

  If you want to display the Status train even if there is no error, you have
  to use the `{{ Details }}` element to prevent the `{{ Code }}` and/or
  `{{ Signal }}` from being displayed:

  ```shell
  export GBT_CAR_STATUS_DISPLAY=1
  export GBT_CAR_STATUS_FORMAT=' {{ Symbol }}{{ Details }} '
  ```

  Then you can modify the format of the `{{ Details }}` element like this for
  when there is an error:

  ```shell
  export GBT_CAR_STATUS_DETAILS_FORMAT=' {{ Code }} {{ Signal }}'
  ```

- `GBT_CAR_STATUS_SYMBOL_BG`

  Background color of the `{{ Symbol }}` element.

- `GBT_CAR_STATUS_SYMBOL_FG`

  Foreground color of the `{{ Symbol }}` element.

- `GBT_CAR_STATUS_SYMBOL_FM='bold'`

  Formatting of the `{{ Symbol }}` element.

- `GBT_CAR_STATUS_SYMBOL_FORMAT`

  Format of the `{{ Symbol }}` element. The format is either `{{ Error }}` if
  the last command returned non zero return code otherwise `{{ User }}` is
  used.

- `GBT_CAR_STATUS_SIGNAL_BG`

  Background color of the `{{ Signal }}` element.

- `GBT_CAR_STATUS_SIGNAL_FG`

  Foreground color of the `{{ Signal }}` element.

- `GBT_CAR_STATUS_SIGNAL_FM`

  Formatting color of the `{{ Signal }}` element.

- `GBT_CAR_STATUS_SIGNAL_TEXT`

  Text of the `{{ Signal }}` element.

- `GBT_CAR_STATUS_CODE_BG='red'`

  Background color of the `{{ Code }}` element.

- `GBT_CAR_STATUS_CODE_FG='light_gray'`

  Foreground color of the `{{ Code }}` element.

- `GBT_CAR_STATUS_CODE_FM='none'`

  Formatting of the `{{ Code }}` element.

- `GBT_CAR_STATUS_CODE_TEXT`

  Text content of the `{{ Code }}` element. The return code.

- `GBT_CAR_STATUS_ERROR_BG='red'`

  Background color of the `{{ Error }}` element.

- `GBT_CAR_STATUS_ERROR_FG='light_gray'`

  Foreground color of the `{{ Error }}` element.

- `GBT_CAR_STATUS_ERROR_FM='none'`

  Formatting of the `{{ Error }}` element.

- `GBT_CAR_STATUS_ERROR_TEXT='✘'`

  Text content of the `{{ Error }}` element.

- `GBT_CAR_STATUS_OK_BG='green'`

  Background color of the `{{ Ok }}` element.

- `GBT_CAR_STATUS_OK_FG='light_gray'`

  Foreground color of the `{{ Ok }}` element.

- `GBT_CAR_STATUS_OK_FM='none'`

  Formatting of the `{{ Ok }}` element.

- `GBT_CAR_STATUS_OK_TEXT='✔'`

  Text content of the `{{ Ok }}` element.

- `GBT_CAR_STATUS_DISPLAY`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_STATUS_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_STATUS_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_STATUS_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_STATUS_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_STATUS_SEP_FM`

  Formatting of the separator for this car.


#### `Time` car

Car that displays current date and time.

- `GBT_CAR_TIME_BG='light_blue'`

  Background color of the car.

- `GBT_CAR_TIME_FG='light_gray'`

  Foreground color of the car.

- `GBT_CAR_TIME_FM='none'`

  Formatting of the car.

- `GBT_CAR_TIME_FORMAT=' {{ DateTime }} '`

  Format of the car.

- `GBT_CAR_TIME_DATETIME_BG`

  Background color of the `{{ DateTime }}` element.

- `GBT_CAR_TIME_DATETIME_FG`

  Foreground color of the `{{ DateTime }}` element.

- `GBT_CAR_TIME_DATETIME_FM`

  Formatting of the `{{ DateTime }}` element.

- `GBT_CAR_TIME_DATETIME_FORMAT='{{ Date }} {{ Time }}'`

  Format of the `{{ DateTime }}` element.

- `GBT_CAR_TIME_DATE_BG`

  Background color of the `{{ Date }}` element.

- `GBT_CAR_TIME_DATE_FG`

  Foreground color of the `{{ Date }}` element.

- `GBT_CAR_TIME_DATE_FM`

  Formatting of the `{{ Date }}` element.

- `GBT_CAR_TIME_DATE_FORMAT='Mon 02 Jan'`

  Format of the `{{ Date }}` element. The format is using placeholders as
  described in the [`time.Format()`](https://golang.org/src/time/format.go#L87)
  Go function. For example `January` is a placeholder for current full month
  name and `PM` is a placeholder `AM` if the current time is before noon or
  `PM` if the current time is after noon. So in order to display date in the
  format of `YYYY-MM-DD`, the value of this variable should be `2006-01-02`.

- `GBT_CAR_TIME_TIME_BG`

  Background color of the `{{ Host }}` element.

- `GBT_CAR_TIME_TIME_FG='light_yellow'`

  Foreground color of the `{{ Host }}` element.

- `GBT_CAR_TIME_TIME_FM`

  Formatting of the `{{ Host }}` element.

- `GBT_CAR_TIME_TIME_FORMAT='15:04:05'`

  Text content of the `{{ Host }}` element. The format principles are the same
  like in the case of the `GBT_CAR_TIME_DATE_FORMAT` variable above. So in
  order to display time in the 12h format, the value of this variable should be
  `03:04:05 PM`.

- `GBT_CAR_TIME_DISPLAY='1'`

  Whether to display this car if it's in the list of cars (`GBT_CARS`).

- `GBT_CAR_TIME_WRAP='0'`

  Whether to wrap the prompt line in front of this car.

- `GBT_CAR_TIME_SEP_TEXT`

  Text content of the separator for this car.

- `GBT_CAR_TIME_SEP_BG`

  Background color of the separator for this car.

- `GBT_CAR_TIME_SEP_FG`

  Foreground color of the separator for this car.

- `GBT_CAR_TIME_SEP_FM`

  Formatting of the separator for this car.


Benchmark
---------

Benchmark of GBT can be done by faking the output of GBT by a testing script
which executes as minimum of commands as possible. For simplicity, the test will
produce output of the Git car only and will be done from within a directory with
a Git repository.

The testing script is using exactly the same commands like GBT to determine the
Git branch, whether the Git repository contains any change and whether it's
ahead/behind of the remote branch. The script has the following content and is
stored in `/tmp/test.sh`:

```shell
BRANCH="$(git symbolic-ref HEAD)"
[ -z "$(git status --porcelain)" ] && DIRTY_ICON='%{\e[38;5;2m%}✔' || DIRTY_ICON='%{\e[38;5;1m%}✘'
[[ "$(git rev-list --count HEAD..@{upstream})" == '0' ]] && AHEAD_ICON='' || AHEAD_ICON=' ⬆'
[[ "$(git rev-list --count @{upstream}..HEAD)" == '0' ]] && BEHIND_ICON='' || BEHIND_ICON=' ⬇'

echo -en "%{\e[0m%}%{\e[48;5;7m%}%{\e[38;5;0m%} %{\e[48;5;7m%}%{\e[38;5;0m%}%{\e[48;5;7m%}%{\e[38;5;0m%} %{\e[48;5;7m%}%{\e[38;5;0m%}${BRANCH##*/}%{\e[48;5;7m%}%{\e[38;5;0m%} %{\e[48;5;7m%}%{\e[38;5;0m%}%{\e[48;5;7m%}$DIRTY_ICON%{\e[48;5;7m%}%{\e[38;5;0m%}%{\e[48;5;7m%}%{\e[38;5;0m%}%{\e[48;5;7m%}%{\e[38;5;0m%}$AHEAD_ICON%{\e[48;5;7m%}%{\e[38;5;0m%}%{\e[48;5;7m%}%{\e[38;5;0m%}$BEHIND_ICON%{\e[48;5;7m%}%{\e[38;5;0m%} %{\e[0m%}"
```

The testing script produces the same output like GBT when run by Bash or ZSH:

```shell
bash /tmp/test.sh > /tmp/a
zsh /tmp/test.sh > /tmp/b
GBT_SHELL='zsh' GBT_CARS='Git' gbt > /tmp/c
diff /tmp/{a,b}
diff /tmp/{b,c}
```

We will use ZSH to run 10 measurements of 100 executions of the testing script
by Bash and ZSH as well as of GBT itself.

```shell
# Execution of the testing script by Bash
for N in $(seq 10); do time (for M in $(seq 100); do bash /tmp/test.sh 1>/dev/null 2>&1; done;) done 2>&1 | sed 's/.*  //'
0.95s user 1.05s system 102% cpu 1.944 total
0.94s user 1.06s system 102% cpu 1.944 total
0.93s user 1.05s system 102% cpu 1.930 total
0.91s user 1.10s system 102% cpu 1.954 total
0.92s user 1.07s system 102% cpu 1.933 total
0.97s user 1.03s system 102% cpu 1.943 total
0.92s user 1.07s system 102% cpu 1.931 total
0.92s user 1.08s system 102% cpu 1.949 total
0.89s user 1.11s system 102% cpu 1.938 total
0.93s user 1.07s system 102% cpu 1.944 total
# Execution of the testing script by ZSH
for N in $(seq 10); do time (for M in $(seq 100); do zsh /tmp/test.sh 1>/dev/null 2>&1; done;) done 2>&1 | sed 's/.*  //'
0.89s user 1.08s system 103% cpu 1.909 total
0.82s user 1.15s system 103% cpu 1.906 total
0.82s user 1.15s system 103% cpu 1.903 total
0.84s user 1.13s system 103% cpu 1.907 total
0.88s user 1.10s system 103% cpu 1.915 total
0.88s user 1.09s system 103% cpu 1.907 total
0.84s user 1.14s system 103% cpu 1.919 total
0.85s user 1.11s system 103% cpu 1.901 total
0.89s user 1.08s system 103% cpu 1.914 total
0.96s user 1.01s system 103% cpu 1.908 total
# Execution of GBT
for N in $(seq 10); do time (for M in $(seq 100); do GBT_SHELL='zsh' GBT_CARS='Git' gbt 1>/dev/null 2>&1; done;) done 2>&1 | sed 's/.*  //'
1.03s user 1.19s system 115% cpu 1.922 total
0.98s user 1.18s system 115% cpu 1.874 total
1.06s user 1.11s system 115% cpu 1.880 total
1.02s user 1.14s system 115% cpu 1.867 total
1.04s user 1.17s system 115% cpu 1.918 total
1.05s user 1.10s system 115% cpu 1.853 total
1.07s user 1.11s system 115% cpu 1.895 total
1.01s user 1.18s system 115% cpu 1.903 total
1.08s user 1.03s system 115% cpu 1.825 total
1.05s user 1.09s system 115% cpu 1.844 total
```

From the above is visible that GBT performs faster than Bash and ZSH even if the
testing script was as simple as possible. You can also notice that GBT was using
more CPU than Bash or ZSH. That's probably because of the built-in concurrency
support in Go.


Prompt forwarding
-----------------

In order to enjoy GBT prompt via SSH but also in Docker, Kubectl, Vagrant, MySQL
or in Screen without the need to install GBT everywhere, you can use GBTS (GBT
written in Shell). GBTS is a set of scripts which get forwarded to applications
and remote connections and then executed to generate the nice looking prompt.

You can start using it by doing the following:

```shell
export GBT__HOME='/usr/share/gbt'
source $GBT__HOME/sources/gbts/cmd/local.sh
```

This will automatically create command line aliases for all enabled plugins (by
default `docker`, `gssh`, `kubectl`, `mysql`, `screen`, `ssh`, `su`, `sudo` and
`vagrant`). Then just SSH to some remote server or enter some Docker container
(even via `kubectl`) or Vagrant box and you should get GBT prompt there.

If you want to have some of the default aliase available only on the remote
site, just un-alias them locally:

```shell
unalias sudo su
```

You can also forward your own aliases which will be then available on any remote
site. For example to have `alias ll='ls -l'` on any remote site, just create the
following alias and it will be automatically forwarded:

```shell
alias gbt___ll='ls -l'
```

The idea behind prompt forwarding is coming from Vladimir Babichev
(@[mrdrup](https://github.com/mrdrup)) who was using it for several years
before GBT even existed. After seeing the potential of GBT, he sparked the
implementation of prompt forwarding into GBT which later turned into GBTS.


### Principle

Principle of GBTS is to pass the GBTS scripts to the application and then execute
them. This is done by concatting all the GBTS scripts into one file and encoding
it by Base64 algorithm. Such string, together with few more commands, is then
used as an argument of the application which makes it to store it on the remote
site in the `/tmp/.gbt.<NUM>` file. The same we create the `/tmp/.gbt.<NUM>.bash`
script which is then used as a replacement of the real shell on the remote site.
For SSH it would look like this:

```shell
ssh -t myserver "export GBT__CONF='$GBT__CONF' && echo '<BASE64_ENCODED_STRING>' | base64 -d > \$GBT__CONF && bash --rcfile \$GBT__CONF"
```

In order to make all this invisible, we wrap that command into a function (e.g.
`gbt_ssh`) and assign it to an `alias` of the same name like the original
application (e.g. `ssh`):

```shell
alias ssh='gbt_ssh'
```

The same or very similar principle applies to other supported commands like
`docker`, `gssh` ([GCP
SSH](https://cloud.google.com/sdk/gcloud/reference/compute/ssh)), `kubectl`,
`mysql`, `screen`, `su`, `sudo` and `vagrant`.


### Additional settings

GBTS has few settings which can be used to influence its behaviour. See the
details [here](https://github.com/jtyr/gbt/tree/master/sources/gbts/README.md).


### MacOS users

To make GBTS working correctly between Linux and MacOS and vice versa requires a
little bit of fiddling. The reason is that the basic command line tools like
`date` and `base64` are very old on MacOS and mostly incompatible with the Linux
world. Some tools are even called differently (e.g. `md5sum` is called `md5`).

Therefore if you want to make the remote script verification working (make sure
nobody changed the remote script while using it), the following variables must be
set:

```shell
# Use 'md5' command instead of 'md5sum'
export GBT__SOURCE_MD5_LOCAL='md5'
# Cut the 4th field from the output of 'md5'
export GBT__SOURCE_MD5_CUT_LOCAL='4'
```

If you don't want to use this feature, you can disable it in which case the above
variables won't be required:

```shell
export GBT__SOURCE_SEC_DISABLE=1
```

When using the `ExecTime` car, the following variable must be set:

```shell
# Don't use nanoseconds in the 'ExecTime' car
export GBT__SOURCE_DATE_ARG='+%s'
```

For maximum compatibility with GBT, it's recommended to install GNU `coreutils`
(`brew install coreutils`) and instead of the variable above use these:

```shell
# Use 'gdate' instead of 'date'
export GBT__SOURCE_DATE='gdate'
# Use 'gdate' instead of 'date' (only if you run GBT on a Mac)
export GBT__SOURCE_BASE64_LOCAL='gbase64'
# Use 'gdate' instead of 'date' (only if you are connection to Mac via SSH)
export GBT__SOURCE_BASE64='gbase64'
```

When connecting to MacOS from Linux using `gbt_ssh` and not using `gbase64` on
MacOS, the following variable must be set on Linux to make the Base64 decoding
working on MacOS:

```shell
# Use 'base64 -D' to decode Base64 encoded text
export GBT__SOURCE_BASE64_DEC='-D'
```


### Limitations

- Requires Bash v4.x to run.
- The [color representation](https://bugs.mysql.com/79755) and [support of
  unicode characters](https://bugs.mysql.com/89359) for MySQL is broken in MySQL
  5.6 and above. But it works just fine in all versions of Percona and MariaDB.
- Plugins `su` and `sudo` are not supported on MacOS.


TODO
----

Contribution to the following is more than welcome:

- Optimize generated escape sequence
    - Don't decorate empty string
    - Don't decorate child element with the same attributes used by the parent
- Implement templating language to allow more dynamic configuration
    - Jinja2-like syntax
    - Should be able to refer variables from the local car
        - `GBT_CAR_GIT_BG="{% 'red' if Status == '{{ StatusDirty }}' else 'light_gray' %}"`
    - Should be able to refer ENV variables (e.g. `env('HOME')`)
    - Could be able to refer variables from another car
    - Advanced functionality via pipes (e.g. `<expr> | substr(1,3)`)
- Add support for GBT [plugins](https://golang.org/pkg/plugin/)
    - Load plugins with `GBT_PLUGINS='mycar1:/path/to/mycar1.so;mycar2:/path/to/mycar2.so'`
    - Load the plugin, read the `Car` symbol and assign the reference to the
      `mycar1` in the `carsFactory`
- Implement Vim statusline using GBT as the generator
- Implement Tmux statusline using GBT as the generator
- Add weather car
    - Using Yahoo Weather API
    - Needs to cache the results in a file and refresh only if its timestamp is
      older than certain time. Or perhaps store the last update in env var?
- Add more themes


Author
------

Jiri Tyr


License
-------

MIT
