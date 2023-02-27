#! /bin/bash

find . -type f -name "*.sh"| sort -r | sed 's/\.sh$//1' | sed 's|^./||'
