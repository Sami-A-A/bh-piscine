#! /bin/bash

find . -name "*.sh" | sort -r | sed 's/\.sh$//1' | tr -d "./"