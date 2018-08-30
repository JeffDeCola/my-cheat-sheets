# VISUAL STUDIO CODE CHEAT SHEET

`visual studio code` _is a code development environment._

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## INSTALL ON LINUX OR WINDOWS

Most, if not all of the time VS Code will be
install on Linux.  But a few crazy people like
to use Windows as a Dev Environment.

Go to the visual studio website for how to install,

[code.visualstudio.com](https://code.visualstudio.com/)

## CONFIGURE WINDOWS VERSION OF VSCODE

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

So if you can;t edit in this area, you need
to put your project somewhere else.

I use the directory to keep my projects,

`/mnt/c/Users/<winusername>/home/<bashusername>`

Note, `/mnt/c ` is the exact same as `C:\ `,

So, when you start bash, it puts you in /home/<bashusername>,

Bottom Line, you still have home as `/home/<bashusername>`,
but now you edit projects in `/mnt/c/Users/<winusername>/home/<bashusername>`.

I change my workspace in VS Code to this new area.

```text
C:\Users\<winusername>\home\<bashusername>
```

Your prompt should look like,

```text
<bashusername>@<computername>:/mnt/c/Users/<winusername>/home/<bashusername>`
```

If you use go, change the paths as follows, 

```bash
GOPATH=/mnt/c/Users/<winusername>/home/<bashusername>
GOROOT=/usr/local/go
PATH= ... /usr/local/go/bin ...
```