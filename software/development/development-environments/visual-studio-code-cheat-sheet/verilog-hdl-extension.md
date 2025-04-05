# VS CODE VERILOG HDL EXTENSION CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A handy extension for syntax highlighting, formatting and linting._

Table of Contents

* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/visual-studio-code-cheat-sheet/verilog-hdl-extension.md#install)
  * [INSTALL SIMULATOR](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/visual-studio-code-cheat-sheet/verilog-hdl-extension.md#install-simulator)
  * [INSTALL UNIVERSAL CTAGS](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/visual-studio-code-cheat-sheet/verilog-hdl-extension.md#install-universal-ctags)
  * [JSON USER SETTINGS IN VS CODE](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/visual-studio-code-cheat-sheet/verilog-hdl-extension.md#json-user-settings-in-vs-code)

Documentation and Reference

* [vs code](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/development-environments/visual-studio-code-cheat-sheet)
* [vs code marketplace](https://marketplace.visualstudio.com/items?itemName=mshr-h.VerilogHDL&ssr=false#overview)
* [verilog](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/verilog-cheat-sheet)

## INSTALL

In order to use this extension you must install a `simulator` and `ctags`.

### INSTALL SIMULATOR

Instruction to
[install iverilog](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet).

### INSTALL UNIVERSAL CTAGS

Extension help keeps updated list on how to install.

### JSON USER SETTINGS IN VS CODE

For Windows,

```json
    "verilog.linting.linter": "iverilog",
    "verilog.ctags.path": "C:\\ctags\\ctags.exe",
```

Fort Linux and macOS,

```json
    "verilog.linting.linter": "iverilog",
    "verilog.ctags.path": "/usr/local/bin/ctags",
```
