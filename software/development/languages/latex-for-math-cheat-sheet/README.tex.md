# LaTeX for MATH CHEAT SHEET

`LaTeX for math` _is the math part of LaTex.  But LaTeX is much more,
an advanced typesetting system._

Documentation and reference,

* [LaTeX math](https://en.wikibooks.org/wiki/LaTeX/Mathematics)
* [Another great website for LaTeX math](https://www.overleaf.com/learn/latex/Mathematical_expressions)
* [mathurl.com to help build your formulas](http://mathurl.com/)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## TeX

`TeX` is an advanced typesetting system which was
developed by Donald Knuth in 1978. TeX is mainly popular
because of its ability to handle complex technical
text and in displaying mathematical formula.

## LaTeX

A set of macros built on top of Tex. Built back in the 80s.
LaTeX is designed for the content, not the look of the document.
A LaTeX editor is an app which enables you to write,
edit and publish your paper in LaTeX.

```
\documentclass{article}
\title{Cartesian closed categories and the price of eggs}
\author{Jane Doe}
\date{September 1994}
\begin{document}
   \maketitle
   Hello world!
\end{document}
```

But I don't do any of this, I'm only interested in
the mathematical formulas. Typesetting mathematics
is one of LaTeX's greatest strengths.

## LaTeX for MATH (INLINE & DISPLAY MODE)

I only use LaTeX math capabilities. In order to use with markdown
you need a delimiter, I like to use the dollar sign.

Display level,

\$\$
E=mc^2
\$\$

You will get,

$$
E=mc^2
$$

Display level with numbering,

```txt
\$\$
E=mc^2
\$\$ (1)
```

You will get,

$$
E=mc^2
$$ (1)

For inline,

```txt
Einstein's equation
\$E=mc^2\$
represent energy is equal to matter multiplied by the speed of light squared.
```

You will get,

Einstein's equation
$E=mc^2$
represent energy is equal to matter multiplied by the speed of light squared.

## COMMON MATH EQUATIONS

Einsteins famous equation,

```txt
$$
E=mc^2
$$ (1)
```

$$
E=mc^2
$$ (1)

Pythagorean theorem,

```txt
$$
x^n + y^n = z^n
$$ (2)
```

$$
x^n + y^n = z^n
$$ (2)

A sample Integral,

```txt
$$
\int_{a}^{b} x^2 dx
$$ (3)
```

$$
\int_{a}^{b} x^2 dx
$$ (3)

Limits,

```txt
$$
\lim_{x\to\infty} f(x)
$$ (4)
```

$$
\lim_{x\to\infty} f(x)
$$ (4)

Some trigonometry,

```txt
$$
\sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)
$$ (5)
```

$$
\sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)
$$ (5)

Fractions (binomial coefficients),

```txt
$$
\binom{n}{k} = \frac{n!}{k!(n-k)!}
$$ (6)
```

$$
\binom{n}{k} = \frac{n!}{k!(n-k)!}
$$ (6)

Brackets,

```txt
$$
\left( \frac{x}{y} \right)
$$ (7)
```

$$
\left( \frac{x}{y} \right)
$$ (7)

Arrays in Brackets with spacing (\qquad),

```txt
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
$$ (8)
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
$$ (8)

## LaTeX APPS

Some cool LaTeX apps I use.

### VISUAL STUDIO CODE APP - MARKDOWN + MATH

View your LaTeX in your markdown file in a preview window.

[VS Code Display Latex Preview](https://marketplace.visualstudio.com/items?itemName=goessner.mdmath).

### BROWSER APP - KaTeX

Displays mathematical notation in web browsers.

[KaTeX Overview](https://katex.org/docs/supported.html).

### BROWSER WEBSITE - mathurl.com

[mathurl.com](http://mathurl.com/)
will help you create your equations.

### GITHUB APP - TeXify

[TeXify](https://github.com/apps/texify)
is a github app (on github) that Renders LateX to images in markdown.
It is built on
[readme2tex](https://github.com/leegao/readme2tex)

You write `README.text.md` and push.
At github, it will convert your LaTeX to `/tex/*.svg` images and
created a `README.md` file with links to those images.

This cheat sheet is using this tool.
