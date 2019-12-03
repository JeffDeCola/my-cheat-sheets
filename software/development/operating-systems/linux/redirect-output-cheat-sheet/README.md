# REDIRECT OUTPUT CHEAT SHEET

_Controls how you can redirect output to a file._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

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
