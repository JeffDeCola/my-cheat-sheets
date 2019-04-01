# REDIRECT OUTPUT CHEAT SHEET

`redirect output` _controls how you can redirect output to a file._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

I copied this table from stack exchange.

The first StdOut/StdErr pair are visible in terminal, and the second pair
is visible in file.

|  Syntax   |  StdOut  |  StdErr  |  StdOut  |  StdErr  | existing file |
|-----------|----------|----------|----------|----------|---------------|
|     >     |    no    |   yes    |   yes    |    no    |  overwrite    |
|     >>    |    no    |   yes    |   yes    |    no    |   append      |
|           |          |          |          |          |               |
|    2>     |   yes    |    no    |    no    |   yes    |  overwrite    |
|    2>>    |   yes    |    no    |    no    |   yes    |   append      |
|           |          |          |          |          |               |
|    &>     |    no    |    no    |   yes    |   yes    |  overwrite    |
|    &>>    |    no    |    no    |   yes    |   yes    |   append      |
|           |          |          |          |          |               |
|  | tee    |   yes    |   yes    |   yes    |    no    |  overwrite    |
|  | tee -a |   yes    |   yes    |   yes    |    no    |   append      |
|           |          |          |          |          |               |
|  n.e. (*) |   yes    |   yes    |    no    |   yes    |  overwrite    |
|  n.e. (*) |   yes    |   yes    |    no    |   yes    |   append      |
|           |          |          |          |          |               |
| |& tee    |   yes    |   yes    |   yes    |   yes    |  overwrite    |
| |& tee -a |   yes    |   yes    |   yes    |   yes    |   append      |
