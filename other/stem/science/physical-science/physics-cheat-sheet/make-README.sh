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
echo "Run readme2tex on motion.tex.md"
python -m readme2tex --usepackage "xcolor" --readme motion.tex.md --output motion.md --nocdn
echo "Run readme2tex on energy.tex.md"
python -m readme2tex --usepackage "xcolor" --readme energy.tex.md --output energy.md --nocdn
echo "Run readme2tex on waves.tex.md"
python -m readme2tex --usepackage "xcolor" --readme waves.tex.md --output waves.md --nocdn
echo "Run readme2tex on electricity-and-magnetism.tex.md"
python -m readme2tex --usepackage "xcolor" --readme electricity-and-magnetism.tex.md --output electricity-and-magnetism.md --nocdn
echo "Run readme2tex on light-and-optics.tex.md"
python -m readme2tex --usepackage "xcolor" --readme light-and-optics.tex.md --output light-and-optics.md --nocdn
echo " "

echo "For more information on readme2tex, read my cheat sheet"
echo "    https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/languages/latex-cheat-sheet#how-i-created-this-readmemd"
echo " "

echo "************************************************* make-README.sh (END) *"
echo "************************************************************************"
echo " "
