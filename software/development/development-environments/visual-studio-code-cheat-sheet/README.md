# VISUAL STUDIO CODE CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Visual studio code is a code development environment._

Table of Contents

* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet#install)
* [CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet#configure)
* [USING VS CODE ON WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet#using-vs-code-on-windows)
* [ADDING A WINDOWS SHORTCUT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet#adding-a-windows-shortcut)

Documentation and Reference

* [vs code verilog-HDL extension](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet/verilog-hdl-extension.md)
* [vs code on windows wsl](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet/vs-code-on-windows-wsl.md)

## INSTALL

Go to [code.visualstudio.com](https://code.visualstudio.com/)
for how to install.

## CONFIGURE

_Coming Soon._

## USING VS CODE ON WINDOWS

Running vs code on windows is a lot easier with wsl2.
To run with the older wsl refer to my cheat sheet
[vs-code-on-windows-wsl](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet/vs-code-on-windows-wsl.md).

## ADDING A WINDOWS SHORTCUT

For wsl2,

```text
"C:\Users\jeffry\AppData\Local\Programs\Microsoft VS Code\Code.exe" --remote wsl+Ubuntu-20.04 /home/jeff/jeffs.code-workspace
"C:\Users\jeffry\AppData\Local\Programs\Microsoft VS Code Insiders\Code - Insiders.exe" --remote wsl+Ubuntu-20.04 /home/jeff/jeffs.code-workspace
```

For ssh remote machines (e.g. 4g),

```text
"C:\Users\jeffry\AppData\Local\Programs\Microsoft VS Code\Code.exe" --file-uri "vscode-remote://ssh-remote+4b/home/jeff/jeffs.code-workspace"
"C:\Users\jeffry\AppData\Local\Programs\Microsoft VS Code Insiders\Code - Insiders.exe" --file-uri "vscode-remote://ssh-remote+4b/home/jeff/jeffs.code-workspace"
```
