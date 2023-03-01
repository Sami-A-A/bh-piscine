#! /bin/bash
ls -1 | awk -v now="$(date +"%b  1 %H:%M ")" '{ if(NR%2!=0){ print "-rw-r--r-- 1 1000 1000 0 " now  $0}}' 