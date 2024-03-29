# VS CODE VERILOG HDL EXTENSION CHEAT SHEET

_A handy extension for syntax highlighting, formatting and linting._

Documentation and Reference

* [VS Code marketplace](https://marketplace.visualstudio.com/items?itemName=mshr-h.VerilogHDL&ssr=false#overview)
* My cheat sheet on
  [verilog](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/verilog-cheat-sheet)
  
## INSTALL

In order to use this extension you must install a `simulator` and `ctags`.

### INSTALL SIMULATOR

Instruction to
[install iverilog](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet)

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
