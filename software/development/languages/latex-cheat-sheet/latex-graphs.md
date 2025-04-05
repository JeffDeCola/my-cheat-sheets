# LaTeX GRAPHS CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Use packages to create images._

Table of Contents

* [POPULAR GRAPHIC LaTeX PACKAGES](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/latex-cheat-sheet/latex-graphs.md#popular-graphic-latex-packages)
  * [tikzs](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/latex-cheat-sheet/latex-graphs.md#tikz)

Documentation and Reference

* [LaTex](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet)
* [LaTeX math mode](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md)

## POPULAR GRAPHIC LaTeX PACKAGES

* `tikz` - _Create graph­ics_
* `pgfplots` - _2D/3D graphing, drawing data or function plots_
* `tikz-euclide` - Eu­clidean ge­om­e­try with tikz
* `tikz-3dplot` - Can define a 3d coordinate.

### tikz

`TikZ` is a LaTeX package that allows you to create high quality diagrams.
Tikz is probably the most complex and powerful tool to create
graphic elements in LaTeX.

As an example of `a-point-on-a-circle` in
[my-latex-renders](https://github.com/JeffDeCola/my-latex-renders/tree/master/mathematics/pure/spaces/geometry/a-point-on-a-circle),

```txt
\documentclass[border=5mm]{standalone}
\usepackage{tikz}
\usetikzlibrary{calc}

\begin{document}

\begin{tikzpicture}[scale=1]

    % Draw the circle
    \draw (2,2) circle (3cm);

    % Add a point on the circle
    % For example, a point at 45 degrees (pi/4 radians) from the center
    \fill ($(2,2) + (45:3cm)$) circle (2pt);

\end{tikzpicture}

\end{document}
```

<p align="center">
    <img src="a-point-on-a-circle.svg"
    align="middle"
</p>
