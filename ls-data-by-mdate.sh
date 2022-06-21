#!/bin/bash

git ls-files -z -- "data" | \
xargs -0 -n1 -I{} -- git log -1 --format="%at {}" {} | \
sort -r | \
perl -MTime::Piece -lape 's/(^\d+)/localtime($1)->strftime("%F")/e'