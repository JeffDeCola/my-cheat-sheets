# COMBINATIONS CHEAT SHEET

```txt
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

`combinations` _tbd._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

Einstein's famous equation,

$$
E=mc^2
$$

\begin{tikzpicture}
\newcounter{density}
\setcounter{density}{20}
    \def\couleur{red}
    \path[coordinate] (0,0)  coordinate(A)
                ++( 60:6cm) coordinate(B)
                ++(-60:6cm) coordinate(C);
    \draw[fill=\couleur!\thedensity] (A) -- (B) -- (C) -- cycle;
    \foreach \x in {1,...,15}{%
        \pgfmathsetcounter{density}{\thedensity+10}
        \setcounter{density}{\thedensity}
        \path[coordinate] coordinate(X) at (A){};
        \path[coordinate] (A) -- (B) coordinate[pos=.15](A)
                            -- (C) coordinate[pos=.15](B)
                            -- (X) coordinate[pos=.15](C);
        \draw[fill=\couleur!\thedensity] (A)--(B)--(C)--cycle;
    }
\end{tikzpicture}
