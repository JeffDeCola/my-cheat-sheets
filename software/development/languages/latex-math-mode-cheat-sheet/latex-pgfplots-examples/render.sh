#!/bin/bash -e
# render.sh

echo " "
echo "************************************************************************"
echo "**************************************************** render.sh (START) *"
echo " "

echo "pdflatex latex-pgfplots.txt"
pdflatex latex-pgfplots.txt
echo " "

echo "****************************************************** render.sh (END) *"
echo "************************************************************************"
echo " "