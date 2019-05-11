#!/bin/bash -e
# render.sh

echo " "
echo "************************************************************************"
echo "**************************************************** render.sh (START) *"
echo " "

echo "pdflatex latex-pgfplots.tex"
pdflatex latex-pgfplots.tex
echo " "

echo "****************************************************** render.sh (END) *"
echo "************************************************************************"
echo " "