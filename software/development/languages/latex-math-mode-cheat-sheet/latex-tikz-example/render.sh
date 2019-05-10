#!/bin/bash -e
# render.sh

echo " "
echo "************************************************************************"
echo "**************************************************** render.sh (START) *"
echo " "

echo "pdflatex latex-tikzpicture.txt"
pdflatex latex-tikzpicture.txt
echo " "

echo "****************************************************** render.sh (END) *"
echo "************************************************************************"
echo " "