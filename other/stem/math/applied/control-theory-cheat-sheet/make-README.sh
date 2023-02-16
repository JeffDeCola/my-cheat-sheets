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
python3 -m readme2tex --usepackage "xcolor" --readme README.tex.md --output README.md --nocdn
echo " "

# readme2tex is a little broke, this is the hack i used to put quotes around align, width and height
sed -i 's/align=middle width=/align="middle" width="/' README.md
sed -i 's/pt height=/pt" height="/' README.md
sed -i 's/pt\/>/pt" \/>/' README.md

echo "For more information on readme2tex, read my cheat sheet"
echo "    https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet"
echo " "

echo "************************************************* make-README.sh (END) *"
echo "************************************************************************"
echo " "
