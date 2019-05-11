#!/bin/bash -e
# render.sh

echo " "
echo "************************************************************************"
echo "**************************************************** render.sh (START) *"
echo " "

echo "pdflatex latex-tkz-euclide.txt"
pdflatex latex-tkz-euclide.txt
echo " "

echo "****************************************************** render.sh (END) *"
echo "************************************************************************"
echo " "