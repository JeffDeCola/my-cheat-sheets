# LaTeX MATH MODE CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_The math mode of LaTex._

Table of Contents

* [DECLARING MATH MODE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md#declaring-math-mode)
* [FORMATTING MULTIPLE EQUATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md#formatting-multiple-equations)
* [MATH EXAMPLES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md#math-examples)
* [TEXT EXAMPLES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md#text-examples)

Documentation and Reference

* [LaTex](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet)
* [LaTeX graphs](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-graphs.md)
* [LaTeX math mode document](https://www.overleaf.com/learn/latex/Mathematical_expressions)
* [LaTeX math mode reference](https://en.wikibooks.org/wiki/LaTeX/Mathematics)
* [KaTeX Supported Functions](https://katex.org/docs/supported.html)
* [KaTeX Support Symbols](https://katex.org/docs/support_table.html)
* [mathurl.com](http://mathurl.com/)
  to help build your formulas

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
E=mc^2
$$
```

$$
E=mc^2
$$

### Block Mode equation numbering

This may not work on github,

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

If begin{equation} doesn't work on github, you could do this,

```text
$$
E=mc^2 \qquad (1)
$$
```

$$
E=mc^2 \qquad (1)
$$

## FORMATTING MULTIPLE EQUATIONS

There are two main ways to format multiple equations in LaTeX math mode,

* `begin{aligned}...end{aligned}` - Will align on an ampersand
* `begin{gathered}...end{gathered}` - Every equation centered

**Block mode (using aligned)**

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

**Block mode (using gathered)**

```text
$$
\begin{gathered}
    a=b+c \\
    d+e=f+s
\end{gathered}
$$
```

$$
\begin{gathered}
    a=b+c \\
    d+e=f+s
\end{gathered}
$$

## MATH EXAMPLES

_For brevity, the dollar sign delimiters are not shown._

**Einsteins famous equation**

```text
E=mc^2
```

$$
E=mc^2
$$

**Einsteins famous equation with box (Does not work on github)**

```text
\boxed{E=mc^2}
```

$$
\boxed{E=mc^2}
$$

**Einsteins famous equation with color**

```text
\color{red}{E=mc^2}
```

$$
\color{red}{E=mc^2}
$$

**Pythagorean Theorem**

```text
x^n + y^n = z^n
```

$$
x^n + y^n = z^n
$$

**Pythagorean theorem (add numbers)**

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
\left[
    \begin{array}{ccc}
        1 & 5 & 8\\
        0 & 2 & 4\\
        3 & 3 & -8
    \end{array}
\right]
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
\left[
    \begin{array}{ccc}
        1 & 5 & 8\\
        0 & 2 & 4\\
        3 & 3 & -8
    \end{array}
\right]
$$

**Some cool arrows**

```text
A
\xleftarrow[]{
    \text{this way} }
B
\xrightarrow[]{
    \text{or that way}}
C
```

$$
A
\xleftarrow[]{
    \text{this way} }
B
\xrightarrow[]{
    \text{or that way}}
C
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

**Getting fancy**

```text
\begin{align*}
    E_1 & = \frac{e_{1max}}{\sqrt{2}}\\
    & = \frac{N_1.{\phi}_m.w}{\sqrt{2}}\\
    & = \frac{N_1{\phi}_m.2\pi.f}{\sqrt{2}}\\
    \\
    E_1 & = 4.44 N_1.{\phi}_m.f \\
    \\
    E_1 &= \boxed{4.44 N_1.B_m.A.f}
\end{align*}
```

$$
\begin{align*}
    E_1 & = \frac{e_{1max}}{\sqrt{2}}\\
    & = \frac{N_1.{\phi}_m.w}{\sqrt{2}}\\
    & = \frac{N_1{\phi}_m.2\pi.f}{\sqrt{2}}\\
    \\
    E_1 & = 4.44 N_1.{\phi}_m.f \\
    \\
    E_1 &= \boxed{4.44 N_1.B_m.A.f}
\end{align*}
$$

## TEXT EXAMPLES

**Some big text**

```text
\huge\text{Hello Jeff}
```

$$
\huge\text{Hello Jeff}
$$
