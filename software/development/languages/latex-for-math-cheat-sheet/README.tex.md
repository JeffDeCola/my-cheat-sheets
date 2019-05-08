# LaTeX for MATH CHEAT SHEET

`LaTeX for math` _is the math part of LaTex.  But LaTeX is much more,
an advanced typesetting system._

Documentation and reference,

* [LaTeX math document](https://www.overleaf.com/learn/latex/Mathematical_expressions)
* [LaTeX math reference](https://en.wikibooks.org/wiki/LaTeX/Mathematics)
* [mathurl.com to help build your formulas](http://mathurl.com/)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## TeX

`TeX` is an advanced typesetting system which was
developed by Donald Knuth in 1978. TeX is mainly popular
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

## LaTeX for MATH (DISPLAY & INLINE MODE)

I only use LaTeX math capabilities. In order to use LaTeX
with markdown you need a delimiter, I like to use the dollar sign.
Delimiters are not shown.

Display Mode,

```txt
/$$$/$
E=mc^2
/$$$/$
```

You will get,

$$
E=mc^2
$$

Inline Mode using,

```txt
Einstein's equation
/$$E=mc^2$/$
represent energy is equal to matter multiplied by the speed of light squared.
```

You will get,

Einstein's equation
$E=mc^2$
represent energy is equal to matter multiplied by the speed of light squared.

## COMMON MATH EQUATIONS

Note, dollar sign delimiters not shown,

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

A sample Integral,

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
\quad
\left\{
  \begin{array}{ccc}
  1 & 5 & 8\\
  0 & 2 & 4\\
  3 & 3 & -8
  \end{array}
\right\}
$$

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

This cheat sheet is using this tool.
