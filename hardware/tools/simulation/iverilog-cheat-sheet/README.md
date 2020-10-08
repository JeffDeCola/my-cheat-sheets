# IVERILOG CHEAT SHEET

_How to install and check `iverilog`_

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## INSTALL IVERILOG

I installed `iverilog` for simulation and synthesis.

### Install from package

Gets placed in `/usr/bin`.

This is easier,

```bash
sudo apt-get update
sudo apt-get install verilog
```

### Install from Source

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
```

Build configuration files,

```bash
cd steveicarus/iverilog
sh autoconf.sh
```

Now lets compile your source,

```bash
./configure
make
sudo su
make install
```

## CHECK

Check,

```bash
iverilog -h
```

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

Execute,

```bash
./hello
```