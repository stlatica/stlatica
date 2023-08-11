#!/bin/bash

set -ue -o pipefail

# インポートステートメント内の文字列を抜き取る
IMPORTS=$(awk '/import \(/,/\)/' $1 | grep '"' | awk -F '"' '{print $2}')

# 抜き取った文字列でgo installを実行
for PKG in $IMPORTS; do
    echo "Installing $PKG..."
    go install $PKG
done
