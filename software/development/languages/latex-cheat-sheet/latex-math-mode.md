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
<img src="svgs/ccb175704c18ad5a81177f1274fcd39f.svg?invert_in_darkmode" align="middle" width="63.09925874999999pt" height="26.76175259999998pt" />
represent energy is equal to matter multiplied by the speed of light squared.

An example of LaTeX math display mode,

```txt
    $$
    $E=mc^2$
    $$
```

<p align="center"><img src="svgs/3abb8c75967ebfdd6439c56912f3d75a.svg?invert_in_darkmode" align="middle" width="63.09925874999999pt" height="14.202794099999998pt" /></p>

An example of LaTeX math display mode (with tag),

```txt
    \begin{equation}
        E=mc^2
    \end{equation}
```

<p align="center"><img src="svgs/f334ebc0d2eff9a7f146c11edd071f96.svg?invert_in_darkmode" align="middle" width="382.09755705pt" height="18.312383099999998pt" /></p>

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

<p align="center"><img src="svgs/61e53b5569a8e3e0af95fdfbec895cdf.svg?invert_in_darkmode" align="middle" width="92.47868355pt" height="39.2694126pt" /></p>

An example of LaTeX math display mode (using gathered),

```txt
    $$
    \begin{gathered}
        a=b+c \\
        d+e=f
    \end{gathered}
    $$
```

<p align="center"><img src="svgs/b81b3a8ee17288b65f2e80ae31ac2c24.svg?invert_in_darkmode" align="middle" width="68.03633925pt" height="39.2694126pt" /></p>

## COMMON LaTeX MATH EQUATIONS

For brevity, the dollar sign delimiters are not shown.

As a reminder you can use,

* [KaTeX Supported Functions](https://katex.org/docs/supported.html)
* [KaTeX Support Symbols](https://katex.org/docs/support_table.html)

Einsteins famous equation,

```txt
    E=mc^2
```

<p align="center"><img src="svgs/3abb8c75967ebfdd6439c56912f3d75a.svg?invert_in_darkmode" align="middle" width="63.09925874999999pt" height="14.202794099999998pt" /></p>

Pythagorean theorem,

```txt
    x^n + y^n = z^n
```

<p align="center"><img src="svgs/238cd7abcefb8a6a256d0bec59744770.svg?invert_in_darkmode" align="middle" width="94.44253334999999pt" height="14.937954899999998pt" /></p>

Pythagorean theorem (with box),

```txt
    \boxed{x^n + y^n = z^n}
```

<p align="center"><img src="svgs/2dc0d735a1c4f4b4bb63782b8b4e8875.svg?invert_in_darkmode" align="middle" width="106.44247019999999pt" height="26.11600695pt" /></p>

Pythagorean theorem (numbered using begin{equation}),

```txt
    \begin{equation}
        x^n + y^n = z^n
    \end{equation}
```

<p align="center"><img src="svgs/9c2c2ace77591d2c08d3ea86c859717c.svg?invert_in_darkmode" align="middle" width="397.76920259999997pt" height="16.438356pt" /></p>

Pythagorean theorem (numbered using tag),

```txt
    \begin{equation}
        x^n + y^n = z^n\tag{6}
    \end{equation}
```

<p align="center"><img src="svgs/508f5e725b18bed700d18ff1cea4238c.svg?invert_in_darkmode" align="middle" width="397.76920259999997pt" height="16.438356pt" /></p>

Pythagorean theorem (another way to do numbers),

```txt
    \qquad \qquad
    x^n + y^n = z^n
    \qquad \qquad (4)
```

<p align="center"><img src="svgs/5b71346bca16eedf780a40e41ddef9e8.svg?invert_in_darkmode" align="middle" width="182.02261604999998pt" height="16.438356pt" /></p>

Square root,

```txt
    \sqrt[3]{x}
```

<p align="center"><img src="svgs/8375620ec7847ff048e565abd48c90cd.svg?invert_in_darkmode" align="middle" width="23.093660699999997pt" height="16.438356pt" /></p>

Multiple equations (gathered),

```txt
    \begin{gathered}
        a=b+c \\
        d+e=f
    \end{gathered}
```

<p align="center"><img src="svgs/b81b3a8ee17288b65f2e80ae31ac2c24.svg?invert_in_darkmode" align="middle" width="68.03633925pt" height="39.2694126pt" /></p>

Multiple equations (aligned on ampersand &),

```txt
    \begin{aligned}
        a&=b+c \\
        d+e&=f
    \end{aligned}
```

<p align="center"><img src="svgs/61e53b5569a8e3e0af95fdfbec895cdf.svg?invert_in_darkmode" align="middle" width="92.47868355pt" height="39.2694126pt" /></p>

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

<p align="center"><img src="svgs/7e7e2b03aa45736c18ad92fa9b6f9b47.svg?invert_in_darkmode" align="middle" width="197.88809369999998pt" height="205.18438049999997pt" /></p>

An integral,

```txt
    \int_{a}^{b} x^2 dx
```

<p align="center"><img src="svgs/7434eb168b5dfced915545b6a422e7b8.svg?invert_in_darkmode" align="middle" width="60.50117205pt" height="41.27894265pt" /></p>

Limits,

```txt
    \lim_{x\to\infty} f(x)
```

<p align="center"><img src="svgs/b7ecb947fb5547679f1c7ab9a546d2ce.svg?invert_in_darkmode" align="middle" width="68.4019611pt" height="22.1917806pt" /></p>

Some trigonometry,

```txt
    \sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)
```

<p align="center"><img src="svgs/85b567a60b6ab8fbc319c720f66f8ae2.svg?invert_in_darkmode" align="middle" width="283.3047162pt" height="16.438356pt" /></p>

Fractions (binomial coefficients),

```txt
    \binom{n}{k} = \frac{n!}{k!(n-k)!}
```

<p align="center"><img src="svgs/9cc892f15c1314868714ad2b49649eb5.svg?invert_in_darkmode" align="middle" width="127.98480255pt" height="39.452455349999994pt" /></p>

Brackets,

```txt
    \left( \frac{x}{y} \right)
```

<p align="center"><img src="svgs/2beb7431726139ffb37413b4031f73c1.svg?invert_in_darkmode" align="middle" width="37.5411861pt" height="39.452455349999994pt" /></p>

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

<p align="center"><img src="svgs/7192bf1b080bb1180da0295fcda68a71.svg?invert_in_darkmode" align="middle" width="115.5253011pt" height="59.1786591pt" /></p>

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

<p align="center"><img src="svgs/34937d1b68e46828a9241c64dd9c8c8a.svg?invert_in_darkmode" align="middle" width="267.12360674999996pt" height="59.178683850000006pt" /></p>

Some cool arrows,

```txt
    A \xleftarrow{\text{this way}} B
    \xrightarrow[\text{or that way}]{ } C
```

<p align="center"><img src="svgs/056bb44ae5a648c007f840c99bc3cbdf.svg?invert_in_darkmode" align="middle" width="192.42537105pt" height="31.053602249999997pt" /></p>

Cases,

```txt
    u(x) =
    \begin{cases}
        \exp{x} & \text{if } x \geq 0 \\
        1       & \text{if } x < 0
    \end{cases}
```

<p align="center"><img src="svgs/7cdb2cc12c0acb192d47b17946ace840.svg?invert_in_darkmode" align="middle" width="175.0379664pt" height="49.315569599999996pt" /></p>

Make some big text,

```txt
    $$
    \huge\text{Hello Jeff}
    $$
```

<p align="center"><img src="svgs/e352edc1951d2b05c08e5e2f08d50273.svg?invert_in_darkmode" align="middle" width="130.30480485pt" height="23.675812049999998pt" /></p>
