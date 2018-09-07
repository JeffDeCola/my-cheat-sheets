# VISUAL STUDIO CODE CHEAT SHEET

`visual studio code` _is a code development environment._

This cheat-sheet is heavy on Windows and go, and I probably
should of called it `VS Code for Windows using go`.

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## INSTALL ON LINUX OR WINDOWS

Most, if not all of the time VS Code will be
install on Linux.  But a few crazy people like
to use Windows as a Dev Environment.

Go to [code.visualstudio.com](https://code.visualstudio.com/)
for how to install.

## CONFIGURE - LINUX

This is pretty straight forward and not sure how I
could add anything else useful.

## CONFIGURE - WINDOWS VERSION OF VS CODE

Visual Studio Code on Windows will ALWAYS run
on Windows.  Hence, you want to edit your
files in Windows realm, not linux.

After you have installed VS Code, goto the user settings and set VS Code to point to WSL bash,

```yaml
"terminal.integrated.shell.windows": "C:\\Windows\\sysnative\\bash.exe",
```

You now can use your linux bash shell in VS Code.

Now, we want to change your working directory (where you edit your files).

As we know, bash home is `/home/<bashusername>`.

This directory is actually located in Windows here,

```
C:\Users\<winusername>\AppData\Local\lxss\home\<bashusername>
```

DO NOT touch these files.  If you do, only edit using
a bash editor like nano.

So if you can't edit in this area, you need
to put your project files somewhere else.

I use the directory to keep my projects,

`/mnt/c/Users/<winusername>/home/<bashusername>`

Note, `/mnt/c` in WSL is the exact same as `C:`,

So when you start Windows bash, it puts you in `/home/<bashusername>`,

Bottom Line: You still have home as `/home/<bashusername>`,
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

That is a mouthful and the setup took me a minute to
figure out because you actually need two versions of go.

Here's a diagram that might help explain what's going on,

![IMAGE - USING GO with WINDOWS VS CODE and BASH WSL TERMINAL - IMAGE](../../../docs/pics/using-go-with-windows-vs-code-and-bash-wsl-terminal.jpg)

You will need to install the following on windows,

* [go](https://golang.org/doc/install).
* [git for windows](https://git-scm.com/downloads)
* Obviously VS Code.

When you install go for Windows, it should set the paths as follows,

```bash
GOROOT=C:\Go\
GOPATH C:\Users\<WINDOWSNAME>\go
Path=...\Go\bin;...%GOPATH%\bin
```

To check everything, open a windows command prompt and
type the following,

* `set` Check the go paths above.
* `go version`
* `git version`

You may need to create your workspace directory for go.

When you open VS Code, install the go extension. IMPORTANT NOTE
- VS Code will use the Windows version of go (and git to install),
NOT the go version in bash.

The go extensions will automatically be placed in
`C:\Users\<WINDOWSNAME>\go\bin`

That's it, you're good to go.
