#!/bin/bash -e
# y-equals-2x.sh

echo " "
echo "************************************************************************"
echo "*********************************************** y-equals-2x.sh (START) *"
echo " "

echo "Create .dvi file"
echo "latex y-equals-2x.tex"
latex y-equals-2x.tex
echo " "

echo "Convert .dvi to .svg"
echo "    -n no fonts"
echo "    -a This option forces dvisvgm to vectorize everything"
dvisvgm -n -a y-equals-2x.dvi
echo " "

echo "************************************************* y-equals-2x.sh (END) *"
echo "************************************************************************"
echo " "
