#!/bin/bash
go test . -v
if [ ! -f "test.png" ]; then
    magick "ico/icon.ico" "test.png"
    mv "test-0.png" "test.png"
    rm -f test-*.png
fi
if [ -f "./pixeldiff" ]; then
    rm -f ./pixeldiff
fi
go build .
chmod +x ./pixeldiff
./pixeldiff -i "test.png" -x 100 -y 100 -r 255 -g 0 -b 0
./pixeldiff -i "test.png" -x 100 -y 100 -r 0 -g 255 -b 0
./pixeldiff -i "test.png" -x 100 -y 100 -r 0 -g 0 -b 255
rm -f ./pixeldiff
