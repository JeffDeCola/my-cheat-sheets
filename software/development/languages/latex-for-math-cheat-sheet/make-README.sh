#!/bin/bash -e
# make-README.sh

echo " "
echo "************************************************************************"
echo "*********************************************** make-README.sh (START) *"
echo " "

echo "rm -r svgs/"
rm -r svgs/
echo " "

echo "Run readme2tex"
python -m readme2tex --readme README.tex.md --output README.md --nocdn
echo " "

echo "************************************************* make-README.sh (END) *"
echo "************************************************************************"
echo " "
