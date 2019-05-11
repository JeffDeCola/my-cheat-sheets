#!/bin/bash -e
# y-equals-sqrt-x.sh

echo " "
echo "************************************************************************"
echo "******************************************* y-equals-sqrt-x.sh (START) *"
echo " "

echo "Create .dvi file"
echo "latex y-equals-sqrt-x.tex"
latex y-equals-sqrt-x.tex
echo " "

echo "Convert .dvi to .svg"
echo "    -n no fonts"
echo "    -a This option forces dvisvgm to vectorize everything"
dvisvgm -n -a y-equals-sqrt-x.dvi
echo " "

echo "********************************************* y-equals-sqrt-x.sh (END) *"
echo "************************************************************************"
echo " "
