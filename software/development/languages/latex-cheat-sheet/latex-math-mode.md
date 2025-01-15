# LaTeX MATH MODE CHEAT SHEET

_The math mode of LaTex._

Table of Contents

* [DECLARING MATH MODE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md#declaring-math-mode)
* [FORMATTING MULTIPLE EQUATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md#formatting-multiple-equations)
* [COMMON LaTeX MATH EQUATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md#common-latex-math-equations)

Documentation and Reference

* [LaTeX math mode document](https://www.overleaf.com/learn/latex/Mathematical_expressions)
* [LaTeX math mode reference](https://en.wikibooks.org/wiki/LaTeX/Mathematics)
* [KaTeX Supported Functions](https://katex.org/docs/supported.html)
* [KaTeX Support Symbols](https://katex.org/docs/support_table.html)
* [mathurl.com to help build your formulas](http://mathurl.com/)

## DECLARING MATH MODE

First, there are three ways to declare math mode in LaTeX,

1. In-line mode using **$**
2. Block mode using **$$**
3. Block mode using **begin{equation}** ... **end{equation}**

### In-line Mode

```text
Einstein's equation
$E=mc^2$
represent energy is equal to matter multiplied by the speed of light squared.
```

Einstein's equation
$E=mc^2$
represent energy is equal to matter multiplied by the speed of light squared.

### Block Mode

```text
$$
$E=mc^2$
$$
```

$$
\boxed ab
$$

$$
E=mc^2
$$

### Block Mode using begin{equation}

```text
\begin{equation}
  E=mc^2
\end{equation}
```

$$
  \begin{equation}
      E=mc^2
  \end{equation}
$$

## FORMATTING MULTIPLE EQUATIONS

There are two main ways to format multiple equations in LaTeX math mode,

* `begin{aligned}...end{aligned}` - Will align on an ampersand
* `begin{gathered}...end{gathered}` - Every equation centered

Block mode (using aligned),

```text
$$
\begin{aligned}
  a&=b+c \\
  d+e&=f
\end{aligned}
$$
```

$$
\begin{aligned}
  a&=b+c \\
  d+e&=f
\end{aligned}
$$

Block mode (using gathered),

```text
$$
\begin{gathered}
  a=b+c \\
  d+e=f
\end{gathered}
$$
```

$$
\begin{gathered}
  a=b+c \\
  d+e=f
\end{gathered}
$$

## COMMON LaTeX MATH EQUATIONS

_For brevity, the dollar sign delimiters are not shown._

**Einsteins famous equation**

```text
E=mc^2
```

$$
E=mc^2
$$

**Pythagorean theorem**

```text
x^n + y^n = z^n
```

$$
x^n + y^n = z^n
$$

**Pythagorean theorem (with box)**

```text
\boxed{x^n + y^n = z^n}
```

$$
\boxed{x^n + y^n = z^n}
$$

**Pythagorean theorem (numbered using begin{equation})**

```text
    \begin{equation}
        x^n + y^n = z^n
    \end{equation}
```

$$
\begin{equation}
    x^n + y^n = z^n
\end{equation}
$$

**Pythagorean theorem (numbered using tag)**

```text
    \begin{equation}
        x^n + y^n = z^n \tag{6}
    \end{equation}
```

$$
\begin{equation}
    x^n + y^n = z^n \tag{6}
\end{equation}
$$

**Pythagorean theorem (another way to do numbers)**

```text
    \qquad \qquad
    x^n + y^n = z^n
    \qquad \qquad (4)
```

$$
\qquad \qquad
x^n + y^n = z^n
\qquad \qquad (4)
$$

**Square root**

```text
    \sqrt[3]{x}
```

$$
\sqrt[3]{x}
$$

**Multiple equations (gathered)**

```text
    \begin{gathered}
        a=b+c \\
        d+e=f
    \end{gathered}
```

$$
\begin{gathered}
    a=b+c \\
    d+e=f
\end{gathered}
$$

**Multiple equations (aligned on ampersand &)**

```text
    \begin{aligned}
        a&=b+c \\
        d+e&=f
    \end{aligned}
```

$$
\begin{aligned}
    a&=b+c \\
    d+e&=f
\end{aligned}
$$

**Alignment and spacing (on equal sign)**

```text
    \begin{aligned}
        f(x) =& x^2\! +3x\! +2 \\
        f(x) =& x^2+3x+2 \\
        f(x) =& x^2\, +3x\, +2 \\
        f(x) =& x^2\: +3x\: +2 \\
        f(x) =& x^2\; +3x\; +2 \\
        f(x) =& x^2\ +3x\ +2 \\
        f(x) =& x^2\quad +3x\quad +2 \\
        f(x) =& x^2\qquad +3x\qquad +2
    \end{aligned}
```

$$
\begin{aligned}
    f(x) =& x^2\! +3x\! +2 \\
    f(x) =& x^2+3x+2 \\
    f(x) =& x^2\, +3x\, +2 \\
    f(x) =& x^2\: +3x\: +2 \\
    f(x) =& x^2\; +3x\; +2 \\
    f(x) =& x^2\ +3x\ +2 \\
    f(x) =& x^2\quad +3x\quad +2 \\
    f(x) =& x^2\qquad +3x\qquad +2
\end{aligned}
$$

**An integral**

```text
    \int_{a}^{b} x^2 dx
```

$$
\int_{a}^{b} x^2 dx
$$

**Limits**

```text
    \lim_{x\to\infty} f(x)
```

$$
\lim_{x\to\infty} f(x)
$$

**Some trigonometry**

```text
    \sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)
```

$$
\sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)
$$

**Fractions (binomial coefficients)**

```text
    \binom{n}{k} = \frac{n!}{k!(n-k)!}
```

$$
\binom{n}{k} = \frac{n!}{k!(n-k)!}
$$

**Brackets**

```text
    \left( \frac{x}{y} \right)
```

$$
\left( \frac{x}{y} \right)
$$

**Bracket array**

```text
    \left(
        \begin{array}{ccc}
            1 & 2 & 3\\
            4 & 4 & 9\\
            1 & -8 & 2
        \end{array}
    \right)
```

$$
\left(
    \begin{array}{ccc}
        1 & 2 & 3\\
        4 & 4 & 9\\
        1 & -8 & 2
    \end{array}
\right)
$$

**Arrays in Brackets with spacing (\qquad)**

```text
    \left(
        \begin{array}{ccc}
            1 & 2 & 3\\
            4 & 5 & 9\\
            1 & -8 & 2
        \end{array}
    \right)
    \qquad
    \left\{
        \begin{array}{ccc}
            1 & 5 & 8\\
            0 & 2 & 4\\
            3 & 3 & -8
        \end{array}
    \right\}
```

$$
\left(
    \begin{array}{ccc}
        1 & 2 & 3\\
        4 & 5 & 9\\
        1 & -8 & 2
    \end{array}
\right)
\qquad
\left\{
    \begin{array}{ccc}
        1 & 5 & 8\\
        0 & 2 & 4\\
        3 & 3 & -8
    \end{array}
\right\}
$$

**Some cool arrows**

```text
    A \xleftarrow{\text{this way}} B
    \xrightarrow[\text{or that way}]{ } C
```

$$
A \xleftarrow{\text{this way}} B
\xrightarrow[\text{or that way}]{ } C
$$

**Cases**

```text
    u(x) =
    \begin{cases}
        \exp{x} & \text{if } x \geq 0 \\
        1       & \text{if } x < 0
    \end{cases}
```

$$
u(x) =
\begin{cases}
    \exp{x} & \text{if } x \geq 0 \\
    1       & \text{if } x < 0
\end{cases}
$$

**Make some big text**

```text
$$
    \huge\text{Hello Jeff}
$$
```

$$
\huge\text{Hello Jeff}
$$
