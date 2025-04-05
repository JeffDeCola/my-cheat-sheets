# REDIRECT OUTPUT CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Controls how you can redirect output to a file._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/operating-systems/linux/redirect-output-cheat-sheet#overview)

## OVERVIEW

I copied this table from stack exchange.

The first StdOut/StdErr pair is visible in the terminal, and the
second StdOut/StdErr pair is visible in the file.

|    Syntax     |  StdOut  |  StdErr  |  StdOut  |  StdErr  | existing file |
|---------------|----------|----------|----------|----------|---------------|
|     **>**     |    no    |   yes    |   yes    |    no    |  overwrite    |
|     **>>**    |    no    |   yes    |   yes    |    no    |   append      |
|               |          |          |          |          |               |
|    **2>**     |   yes    |    no    |    no    |   yes    |  overwrite    |
|    **2>>**    |   yes    |    no    |    no    |   yes    |   append      |
|               |          |          |          |          |               |
|    **&>**     |    no    |    no    |   yes    |   yes    |  overwrite    |
|    **&>>**    |    no    |    no    |   yes    |   yes    |   append      |
|               |          |          |          |          |               |
|  **\| tee**    |   yes    |   yes    |   yes    |    no    |  overwrite    |
|  **\| tee -a** |   yes    |   yes    |   yes    |    no    |   append      |
|               |          |          |          |          |               |
|  **n.e. (*)** |   yes    |   yes    |    no    |   yes    |  overwrite    |
|  **n.e. (*)** |   yes    |   yes    |    no    |   yes    |   append      |
|               |          |          |          |          |               |
| **\|& tee**    |   yes    |   yes    |   yes    |   yes    |  overwrite    |
| **\|& tee -a** |   yes    |   yes    |   yes    |   yes    |   append      |

For example,

File will contain jeff, but not on terminal,

```bash
echo "jeff" > test.txt
```

On terminal but the file will be empty,

```bash
echo "jeff" 2> test.txt
```

Both on terminal and in the file,

```bash
echo "jeff" | tee test.txt
```
