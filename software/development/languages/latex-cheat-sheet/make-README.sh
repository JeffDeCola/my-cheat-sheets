#!/bin/sh -e
# make-README.sh

echo " "
echo "************************************************************************"
echo "*********************************************** make-README.sh (START) *"
echo " "

echo "rm -r svgs/"
rm -r svgs/
echo " "

echo "Run readme2tex on README.tex.md"
python3 -m readme2tex --usepackage "xcolor" --readme README.tex.md --output README.md --nocdn
echo " "

# readme2tex is a little broke, this is the hack i used to put quotes around align, width and height
sed -i 's/align=middle width=/align="middle" width="/' README.md
sed -i 's/pt height=/pt" height="/' README.md
sed -i 's/pt\/>/pt" \/>/' README.md

echo "Run readme2tex on latex-math-mode.tex.md"
python3 -m readme2tex --usepackage "xcolor" --readme latex-math-mode.tex.md --output latex-math-mode.md --nocdn
echo " "

# readme2tex is a little broke, this is the hack i used to put quotes around align, width and height
sed -i 's/align=middle width=/align="middle" width="/' latex-math-mode.md
sed -i 's/pt height=/pt" height="/' latex-math-mode.md
sed -i 's/pt\/>/pt" \/>/' latex-math-mode.md

echo "Run readme2tex on latex-graphs.tex.md"
python3 -m readme2tex --usepackage "xcolor" --readme latex-graphs.tex.md --output latex-graphs.md --nocdn
echo " "

# readme2tex is a little broke, this is the hack i used to put quotes around align, width and height
sed -i 's/align=middle width=/align="middle" width="/' latex-graphs.md
sed -i 's/pt height=/pt" height="/' latex-graphs.md
sed -i 's/pt\/>/pt" \/>/' latex-graphs.md

echo "For more information on readme2tex, read my cheat sheet"
echo "    https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet"
echo " "

echo "************************************************* make-README.sh (END) *"
echo "************************************************************************"
echo " "
