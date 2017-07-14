#!/bin/bash

echo "$(</dev/stdin)" > tmp.html

Xvfb :0 -screen 0 2048x2048x24 &
cat tmp.html | webshot




