# IVERILOG CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_iverilog (icaris verilog) is a free tool for simulation and synthesis._

Table of Contents

* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet#install)
  * [INSTALL ON LINUX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet#install-on-linux)
  * [INSTALL ON WINDOWS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet#install-on-windows)
  * [INSTALL ON macOS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet#install-on-macos)
* [CHECK INSTALLATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet#check-installation)
* [SIMPLE EXAMPLE (NO TESTBENCH OR WAVEFORM FILE)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet#simple-example-no-testbench-or-waveform-file)
* [SIMPLE EXAMPLE WITH TESTBENCH AND WAVEFORM .VCD FILE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet#simple-example-with-testbench-and-waveform-vcd-file)

Documentation and Reference

* icarus verilog
  [home page](https://steveicarus.github.io/iverilog/)
* iverilog
  [repo](  https://github.com/steveicarus/iverilog)
* iverilog
  [Installation guide](https://iverilog.fandom.com/wiki/Installation_Guide)
* [verilog](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/verilog-cheat-sheet)
  cheat sheet
* [my-verilog-examples](https://github.com/JeffDeCola/my-verilog-examples)
* [GTKWave](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/gtkwave-cheat-sheet)
  is a free HDL waveform viewer
* [xilinx vivado](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/synthesis/xilinx-vivado-cheat-sheet)
  is a complete IDE for synthesis and analysis of HDL designs
* [digilent ARTY-S7](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/fpga-development-boards/digilent-arty-s7-cheat-sheet)
  is a FPGA development board
  
## INSTALL

### INSTALL ON LINUX

You can install it either from a package or build it from source.

#### Install from a package

Gets placed in `/usr/bin`.

This is easier,

```bash
sudo apt-get update
sudo apt-get install verilog
```

#### Install from Source

Gets placed in `/usr/local/bin` (default).

Goto
[github.com/steveicarus/iverilog](https://github.com/steveicarus/iverilog)
for latest information.

I compiled from source to /usr/local (default),

```bash
git clone git@github.com:steveicarus/iverilog.git
```

I needed a few things,

```bash
sudo apt-get install -y autoconf
sudo apt-get install -y gperf
sudo apt-get install -y flex
sudo apt-get install -y bison
```

Build configuration files,

```bash
cd iverilog
sh autoconf.sh
```

Now lets compile your source,

```bash
./configure
make
sudo su
make install
```

### INSTALL ON WINDOWS

Pre-built binaries are
[here](http://bleyer.org/icarus/)

### INSTALL ON macOS

The GNU Bison tool (packaged with Xcode) needs to be updated to version 3.

```bash
brew install bison
```

Install iverilog,

```bash
brew install icarus-verilog
```

## CHECK INSTALLATION

Check,

```bash
iverilog -h
```

## SIMPLE EXAMPLE (NO TESTBENCH OR WAVEFORM FILE)

Create a verilog file `hello.v`,

```verilog
module main();

initial
  begin
    $display("Hi there");
    $finish ;
  end

endmodule
```

Compile,

```bash
iverilog -o hello hello.v
```

Execute with linux,

```bash
./hello
```

Execute with Windows,

```bash
vvp hello
```

## SIMPLE EXAMPLE WITH TESTBENCH AND WAVEFORM .VCD FILE

In this example we will create, synthesis and test an AND gate.
The flow is as follows,

![IMAGE - iVERILOG to .vcd FLOW - IMAGE](../../../../docs/pics/hardware/iverilog-to-vcd-flow.svg)

First, create your AND gate verilog file `and-gate.v`,

```verilog
module and_gate(
    A,
    B,
    Y
);

    input A, B;
    output Y;

    and(Y, A ,B);

endmodule
```

Next create your testbench verilog file `and-gate-tb.v`,

```verilog
`timescale 1ns / 1ns
`include "and-gate.v"

module and_gate_tb;

reg A, B;
wire Y;

and_gate uut(A, B, Y);

initial begin

    $dumpfile("and-gate-tb.vcd");
    $dumpvars(0, and_gate_tb);

    A = 0;
    B = 0;
    #20;

    A = 1;
    B = 0;
    #20;

    A = 0;
    B = 1;
    #20;

    A = 1;
    B = 1;
    #20;

    $display("test complete");

end

endmodule
```

Now lets compile and simulate,

Compile,

```bash
iverilog -o and-gate-tb.vvp and-gate-tb.v
```

Execute with linux,

```bash
./and-gate-tb.vvp
```

Execute with Windows,

```bash
vvp and-gate-tb.vvp
```

You will now have a `and-gate.vcd` file you can use with a waveform viewer
such as
[GTKWave](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/gtkwave-cheat-sheet).
