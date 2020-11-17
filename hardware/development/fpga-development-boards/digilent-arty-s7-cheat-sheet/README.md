# DIGILENT ARTY-S7 CHEAT SHEET

_A FPGA development board.  Just some of my notes on this board._

Table of Contents,

* [SOME SPECS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/fpga-development-boards/digilent-arty-s7-cheat-sheet#some-specs)
* [PROGRAM ONBOARD FLASH WITH YOUR DESIGN](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/fpga-development-boards/digilent-arty-s7-cheat-sheet#program-onboard-flash-with-your-design)
* [INPUT & OUTPUT USING PMOD PINS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/fpga-development-boards/digilent-arty-s7-cheat-sheet#input--output-using-pmod-pins)

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

Some specs,

* 4 Pmod ports
  * 2 high speed JA, JB Pmod ports
  * 2 Standard JC, JD Pmod ports
  * 8 usable FPGA signals on each Pmod
  * 32 total FPGA I/O (8x24)
* Each 12-pin Pmod connector provides
  * 2 3.3V VCC signals (pins 6 and 12)
  * 2 Ground signals (pins 5 and 11)
  * 8 logic signals (remainder)

To access the pins use
[Xilinx Vivado](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/synthesis/xilinx-vivado-cheat-sheet).
More specifically, the constraint file.

This is how I set up the led and button on a breadboard,

![IMAGE - arty-s7-input-and-output-using-pmod - IMAGE](../../../../docs/pics/arty-s7-input-and-output-using-pmod.jpg)

This is the Pmod pin map,

![IMAGE - pmod-connector- IMAGE](../../../../docs/pics/pmod-connector.png)

|             | Pmod JA    | Pmod JB    | Pmod JC    | Pmod JD    |
|-------------|------------|------------|------------|------------|
| Pmod Type   | High-Speed | High-Speed | Standard   | Standard   |
| Shared pins | -          | -          | IO34-IO41  | IO26-IO33  |
| Pin 1       | L17        | P17        | U15        | V15        |
| Pin 2       | L18        | P18        | V16        | U12        |
| Pin 3       | M14        | R18        | U17        | V13        |
| Pin 4       | N14        | T18        | U18        | T12        |
| Pin 7       | M16        | P14        | U16        | T13        |
| Pin 8       | M17        | P15        | P13        | R11        |
| Pin 9       | M18        | N15        | R13        | T11        |
| Pin 10      | N18        | P16        | V14        | U11        |
