# LaTeX MATH MODE CHEAT SHEET

`LaTeX math mode` _is the math mode of LaTex. Where LaTeX is
an advanced markup language used for typesetting._

Documentation and reference,

* [LaTeX math mode document](https://www.overleaf.com/learn/latex/Mathematical_expressions)
* [LaTeX math mode reference](https://en.wikibooks.org/wiki/LaTeX/Mathematics)
* [mathurl.com to help build your formulas](http://mathurl.com/)

LaTeX math mode functions & symbols,

* [KaTeX Supported Functions](https://katex.org/docs/supported.html)
* [KaTeX Support Symbols](https://katex.org/docs/support_table.html)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## DECLARING MATH MODE

First, there are three ways to `declare math mode` in LaTeX,

* In-line mode using \$...\$
* Display mode (Newline, center) using \$\$...\$\$
* Display mode (Newline, center with tag) using begin{equation}...end{equation}

An example of LaTeX math in-line mode,

```txt
    Einstein's equation
    $E=mc^2$
    represent energy is equal to matter multiplied by the speed of light squared.
```

Einstein's equation
$E=mc^2$
represent energy is equal to matter multiplied by the speed of light squared.

An example of LaTeX math display mode,

```txt
    $$
    $E=mc^2$
    $$
```

$$
E=mc^2
$$

An example of LaTeX math display mode (with tag),

```txt
    \begin{equation}
        E=mc^2
    \end{equation}
```

\begin{equation}
    E=mc^2
\end{equation}

This is sometimes hard to render being{equation} correctly,
so I kind of stay away from it.  I have a way to make my own tags below.

## FORMATING MULTIPLE EQUATIONS

There are two main ways to format multiple equations in LaTeX math mode,

* `begin{aligned}...end{aligned}` - Will align on an ampersand
* `begin{gathered}...end{gathered}` - Every equation centered

An example of LaTeX math display mode (using aligned),

```txt
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

An example of LaTeX math display mode (using gathered),

```txt
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

For brevity, the dollar sign delimiters are not shown.

As a reminder you can use,

* [KaTeX Supported Functions](https://katex.org/docs/supported.html)
* [KaTeX Support Symbols](https://katex.org/docs/support_table.html)

Einsteins famous equation,

```txt
    E=mc^2
```

$$
E=mc^2
$$

Pythagorean theorem,

```txt
    x^n + y^n = z^n
```

$$
x^n + y^n = z^n
$$

Pythagorean theorem (with box),

```txt
    \boxed{x^n + y^n = z^n}
```

$$
\boxed{x^n + y^n = z^n}
$$

Pythagorean theorem (numbered using begin{equation}),

```txt
    \begin{equation}
        x^n + y^n = z^n
    \end{equation}
```

\begin{equation}
    x^n + y^n = z^n
\end{equation}

Pythagorean theorem (numbered using tag),

```txt
    \begin{equation}
        x^n + y^n = z^n\tag{6}
    \end{equation}
```

\begin{equation}
    x^n + y^n = z^n\tag{6}
\end{equation}

Pythagorean theorem (another way to do numbers),

```txt
    \qquad \qquad
    x^n + y^n = z^n
    \qquad \qquad (4)
```

$$
\qquad \qquad
x^n + y^n = z^n
\qquad \qquad (4)
$$

Square root,

```txt
    \sqrt[3]{x}
```

$$
\sqrt[3]{x}
$$

Multiple equations (gathered),

```txt
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

Multiple equations (aligned on ampersand &),

```txt
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

Alignment and spacing (on equal sign),

```txt
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

An integral,

```txt
    \int_{a}^{b} x^2 dx
```

$$
\int_{a}^{b} x^2 dx
$$

Limits,

```txt
    \lim_{x\to\infty} f(x)
```

$$
\lim_{x\to\infty} f(x)
$$

Some trigonometry,

```txt
    \sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)
```

$$
\sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)
$$

Fractions (binomial coefficients),

```txt
    \binom{n}{k} = \frac{n!}{k!(n-k)!}
```

$$
\binom{n}{k} = \frac{n!}{k!(n-k)!}
$$

Brackets,

```txt
    \left( \frac{x}{y} \right)
```

$$
\left( \frac{x}{y} \right)
$$

Bracket array,

```txt
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

Arrays in Brackets with spacing (\qquad),

```txt
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

Some cool arrows,

```txt
    A \xleftarrow{\text{this way}} B
    \xrightarrow[\text{or that way}]{ } C
```

$$
A \xleftarrow{\text{this way}} B
\xrightarrow[\text{or that way}]{ } C
$$

Cases,

```txt
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

Make some big text,

```txt
    $$
    \huge\text{Hello Jeff}
    $$
```

$$
\huge\text{Hello Jeff}
$$
