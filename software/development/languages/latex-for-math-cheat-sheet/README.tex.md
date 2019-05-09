# LaTeX for MATH CHEAT SHEET

`LaTeX for math` _is the math part of LaTex.  But LaTeX is much more,
an advanced typesetting system._

Documentation and reference,

* [LaTeX math document](https://www.overleaf.com/learn/latex/Mathematical_expressions)
* [LaTeX math reference](https://en.wikibooks.org/wiki/LaTeX/Mathematics)
* [mathurl.com to help build your formulas](http://mathurl.com/)

What LaTeX math functions can I use,

* [KaTeX Supported Functions](https://katex.org/docs/supported.html)
* [KaTeX Support Table](https://katex.org/docs/support_table.html)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## TeX

`TeX` is an advanced typesetting system which was
developed by Donald Knuth in 1978. It is a markup language
for describing a document. TeX is mainly popular
because of its ability to handle complex technical
text and in displaying mathematical formula.

## LaTeX

`LaTeX` is a set of macros built on top of TeX. Built back in the 80s.
LaTeX is designed for the content, not the look of the document.
A LaTeX editor is an app which enables you to write,
edit and publish your paper in LaTeX.

```txt
\documentclass{article}
\title{Cartesian closed categories and the price of eggs}
\author{Jane Doe}
\date{September 1994}
```

But I don't do any of this, I'm only interested in
the mathematical formulas. Typesetting mathematics
is one of LaTeX's greatest strengths.

## HOW I CREATED THIS README.md

I use a program called [readme2tex](https://github.com/leegao/readme2tex).
I run everything on my local machine.

It works by taking your LaTeX math formula, rendering an image and
referencing that image using HTML.

First, I created and edit in README.tex.md.
I have my LaTeX for math markdown in there delimited by dollar signs.

Second, I run `readme2tex`. This will,

* Create a new README.md file.
* Creates *.svg images for each of the LateX for math formulas.
* Add those image links automatically in your README.md file.

It created the images using the latex program and geometry package you
must install on your machine.  See the end of this cheat sheet for how I installed.

I used to use a github app called
[TeXify](https://github.com/apps/texify),
but it is not that reliable and I don't like
other programs having write access to my repo.  I feel
like i loose control.

## LaTeX for MATH (DISPLAY & INLINE MODE)

As mentioned, I only use LaTeX math capabilities. In order to use LaTeX
with markdown you need a delimiter, I like to use the dollar sign.
Delimiters are not shown for simplicity.

Display Mode,

```txt
    E=mc^2
```

You will get,

$$
E=mc^2
$$

Inline Mode using,

```txt
    Einstein's equation
    E=mc^2
    represent energy is equal to matter multiplied by the speed of light squared.
```

You will get,

Einstein's equation
$E=mc^2$
represent energy is equal to matter multiplied by the speed of light squared.

## COMMON LaTeX MATH EQUATIONS

Note, dollar sign delimiters not shown.  Also, you can't use begin{equation}
because that is not part of LateX for math.  You are now getting into LaTeX
typesetting capabilities.

You can use,

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

Pythagorean theorem (my solution to number, but it stinks),

```txt
    x^n + y^n = z^n
    \qquad \qquad (1)
```

$$
x^n + y^n = z^n
\qquad \qquad (1)
$$

Square root,

```txt
    \sqrt[3]{x}
```

$$
\sqrt[3]{x}
$$

Relations
Alignment (on equal sign),

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

## LIMITATIONS - I WOULD IKE TO DO

Color,

```txt
    \color{red} E=mc^2
```

Line numbers,

```txt
    \tag{1} E=mc^2
```

## LaTeX APPS

Some cool LaTeX apps I use.

### VISUAL STUDIO CODE APP - MARKDOWN + MATH

View LaTeX math formulas in Visual Studio Code preview window.

[MARKDOWN + MATH](https://marketplace.visualstudio.com/items?itemName=goessner.mdmath)

### BROWSER SUPPORT - KaTeX

Displays mathematical notation in web browsers.

[KaTeX](https://katex.org/docs/supported.html)

### BROWSER SANDBOX - mathurl.com

Sandbox to create your LaTeX math formulas online.

[mathurl.com](http://mathurl.com/)

### GITHUB APP - TeXify

[TeXify](https://github.com/apps/texify)
is a github app (on github) that Renders LateX to images in markdown.
Its built on
[readme2tex](https://github.com/leegao/readme2tex).

You write `README.text.md` and push.
At github, it will convert your LaTeX to `/tex/*.svg` images and
created a `README.md` file with links to those images.

This app is not that stable.

### APP - readme2tex

[readme2tex](https://github.com/leegao/readme2tex).

This cheat sheet is using this app.

You need Python 2.7 or above and pip installed.

Install full latex (I really only wanted latex and
latex geometry package) and dvisvgm,

```bash
sudo apt-get update
sudo apt-get upgrade -y
sudo apt-get install texlive-full
sudo apt install texlive-extra-utils
```

Install cairocffi and other dependencies you may need,

```bash
sudo apt-get install libffi6 libffi-dev
pip install --user cairocffi
pip install 'setuptools<36'
```

Now install readme2tex,

```bash
git clone https://github.com/leegao/readme2tex
cd readme2tex
sudo python setup.py develop
```

Again, I do not want to install the full version of latex
(3 gigs), but can't figure out how to install latex
and just the geometry package.

Usage,

```bash
python -m readme2tex --readme README.tex.md --output README.md --nocdn
```
