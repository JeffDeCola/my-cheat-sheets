# LaTeX CHEAT SHEET

_LaTeX is an advanced markup language used for typesetting._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#overview)
* [LaTeX .tex FILE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#latex-tex-file)
* [CREATE A .svg IMAGE FILE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#create-a-svg-image-file)
* [INSTALL LaTeX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#install-latex)
  * [LINUX](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#linux)
  * [macOS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#macos)
  * [CHECK INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#check-install)
* [INSTALL PACKAGES USING TEXLIVE MANAGER (tlmgr)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#install-packages-using-texlive-manager-tlmgr)
* [CONVERTING TO OTHER FILE FORMATS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#converting-to-other-file-formats)
* [POPULAR LaTeX PACKAGES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#popular-latex-packages)

Documentation and Reference

* [my-latex-renders](https://github.com/JeffDeCola/my-latex-renders)

## OVERVIEW

LateX is is a markup language used for describing a documents.

* [LaTeX math mode](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md)
can display math equations in LaTeX
* [LaTeX graphs](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-graphs.md)
can graph 2D and 3D images via LaTeX packages

It stated with `TeX` which is an advanced typesetting system
developed by Donald Knuth in 1978. It is a markup language
for describing a document.

TeX is designed to describe the content, not the look of the document.

`LaTeX` is a set of macros built on top of TeX. Built back in the 80s.

## LaTeX .tex FILE

The `.tex` syntax can look something like,

```latex
    \documentclass{article}
    \title{Cartesian closed categories and the price of eggs}
    \author{Jane Doe}
    \date{September 1994}
    \usepackage{amsmath}

    \begin{document}

    \begin{equation}
        E=mc^2
    \end{equation}

    \end{document}
```

## CREATE A .svg IMAGE FILE

The flow to create an image file `.svg` from a `.tex` file looks like,

![IMAGE - latex-software-flow-to-create-svg-image-file - IMAGE](../../../../docs/pics/latex-software-flow-to-create-svg-image-file.jpg)

## INSTALL LaTeX

### LINUX

This is to install a full version with all the packages,

```bash
sudo apt-get update
sudo apt-get upgrade -y
sudo apt-get install texlive-full
```

I'm not sure how to install a lighter version.

### macOS

Download and install the full version called `MacTex` from
[here](http://tug.org/mactex/mactex-download.html).

### CHECK INSTALL

Make sure you have the following,

```bash
latex -version
tlmgr -version
dvisvgm -V1
ghostscript -v
```

## INSTALL PACKAGES USING TEXLIVE MANAGER (tlmgr)

After you install, you can update and install other packages located
[here](https://ctan.org/)
using `tlmgr`. For example,

```bash
sudo tlmgr update --self
sudo tlmgr update --self --all
sudo tlmgr install standalone
sudo tlmgr install circuitikz
sudo tlmgr install amsmath
```

## CONVERTING TO OTHER FILE FORMATS

You can created a LaTeX File (.tex) and convert to many formats,

* `pdflatex` .tex -> .pdf
* `latex` .tex -> .dvi
* `dvisvgm` .dvi -> .svg
* `ghostscript` .pdf -> many types

## POPULAR LaTeX PACKAGES

There are hundreds of packages you can use with LateX such as,

* `xcolor` - _Add color_
* `amsmath` & `amsfonts` - _For rendering matrices_
* `tikz` - _Refer to my
  [LaTeX graphs](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md)
  cheat sheet_
