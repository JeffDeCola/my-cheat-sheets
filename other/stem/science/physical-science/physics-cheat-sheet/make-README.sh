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
python3 -m readme2tex --usepackage "xcolor" --readme README.tex.md --output README.md --nocdn
# readme2tex is a little broke, this is the hack i used to put quotes around align, width and height
sed -i 's/align=middle width=/align="middle" width="/' README.md
sed -i 's/pt height=/pt" height="/' README.md
sed -i 's/pt\/>/pt" \/>/' README.md

echo "Run readme2tex on motion.tex.md"
python3 -m readme2tex --usepackage "xcolor" --readme motion.tex.md --output motion.md --nocdn
# readme2tex is a little broke, this is the hack i used to put quotes around align, width and height
sed -i 's/align=middle width=/align="middle" width="/' motion.md
sed -i 's/pt height=/pt" height="/' motion.md
sed -i 's/pt\/>/pt" \/>/' motion.md

echo "Run readme2tex on energy.tex.md"
python3 -m readme2tex --usepackage "xcolor" --readme energy.tex.md --output energy.md --nocdn
# readme2tex is a little broke, this is the hack i used to put quotes around align, width and height
sed -i 's/align=middle width=/align="middle" width="/' energy.md
sed -i 's/pt height=/pt" height="/' energy.md
sed -i 's/pt\/>/pt" \/>/' energy.md

echo "Run readme2tex on waves.tex.md"
python3 -m readme2tex --usepackage "xcolor" --readme waves.tex.md --output waves.md --nocdn
# readme2tex is a little broke, this is the hack i used to put quotes around align, width and height
sed -i 's/align=middle width=/align="middle" width="/' waves.md
sed -i 's/pt height=/pt" height="/' waves.md
sed -i 's/pt\/>/pt" \/>/' waves.md

echo "Run readme2tex on electricity-and-magnetism.tex.md"
python3 -m readme2tex --usepackage "xcolor" --readme electricity-and-magnetism.tex.md --output electricity-and-magnetism.md --nocdn
# readme2tex is a little broke, this is the hack i used to put quotes around align, width and height
sed -i 's/align=middle width=/align="middle" width="/' electricity-and-magnetism.md
sed -i 's/pt height=/pt" height="/' electricity-and-magnetism.md
sed -i 's/pt\/>/pt" \/>/' electricity-and-magnetism.md

echo "Run readme2tex on light-and-optics.tex.md"
python3 -m readme2tex --usepackage "xcolor" --readme light-and-optics.tex.md --output light-and-optics.md --nocdn
# readme2tex is a little broke, this is the hack i used to put quotes around align, width and height
sed -i 's/align=middle width=/align="middle" width="/' light-and-optics.md
sed -i 's/pt height=/pt" height="/' light-and-optics.md
sed -i 's/pt\/>/pt" \/>/' light-and-optics.md

echo " "

echo "For more information on readme2tex, read my cheat sheet"
echo "    https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet"
echo " "

echo "************************************************* make-README.sh (END) *"
echo "************************************************************************"
echo " "