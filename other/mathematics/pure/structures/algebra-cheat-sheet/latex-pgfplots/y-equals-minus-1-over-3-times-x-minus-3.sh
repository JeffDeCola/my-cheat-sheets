#!/bin/bash -e
# y-equals-minus-1-over-3-times-x-minus-3.sh

echo " "
echo "************************************************************************"
echo "******************* y-equals-minus-1-over-3-times-x-minus-3.sh (START) *"
echo " "

echo "Create .dvi file"
echo "latex  y-equals-minus-1-over-3-times-x-minus-3.tex"
latex  y-equals-minus-1-over-3-times-x-minus-3.tex
echo " "

echo "Convert .dvi to .svg"
echo "    -n no fonts"
echo "    -a This option forces dvisvgm to vectorize everything"
dvisvgm -n -a  y-equals-minus-1-over-3-times-x-minus-3.dvi
echo " "

echo "********************* y-equals-minus-1-over-3-times-x-minus-3.sh (END) *"
echo "************************************************************************"
echo " "
