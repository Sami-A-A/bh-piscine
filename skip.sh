#! /bin/bash
ls -1 | awk '{ if(NR%2!=0){ print "-rw-r--r-- 1 1000 1000 0 Feb 28 12:57 " $0 } }'