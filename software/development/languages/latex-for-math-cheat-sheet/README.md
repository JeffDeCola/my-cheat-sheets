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

## LaTeX for MATH (DISPLAY & INLINE MODE)

I only use LaTeX math capabilities. In order to use LaTeX
with markdown you need a delimiter, I like to use the dollar sign.
Delimiters are not shown.

Display Mode,

```txt
E=mc^2
```

You will get,

<p align="center"><img alt="$$&#10;E=mc^2&#10;$$" src="svgs/3abb8c75967ebfdd6439c56912f3d75a.svg" align="middle" width="62.901135pt" height="14.175084pt"/></p>

Inline Mode using,

```txt
Einstein's equation
E=mc^2
represent energy is equal to matter multiplied by the speed of light squared.
```

You will get,

Einstein's equation
<img alt="$E=mc^2$" src="svgs/ccb175704c18ad5a81177f1274fcd39f.svg" align="middle" width="62.9013pt" height="26.70657pt"/>
represent energy is equal to matter multiplied by the speed of light squared.

## COMMON MATH EQUATIONS

Note, dollar sign delimiters not shown,

Einsteins famous equation,

```txt
E=mc^2
```

<p align="center"><img alt="$$&#10;E=mc^2&#10;$$" src="svgs/3abb8c75967ebfdd6439c56912f3d75a.svg" align="middle" width="62.901135pt" height="14.175084pt"/></p>

Pythagorean theorem,

```txt
x^n + y^n = z^n
```

<p align="center"><img alt="$$&#10;x^n + y^n = z^n&#10;$$" src="svgs/238cd7abcefb8a6a256d0bec59744770.svg" align="middle" width="94.22292pt" height="14.9075025pt"/></p>

A sample Integral,

```txt
\int_{a}^{b} x^2 dx
```

<p align="center"><img alt="$$&#10;\int_{a}^{b} x^2 dx&#10;$$" src="svgs/7434eb168b5dfced915545b6a422e7b8.svg" align="middle" width="60.399075pt" height="41.24901pt"/></p>

Limits,

```txt
\lim_{x\to\infty} f(x)
```

<p align="center"><img alt="$$&#10;\lim_{x\to\infty} f(x)&#10;$$" src="svgs/b7ecb947fb5547679f1c7ab9a546d2ce.svg" align="middle" width="68.289045pt" height="22.14564pt"/></p>

Some trigonometry,

```txt
\sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)
```

<p align="center"><img alt="$$&#10;\sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)&#10;$$" src="svgs/85b567a60b6ab8fbc319c720f66f8ae2.svg" align="middle" width="282.74565pt" height="16.376943pt"/></p>

Fractions (binomial coefficients),

```txt
\binom{n}{k} = \frac{n!}{k!(n-k)!}
```

<p align="center"><img alt="$$&#10;\binom{n}{k} = \frac{n!}{k!(n-k)!}&#10;$$" src="svgs/9cc892f15c1314868714ad2b49649eb5.svg" align="middle" width="127.93704pt" height="39.30498pt"/></p>

Brackets,

```txt
\left( \frac{x}{y} \right)
```

<p align="center"><img alt="$$&#10;\left( \frac{x}{y} \right)&#10;$$" src="svgs/2beb7431726139ffb37413b4031f73c1.svg" align="middle" width="37.49592pt" height="39.30498pt"/></p>

Bracket array,

```txt
TBD
```

<p align="center"><img alt="$$&#10;\left(&#10; \begin{array}{ccc}&#10;  1 &amp; 2 &amp; 3\\&#10;  4 &amp; 4 &amp; 9\\&#10;  1 &amp; -8 &amp; 2&#10; \end{array}&#10;\right)&#10;$$" src="svgs/239df225910aa3c7dccfe23cef3e1681.svg" align="middle" width="115.47162pt" height="59.068185pt"/></p>

Arrays in Brackets with spacing (\qquad),

```txt
TBD
```

<p align="center"><img alt="$$&#10;\left(&#10; \begin{array}{ccc}&#10;  1 &amp; 2 &amp; 3\\&#10;  4 &amp; 5 &amp; 9\\&#10;  1 &amp; -8 &amp; 2&#10; \end{array}&#10;\right)&#10;\quad&#10;\left\{&#10;  \begin{array}{ccc}&#10;  1 &amp; 5 &amp; 8\\&#10;  0 &amp; 2 &amp; 4\\&#10;  3 &amp; 3 &amp; -8&#10;  \end{array}&#10;\right\}&#10;$$" src="svgs/8e653c22e26b54bf3d59ac7699ce9318.svg" align="middle" width="250.63005pt" height="59.12346pt"/></p>

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
