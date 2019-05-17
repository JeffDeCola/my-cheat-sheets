#!/bin/bash -e
# make-README.sh

echo " "
echo "************************************************************************"
echo "*********************************************** make-README.sh (START) *"
echo " "

echo "rm -r svgs/"
rm -r svgs/
echo " "

echo "Run readme2tex on README.tex.md"
python -m readme2tex --usepackage "xcolor" --readme README.tex.md --output README.md --nocdn
echo " "

echo "Run readme2tex on latex-math-mode.tex.md"
python -m readme2tex --usepackage "xcolor" --readme latex-math-mode.tex.md --output latex-math-mode.md --nocdn
echo " "

echo "Run readme2tex on latex-graphs.tex.md"
python -m readme2tex --usepackage "xcolor" --readme latex-graphs.tex.md --output latex-graphs.md --nocdn
echo " "

echo "For more information on readme2tex, read my cheat sheet"
echo "    https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-math-mode-cheat-sheet#app---readme2tex"
echo " "

echo "************************************************* make-README.sh (END) *"
echo "************************************************************************"
echo " "
