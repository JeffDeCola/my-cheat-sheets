# COMBINATIONS CHEAT SHEET

```txt
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

`combinations` _tbd._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

Einstein's famous equation,

<p align="center"><img alt="$$&#10;E=mc^2&#10;$$" src="svgs/3abb8c75967ebfdd6439c56912f3d75a.svg" align="middle" width="62.901135pt" height="14.175084pt"/></p>

<p align="center"><img alt="$$&#10;\huge\text{Hello \LaTeX}&#10;$$" src="svgs/d27ecd9d6334c7a020001926c8000801.svg" align="middle" width="159.690135pt" height="30.925785pt"/></p>

<p align="center"><img alt="$$&#10;\begin{tikzpicture}&#10;\newcounter{density}&#10;\setcounter{density}{20}&#10;    \def\couleur{blue}&#10;    \path[coordinate] (0,0)  coordinate(A)&#10;                ++( 60:6cm) coordinate(B)&#10;                ++(-60:6cm) coordinate(C);&#10;    \draw[fill=\couleur!\thedensity] (A) -- (B) -- (C) -- cycle;&#10;    \foreach \x in {1,...,15}{%&#10;        \pgfmathsetcounter{density}{\thedensity+10}&#10;        \setcounter{density}{\thedensity}&#10;        \path[coordinate] coordinate(X) at (A){};&#10;        \path[coordinate] (A) -- (B) coordinate[pos=.15](A)&#10;                            -- (C) coordinate[pos=.15](B)&#10;                            -- (X) coordinate[pos=.15](C);&#10;        \draw[fill=\couleur!\thedensity] (A)--(B)--(C)--cycle;&#10;    }&#10;\end{tikzpicture}&#10;$$" src="svgs/68bfc12202ea59bca6c4fe8f73cab682.svg" align="middle" width="2514.93pt" height="16.376943pt"/></p>