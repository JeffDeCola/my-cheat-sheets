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
# python3 -m readme2tex --usepackage "xcolor" --readme README.tex.md --output README.md --nocdn
python3 -m readme2tex --readme README.tex.md --output README.md --nocdn
echo " "

echo "For more information on readme2tex, read my cheat sheet"
echo "    https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet"
echo " "

echo "************************************************* make-README.sh (END) *"
echo "************************************************************************"
echo " "
