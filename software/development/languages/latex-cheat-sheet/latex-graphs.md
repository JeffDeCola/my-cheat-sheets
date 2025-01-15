# LaTeX GRAPHS CHEAT SHEET

_Use packages to create images._

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
