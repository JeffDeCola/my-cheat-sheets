#!/bin/bash -e
# render.sh

echo " "
echo "************************************************************************"
echo "**************************************************** render.sh (START) *"
echo " "

echo "pdflatex latex-tkz-euclide.tex"
pdflatex latex-tkz-euclide.tex
echo " "

echo "****************************************************** render.sh (END) *"
echo "************************************************************************"
echo " "