# IVERILOG CHEAT SHEET

_How to install and check `iverilog` (Icaris Verilog), a tool for
simulation and synthesis._

Documentation and reference,

* Icarus Verilog
  [Home page](http://iverilog.icarus.com/)
* [Repo](  https://github.com/steveicarus/iverilog)
* [Installation guide](https://iverilog.fandom.com/wiki/Installation_Guide)

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## INSTALL IVERILOG

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

