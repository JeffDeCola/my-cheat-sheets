# SYSTEMVERILOG CHEAT SHEET

```text
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

_Verilog is a Hardware Description Language (HDL) used to describe a digital system._

Table of Contents,

* [OVERVIEW AND HDL LEVELS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#overview-and-hdl-levels)
* [SAMPLE CODE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#sample-code)
* [DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#data-types)
* [OPERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#operators)
* [CONTROL STATEMENTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#control-statements)
* [VARIABLE ASSIGNMENT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#variable-assignment)
  * [INITIAL BLOCKS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#initial-blocks)
  * [ALWAYS BLOCKS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-blocks)
  * [ASSIGN STATEMENT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#assign-statement)
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

## SAMPLE CODE

Let's look at a block box,

```verilog
// Description of this Module
module name (
    in1,        // 1-bit input
    in2[3:0],   // 4-bit input  (little-endian)
    out1,       // 1-bit output
    out2[0:1],  // 2-bit output  (big-endian)
    clk,        // Clock
    rst         // Reset
);

// INPUT/OUTPUT
input in1, in2;
input clk, rst;
output out1, out2;

// DATA TYPE

```

## DATA TYPES

```verilog
// DRIVER
    reg                 // Holds a state
    wire                // Connecting things
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

* **COMBINATIONAL ELEMENTS** Two ways to model - **assign** and **always** statements
* **SEQUENTIAL ELEMENTS**  One way to model - **always** statement
* **INITIAL STATEMENTS** Used in test benches only

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

### ALWAYS BLOCKS

An always block,

* Executes always, unlike initial blocks which execute only once
* Should have a sensitive list or a delay associated with it

```verilog
always  @ (a or b or sel)
begin
  y = 0;
  if (sel == 0) begin
    y = a;
  end else begin
    y = b;
  end
end
```

### ASSIGN STATEMENT

An assign statement,

* Used to model combinational logic

```verilog
assign out = (enable) ? data : 1'bz;
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
