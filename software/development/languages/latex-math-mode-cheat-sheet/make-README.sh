#!/bin/bash -e
# make-README.sh

echo " "
echo "************************************************************************"
echo "*********************************************** make-README.sh (START) *"
echo " "

echo "rm -r svgs/"
rm -r svgs/
echo " "

echo "Run readme2tex"
python -m readme2tex --usepackage "xcolor" --usepackage "booktabs" --readme README.tex.md --output README.md --nocdn
echo " "

echo "For more information on readme2tex, read my cheat sheet"
echo "    https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-math-mode-cheat-sheet#app---readme2tex"
echo " "

echo "************************************************* make-README.sh (END) *"
echo "************************************************************************"
echo " "
