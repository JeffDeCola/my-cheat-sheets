# LS_COLORS CHEAT SHEET

`LS_COLORS` _is an environment variable used to set different colors
for different kinds of files when using the `ls` command._

Setting `ls` color for macOS is
[here](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/macos/lscolors-cheat-sheet)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## METHOD 1 - SET LS_COLORS ENV VARIABLE

This method sets the `LS_COLOR` env variable.

### CONFIGURE

To add color to `ls`, run the following,

```bash
alias ls='ls --color=auto'
LS_COLORS='di=01;35;100:fi=33:ex=93:*.monkey=31'
export LS_COLORS
```

Where the file types are,

* di = directory
* fi = file
* ex = file which is executable (i.e. has 'x' set in permissions).

And the color codes are,

* 31 = red
* 33 = orange
* 35 = purple
* 93 = yellow
* 100 = dark grey background

Simply put, `ls --colors=auto` tells the display files in different
colors based on the setting of the `LS_COLORS` env variable.

So with this example,

* Directories will be bold purple with a dark grey background
* Files will be orange
* Executables will be yellow
* `*.monkey` will be red

Do a `ls` and you should see your changes.

Check out the env variable,

```bash
echo $LS_COLORS
env | grep LS_COLORS
```

### FIX THE 755/755 FILE ISSUE

Sometimes a file could have a 777 permission. So the `ex`
will override a lot of settings.

This is easily fixed by setting `ex=0`.

Hence the above example would be,

```bash
alias ls='ls --color=auto'
LS_COLORS='di=01;35;100:fi=33:ex=0:*.monkey=31'
export LS_COLORS
```

### FIX BASH FOR UBUNTU FOR WINDOWS DIRECTORY COLOR ISSUE

Microsoft always has little issues,

For directories, set 'ow=01;35;100' which is `other-writable`.
Hence,

```bash
alias ls='ls --color=auto'
LS_COLORS='di=01;35;100:ow=01;35;100:fi=33:ex=0:*.monkey=31'
export LS_COLORS
```

### TO TURN THIS OFF

```bash
unalias ls
```

## METHOD 2 - USE DIRCOLORS TO SET LS_COLORS

Time to up our game. `dircolors` will allow us to set
the `LS_COLOR` env variable from a file.

First, lets create a nice file `~/.dircolors` to keep
our settings.

You can automatically generate this file using `dircolors`,
and edit as will.

```bash
dircolors --print-database > ~/.dircolors
```

Now update your `LS_COLORS` env variable,

```bash
eval "`dircolors -b ~/.dircolors`"
alias ls='ls --color=auto'
```

Check out the updated env variable,

```bash
echo $LS_COLORS
env | grep LS_COLORS
```

I modified this file with the few things explained above.

### MY SETTINGS IN .BASH

I use the following code in my `.bashrc` file,

```bash
test -r ~/.dircolors && eval "$(dircolors ~/.dircolors)"
alias ls='ls --color=auto'
alias grep='grep --color=auto'
alias fgrep='fgrep --color=auto'
alias egrep='egrep --color=auto'
```

Then I created a `.directories` file and updated
a few things as explained above.

## SETTINGS

Some settings for LS_COLORS.

* FILE TYPES
  * bd  = (BLOCK, BLK) Block device (buffered) special file
  * cd  = (CHAR, CHR) Character device (unbuffered) special file
  * di  = (DIR)  Directory
  * do  = (DOOR) Door
  * ex  = (EXEC) Executable file (ie. has 'x' set in permissions)
  * fi  = (FILE) Normal file
  * ln  = (SYMLINK, LINK, LNK)   Symbolic link. If you set this to
    'target' instead of a numerical value, the color is as for the file pointed to.
  * mi  = (MISSING)  Non-existent file pointed to by a symbolic link
    (visible when you type ls -l)
  * no  = (NORMAL, NORM) Normal (non-filename) text. Global default,
    although everything should be something
  * or  = (ORPHAN)   Symbolic link pointing to an orphaned non-existent file
  * ow  = (OTHER_WRITABLE)   Directory that is other-writable (o+w) and not
    sticky
  * pi  = (FIFO, PIPE)   Named pipe (fifo file)
  * sg  = (SETGID) File that is setgid (g+s)
  * so  = (SOCK) Socket file
  * st  = (STICKY) Directory with the sticky bit set (+t) and not other-writable
  * su  = (SETUID) File that is setuid (u+s)
  * tw  = (STICKY_OTHER_WRITABLE)    Directory that is sticky and other-writable
    (+t,o+w)
* EXTENSIONS
  * `*.extension` = Every file using this extension e.g. *.rpm = files with
    the ending .rpm
* COLORS
  * 31  = red
  * 32  = green
  * 33  = orange (yellow)
  * 34  = blue
  * 35  = purple
  * 36  = cyan
  * 37  = grey
  * 90  = dark grey
  * 91  = light red
  * 92  = light green
  * 93  = yellow
  * 94  = light blue
  * 95  = light purple
  * 96  = turquoise
  * 97  = white
* BACKGROUND COLORS
  * 40  = black background
  * 41  = red background
  * 42  = green background
  * 43  = orange background (yellow)
  * 44  = blue background
  * 45  = purple background
  * 46  = cyan background
  * 47  = grey background
  * 100 = dark grey background
  * 101 = light red background
  * 102 = light green background
  * 103 = yellow background
  * 104 = light blue background
  * 105 = light purple background
  * 106 = turquoise background
  * 107 = white background
* EFFECTS
  * 00  = default color
  * 01  = bold
  * 04  = underlined
  * 05  = flashing text (disabled on some terminals)
  * 07  = reverse field (exchange foreground and background color)
  * 08  = concealed (invisible)
