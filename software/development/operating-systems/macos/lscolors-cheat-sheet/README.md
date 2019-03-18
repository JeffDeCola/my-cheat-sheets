# LSCOLORS CHEAT SHEET

`LSCOLORS` _is an environment variable used to set different colors
for different kinds of files when using the `ls` command._

The `ls` that comes with macOS is from `BSD`, and lacks some
of the features of its GNU sibling.

Setting `ls` color for linux is
[here](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/ls_colors-cheat-sheet)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## BSD - SET LSCOLORS ENV VARIABLE

Its pretty simple,

```bash
export CLICOLOR=1
export LSCOLORS=ExGxBxDxCxEgEdxbxgxcxd
```

For grep add,

```bash
export GREP_OPTIONS='--color=always'
export GREP_COLOR='01;31'
```

A limitation is that you can't set different file
extensions like .gz.

## BSD - SETTINGS

Settings for LSCOLORS.

* FILE TYPES (Foreground, background)
  * directory = e, x
  * symbolic = f, x
  * socket = c, x
  * pipe = d, x
  * executable = b, x
  * block = e, g
  * character = e, d
  * executable = a, b
  * executable = a, g
  * directory = a, c
  * directory = a, d
* COLORS
  * a = Black
  * b = Red
  * c = Green
  * d = Brown
  * e = Blue
  * f = Magenta
  * g = Cyan
  * h = Light grey
  * A = Bold black, usually shows up as dark grey
  * B = Bold red
  * C = Bold green
  * D = Bold brown, usually shows up as yellow
  * E = Bold blue
  * F = Bold magenta
  * G = Bold cyan
  * H = Bold light grey; looks like bright white
  * x = Default foreground or background
