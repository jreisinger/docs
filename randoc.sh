#!/bin/bash

NDOCS=$(find . -type f -iname '*.md' | wc -l)
N=$(shuf -i 1-"$NDOCS" -n 1)
DOC=$(find . -type f | sed -n "${N}"p)
echo "$DOC"