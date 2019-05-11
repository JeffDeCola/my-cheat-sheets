#!/bin/bash -e
# make-README.sh

echo " "
echo "************************************************************************"
echo "*********************************************** make-README.sh (START) *"
echo " "

echo "Create .dvi file"
echo "latex y-equal-2x.tex"
latex y-equal-2x.tex
echo " "

echo "Convert .dvi to .svg"
echo "    -n no fonts"
echo "    -a This option forces dvisvgm to vectorize everything"
dvisvgm -n -a y-equal-2x.dvi
echo " "

echo "************************************************* make-README.sh (END) *"
echo "************************************************************************"
echo " "
