# DIGILENT ARTY-S7 CHEAT SHEET

_A FPGA development board.  Just some of my notes on this board._

Documentation and reference,

* [Digilent ARTY-S7](https://reference.digilentinc.com/reference/programmable-logic/arty-s7/start)
  Home page
* [Xilinx Vivado](https://www.xilinx.com/products/design-tools/vivado.html)
  home page
* My verilog examples repo [my-systemverilog-examples](https://github.com/JeffDeCola/my-systemverilog-examples)

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## PROGRAM ONBOARD FLASH WITH YOUR DESIGN

Reprogram the onboard flash so it boots with your design.

* Open Xilinx Vivado and create your .bit and .bin file
  * Make sure in project settings you click checkbox to also create .bin file
* In Hardware manager choose `Add Configuration Memory Device`
  * Pick Spansion part number `s25fl128sxxxxxx0`
* It should program your deivce
