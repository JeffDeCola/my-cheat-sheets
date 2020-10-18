# SYSTEMVERILOG CHEAT SHEET

```text
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

_Verilog is a Hardware Description Language (HDL) used to describe a digital system._

Table of Contents,

* [OVERVIEW AND HDL LEVELS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#overview-and-hdl-levels)
* [BASIC STRUCTURE - THE MODULE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#basic-structure---the-module)
* [DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#data-types)
* [OPERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#operators)
* [CONTROL STATEMENTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#control-statements)
* [VARIABLE ASSIGNMENT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#variable-assignment)
  * [ASSIGN STATEMENT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#assign-statement)
  * [ALWAYS BLOCKS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-blocks)
  * [INITIAL BLOCKS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#initial-blocks)
* [TASK AND FUNCTION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#task-and-function)
* [GATE PRIMITIVES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#gate-primitives)

Documentation and reference,

* [iverilog](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet)
  is a free verilog simulator and synthesis tool
* [GTKWave](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/gtkwave-cheat-sheet)
  is a free waveform viewer
* VS Code [Verilog HDL Extension](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/visual-studio-code-cheat-sheet/verilog-hdl-extension.md)
 for syntax highlighting and linting
* Check out my repo [my-systemverilog-examples](https://github.com/JeffDeCola/my-systemverilog-examples)

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## OVERVIEW AND HDL LEVELS

Verilog is used to describe digital system at the,

* Behavioral Level - Good for testing
* Register Transfer Level (RTL)
* Gate Level
* Switch Level

The word Verilog is a combination of the words `verification` and `logic`
because the language was first suggested as a simulation and verification tool.

## BASIC STRUCTURE - THE MODULE

Let's look at a black box of a ,

```verilog
// Description of this Module
module name (
    input in1,        // 1-bit input
    input in2[3:0],   // 4-bit input  (little-endian)
    intpu out1,       // 1-bit output
    out2[0:1],  // 2-bit output  (big-endian)
    clk,        // Clock
    rst         // Reset
);

// PORT DECLARATION
input in1, in2;
input clk, rst;
output out1, out2;

// DATA TYPES

// CODE STARTS HERE


```

## DATA TYPES

```verilog
// DRIVER
    reg                 // Holds a state
    wire                // Connecting things - represents a physical wire
```

## OPERATORS

```verilog
// ARITHMETIC
    *                   // Multiply
    /                   // Division
    +                   // Add
    -                   // Subtract
    %                   // Modulus
    +                   // Unary plus
    -                   // Unary minus
// LOGICAL
    !                   // Logical Not
    &&                  // Logical And
    ||                  // Logical Nor
// RELATIONAL
    tbd                 // tbd
// EQUALITY
    tbd                 // tbd
// REDUCTION
    tbd                 // tbd
// SHIFT
    tbd                 // tbd
// CONCATENATION
    tbd                 // tbd
// CONDITIONAL
    tbd                 // tbd
```

## CONTROL STATEMENTS

```verilog
// IF-ELSE
    if (enable == 1'b1) begin
        //stuff
    end else begin
        // more stuff
    end
// CASE
    tbd                 // tbd
// WHILE
    tbd                 // tbd
// FOR LOOP
    tbd                 // tbd
// REPEAT
    tbd                 // tbd
```

## VARIABLE ASSIGNMENT

* **COMBINATIONAL ELEMENTS** Two ways to model - **assign statement** and
  **always block**
* **SEQUENTIAL ELEMENTS**  One way to model - **always block**
* **INITIAL STATEMENTS** Used in test benches only

### ASSIGN STATEMENT

An assign statement,

* Used to model combinational logic
* Only use non-blocking statements

```verilog
assign b <= a;      // another way to write it
b = a;              // blocking
b <= a;             // non-blocking
```

### ALWAYS BLOCKS

An always block,

* Executes always, unlike initial blocks which execute only once
* Should have a sensitive list or a delay associated with it
* Used for both combinational and sequential logic

```verilog
always @ (posedge clk)
begin
    b <= a;
end
```

### INITIAL BLOCKS

An initial block is executed only once when simulation starts.

```verilog
// INITIAL BLOCK
initial begin
    clk = 0;
    reset = 0;
    req_0 = 0;
end
```

## TASK AND FUNCTION

Just like other languages when repeating the same old things again and again.

* Tasks can have a delay.
* Functions can return a value, whereas tasks can not.

```verilog
function parity;
input [31:0] data;
integer i;
begin
  parity = 0;
  for (i= 0; i < 32; i = i + 1) begin
    parity = parity ^ data[i];
  end
end
endfunction
```

## GATE PRIMITIVES

You can model gate primitives at the gate level with,

```verilog
    and(xy, x, y);
    nand(xy, x, y);
    or();
    nor();
    or();
    nor();
```
