#!/bin/bash

if [ "$1" == "onlp" ]; then
	echo "Build Target is onlp"
elif [ "$1" == "openBMC" ]; then
	echo "Build Target is OpenBMC"
elif [ "$1" == "openBMCVoyager" ]; then
	echo "Build Target is OpenBMCVoyager"
else
	echo "Build Target is Dummy"
fi
