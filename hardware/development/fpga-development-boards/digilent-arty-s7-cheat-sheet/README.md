# DIGILENT ARTY-S7 CHEAT SHEET

_A FPGA development board.  Just some of my notes on this board._

Documentation and reference,

* [Digilent ARTY-S7](https://reference.digilentinc.com/reference/programmable-logic/arty-s7/start)
  home page
* [Xilinx Vivado](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/synthesis/xilinx-vivado-cheat-sheet)
* My verilog examples repo [my-systemverilog-examples](https://github.com/JeffDeCola/my-systemverilog-examples)

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## SOME SPECS

A very short list,

* Xilinx Spartan-7 FPGA (Part Number XC7S50-CSGA324)
* 52,160 logic cells
* 65,200 flip-flops
* 250 pins

## PROGRAM ONBOARD FLASH WITH YOUR DESIGN

Reprogram the onboard flash so it boots with your design.

* Open Xilinx Vivado and create your .bit and .bin file
  * Make sure in project settings you click checkbox to also create .bin file
* In Hardware manager choose `Add Configuration Memory Device`
  * Pick Spansion part number `s25fl128sxxxxxx0`
* Hardware Manager should program your dev board

## INPUT & OUTPUT USING PMOD PINS


This may help,

raspberry-pi-input-and-output-using-gpio-pins
![IMAGE - raspberry-pi-input-and-output-using-gpio-pins - IMAGE](????)

