# SYSTEMVERILOG CHEAT SHEET

_Verilog is a Hardware Description Language (HDL) used to describe a digital system._

Table of Contents,

* [OVERVIEW & HDL LEVELS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#overview--hdl-levels)
* [BASIC SYNTAX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#basic-syntax)
  * [BASIC STRUCTURE (THE MODULE)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#basic-structure-the-module)
  * [NUMBERS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#numbers)
  * [DATA TYPES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#data-types)
  * [SCALAR, VECTOR & ARRAYS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#scalar-vector--arrays)
  * [OPERATORS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#operators)
* [MODELING COMBINATIONAL & SEQUENTIAL LOGIC (USING 3 BASIC BUILDING BLOCKS)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#modeling-combinational--sequential-logic-using-3-basic-building-blocks)
  * [ASSIGN STATEMENT](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#assign-statement)
  * [ALWAYS BLOCK (WHERE THE MAGIC HAPPENS)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-block-where-the-magic-happens)
  * [INITIAL BLOCK (TESTBENCH)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#initial-block-testbench)
* [MORE SYNTAX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#more-syntax)
  * [GATE PRIMITIVES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#gate-primitives)
  * [CONCATENATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#concatenation)
  * [CONTROL STATEMENTS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#control-statements)
  * [TASKS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#tasks)
  * [FUNCTIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#functions)

Documentation and reference,

* [iverilog](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/iverilog-cheat-sheet)
  is a free verilog simulator and synthesis tool
* [GTKWave](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/tools/simulation/gtkwave-cheat-sheet)
  is a free waveform viewer
* VS Code [Verilog HDL Extension](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/development-environments/visual-studio-code-cheat-sheet/verilog-hdl-extension.md)
 for syntax highlighting and linting
* Check out my repo [my-systemverilog-examples](https://github.com/JeffDeCola/my-systemverilog-examples)

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## OVERVIEW & HDL LEVELS

Verilog is used to describe digital system at the,

* Behavioral Level - Good for testing
* Register Transfer Level (RTL)
* Gate Level
* Switch Level

The word Verilog is a combination of the words `verification` and `logic`
because the language was first suggested as a simulation and verification tool.

## BASIC SYNTAX

First, the basics.

### BASIC STRUCTURE (THE MODULE)

A `module` is the basic building block of verilog.
It has input/output as well as the description of what it does.
Think of it like a black box.

Here is the structure of some verilog code I coped from my
[left-shift-register](https://github.com/JeffDeCola/my-systemverilog-examples/blob/master/sequential-logic/shifters/left-shift-register/left-shift-register.v).

```verilog
// A 2-bit left-shift-register
module left_shift_register(
    input        clk,           // clk
    input        rst,           // Reset
    input        d,             // d
    output [3:0] out            // out
);

// DATA TYPES
reg [3:0] out;

// SEQUENTIAL CODE
always @ (posedge clk) begin
    if (rst) begin
        out <= 4'b0000;      
    end else begin
        out <= {out[2:0], d};
    end
end

endmodule
```

### NUMBERS

`[size]'[base_format][number]`

```verilog
    4’b1001     // 1001
    12’hFA3     // 1111 1001 0011
```

### DATA TYPES

I just want to focus on these two data types,  There are others, but
I find these the most useful.

```verilog
    // DRIVER
        reg                 // Holds a state
        wire                // Connecting things (Represents a physical wire)
```

Note the name `reg` does not necessarily mean that the value is
a register. (It could be, it does not have to be).

Variables and constants are also data types,

```verilog
    // VARIABLES
        wire    [3:0] out   // out is a vector variable
        input   clk         // clk is a scalar variable
    // CONSTANT
        parameter [3:0] state = 2'b0001;
```

### SCALAR, VECTOR & ARRAYS

This is simple,

```verilog
    // SCALAR
        input   clk              // clk is a scalar variable
    // VECTOR
        reg  [3:0] y2            // y2 is a 4-bit scalar variable
    // ARRAY - Used for memory
        reg  [7:0] y3 [0:1][0:3] // y3 is a 2D array (2 rows, 4 col) each 8-bit wide
```

### OPERATORS

Pretty much in order of precedence,

```verilog
        ~                   // Bitwise NOT
    // ARITHMETIC
        *                   // Multiply
        /                   // Division
        %                   // Modulus
        +                   // Add
        -                   // Subtract
        +                   // Unary plus
        -                   // Unary minus
    // SHIFT
        >>                  // Right Shift
        <<                  // Left Shift
        <<<                 // Right Arithmetic Shift
        >>>                 // Left Arithmetic Shift
    // RELATIONAL (COMPARISON)
        >                   // Greater than
        <                   // Less than
        >=                  // Greater than or equal
        <=                  // Less than or equal
    // EQUALITY
        ==                  // Equal
        !=                  // Not Equal
    // LOGICAL (BOOLEAN)
        !                   // Logical NOT
        &&                  // Logical AND
        ||                  // Logical OR
    // REDUCTION
        &                   // Bitwise AND
        ~&                  // Bitwise NAND
        ^                   // Bitwise XOR
        ~^                  // Bitwise XNOR
        |                   // Bitwise OR
        ~|                  // Bitwise NOR
    // CONCATENATION
        {}                  // Concatenation
    // CONDITIONAL
        ?:                  // Conditional (Operates on 3 inputs)
```

## MODELING COMBINATIONAL & SEQUENTIAL LOGIC (USING 3 BASIC BUILDING BLOCKS)

* **COMBINATIONAL LOGIC**
  * Blocks that do not have memory
  * Modeled using
    [Continuous Assignment Statement (assign)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#assign-statement)
  * Modeled using an
    [Always Block](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-block-where-the-magic-happens)
    with
    [Non-Blocking Procedural Assignment Statements (<=)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-block-where-the-magic-happens)
  * Modeled using an
    [Always Block](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-block-where-the-magic-happens)
    with
    [Blocking Procedural Assignment Statements (=)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-block-where-the-magic-happens)
* **SEQUENTIAL LOGIC**
  * Blocks that have memory and triggered by a `clk` event (sensitivity list)
  * Modeled using an
    [Always Block](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-block-where-the-magic-happens)
    with
    [Non-Blocking Procedural Assignment Statements (<=)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-block-where-the-magic-happens)
  * Modeled using an
    [Always Block](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-block-where-the-magic-happens)
    with
    [Blocking Procedural Assignment Statements (=)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#always-block-where-the-magic-happens)

### ASSIGN STATEMENT

The **Continuous Assignment Statement (assign)** is used for combinational logic.

An example of combinational logic (AND gate),

```verilog
    // Model AND gate
    assign xy = x & y;
```

### ALWAYS BLOCK (WHERE THE MAGIC HAPPENS)

The **Always Block** is used for both combinational and sequential logic.

* Executes always, unlike
  [Initial Blocks](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/hardware/development/languages/systemverilog-cheat-sheet#initial-block-testbench)
  which execute only once
* Should have a sensitive list or a delay associated with it

Syntax,

```verilog
    always @ (sensitivity list) begin
        procedural statement;
    end
```

The procedural statements are either,

* **Non-Blocking Procedural Assignment Statements** (**<=**)
  * Values are assigned concurrently
* **Blocking Procedural Assignment Statements** (**=**) ONLY USE IN TESTBENCHES
  * Values are assigned sequentially

An example of combinational logic (AND gate),

```verilog
    // ALWAYS BLOCK with NON-BLOCKING PROCEDURAL ASSIGNMENT STATEMENT
    always @ ( * ) begin
        xy <= x & y;
    end

    // ALWAYS BLOCK with BLOCKING PROCEDURAL ASSIGNMENT STATEMENT - ONLY USE IN TESTBENCHES
    always @ ( * ) begin
        xy = x & y;
    end
```

An example of sequential logic (d-flip-flop),

```verilog
    // ALWAYS BLOCK with NON-BLOCKING PROCEDURAL ASSIGNMENT STATEMENT
    always @ (posedge clk) begin
        q <= d;
        q_bar <= !d;
    end

    // ALWAYS BLOCK with BLOCKING PROCEDURAL ASSIGNMENT STATEMENT - ONLY USE IN TESTBENCHES
    always @ (posedge clk) begin
        q = d;
        q_bar = !d;
    end
```

### INITIAL BLOCK (TESTBENCH)

Unlike always block, Initial Blocks executed only once when simulation starts.

* Used for testbenches
* Can not synthesize
* No sensitivity list

Syntax,

```verilog
    initial begin
        procedural statement;
    end
```

An example of initializing an d-flip-flop for testing,

```verilog
    // INITIAL BLOCK INITIALIZING D-FLIP FLOP
    initial begin
        CLK = 0;
        RST = 0;
        D = 0;

        #15; D = 1;
        #10; D = 0;
        #20; D = 1;
        etc...
    end
```

## MORE SYNTAX

Let's learn a little more about what you can do.

### GATE PRIMITIVES

Not really needed, but you can model gate primitives
such as `and`, `nand`, `or`, `nor`, `xor`.

For example,

```verilog
    and(xy, x, y);
```

### CONCATENATION

Example,

```verilog
    wire [2:0] y;
    assign y = {a, b, 1'b1};
```

### CONTROL STATEMENTS

```verilog
    // IF-ELSE
        if (expression) begin
            statements;
        end else begin
            statements;
        end
    // CASE
        case (expression) begin
            xx: statements
            xx: statements
            xx: statements
            xx: statements
        end
    // WHILE
        repeat (expression) begin
            statements;
        end
    // FOR LOOP
        for (init; condition; increment) begin
            statements;
        end
    // REPEAT
        repeat (times) begin
            statements;
        end
```

### TASKS

Tasks are used in testbenches,

```verilog
    $display("Start Simulation");
    $finish;
```

### FUNCTIONS

Functions are used to repeat code. Unlike tasks, functions can return a value.

```verilog
    module simple_function();

        function  myfunction;
            input a, b, c, d;
            begin
                myfunction = ((a+b) + (c-d));
            end
        endfunction

    endmodule
```
