# LaTeX GRAPHS CHEAT SHEET

`LaTeX graphs` _uses packages to create images._

Refer to my repo
[my-latex--graphs](https://github.com/JeffDeCola/my-latex-pgfplots-graphs)
for lots of graph examples.

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## POPULAR GRAPHIC LaTeX PACKAGES

* `tikz` - _Create graph­ics_
* `pgfplots` - _2D/3D graphing, drawing data or function plots_
* `tikz-euclide` - Eu­clidean ge­om­e­try with tikz
* `tikz-3dplot` - Can define a 3d coordinate.

### tikz

`TikZ` is a LaTeX package that allows you to create high quality diagrams.
Tikz is probably the most complex and powerful tool to create
graphic elements in LaTeX.

When you use `readme2tex make sure you add it in,

```bash
python -m readme2tex --usepackage "tikz" ....etc...
```

As an example of a circle,

```txt
    \begin{tikzpicture}
        \draw (2,2) circle (3cm);
    \end{tikzpicture}
```

Sadly, I could not get this to work with `readme2tex`.
