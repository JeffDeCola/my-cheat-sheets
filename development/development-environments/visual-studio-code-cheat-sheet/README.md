# VISUAL STUDIO CODE CHEAT SHEET

`visual studio code` _is a code development environment._

this cheat-sheet is heavy on windows and go, and I probably
should of called it `Vs Code windows using go cheat sheet`.

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## INSTALL ON LINUX OR WINDOWS

Most, if not all of the time VS Code will be
install on Linux.  But a few crazy people like
to use Windows as a Dev Environment.

Go to the visual studio website for how to install,

[code.visualstudio.com](https://code.visualstudio.com/)

## CONFIGURE - LINUX

This is pretty straight forward and not sure I
could add anything useful here.

## CONFIGURE - WINDOWS VERSION OF VS CODE

Visual Studio Code on Windows will ALWAYS run
on Windows.  Hence, you want to edit your
files in Windows realm, not linux.

First, in VS Code user settings, set VS Code to point to WSL bash,

```yaml
"terminal.integrated.shell.windows": "C:\\Windows\\sysnative\\bash.exe",
```

You now can use your linux bash shell in VS Code.

Now, we want to change where you edit your files.

As we know, bash home for user is `/home/<bashusername>`.

This directory is actually located in Windows here,

```text
C:\Users\<winusername>\AppData\Local\lxss\home\<bashusername>
```

DO NOT touch these files.  If you do, only edit using
a bash editor like nano.

So if you can't edit in this area, you need
to put your project files somewhere else.

I use the directory to keep my projects,

`/mnt/c/Users/<winusername>/home/<bashusername>`

Note, `/mnt/c ` in WSL is the exact same as `C:\ `,

So, when you start Windows bash, it puts you in /home/<bashusername>,

Bottom Line, you still have home as `/home/<bashusername>`,
but now you edit projects in `/mnt/c/Users/<winusername>/home/<bashusername>`.

I change my workspace in VS Code to this new area,

```text
C:\Users\<winusername>\home\<bashusername>
```

Your prompt should look like,

```text
<bashusername>@<computername>:/mnt/c/Users/<winusername>/home/<bashusername>
```

## USING GO with WINDOWS VS CODE and BASH WSL TERMINAL

That is a mouthful and the setup took me forever to
figure out because you actually need two versions of go.

Here's a diagram that might help explain what's going on,

![IMAGE - USING GO with WINDOWS VS CODE and BASH WSL TERMINAL - IMAGE](../../../docs/pics/using-go-with-windows-vs-code-and-bash-wsl-terminal.jpg)

You will need to install the following on windows,

* [go](https://golang.org/doc/install).
* [git for windows](https://git-scm.com/downloads)
* Obviously VS Code.

When you install go, it should set the paths as follows,

```bash
GOROOT=C:\Go\
GOPATH C:\Users\<WINDOWSNAME>\go
Path=...\Go\bin;...%GOPATH%\bin
```

Open a windows command prompt and type the following,

* `set` Check the go paths above.
* `go version`
* `git version`

You may need to create your workspace directory for go.

When you open VS Code, You install the go extension.

IMPORTANT - Now here's the trick, it will use the Windows
version of go (and git to install), NOT the go version in bash.

The go extensions will be placed in
`C:\Users\<WINDOWSNAME>\go\bin`

You're good to go.  Yep, that joke has never been used before. :)
