#!/bin/bash
go test . -v
if [ ! -f "test.png" ]; then
    magick "ico/icon.ico" "test.png"
    mv "test-0.png" "test.png"
    rm -f test-*.png
fi
if [ -f "./PixelDiff" ]; then
    rm -f ./PixelDiff
fi
go build .
chmod +x ./PixelDiff
./PixelDiff -i "test.png" -x 100 -y 100 -r 255 -g 0 -b 0
./PixelDiff -i "test.png" -x 100 -y 100 -r 0 -g 255 -b 0
./PixelDiff -i "test.png" -x 100 -y 100 -r 0 -g 0 -b 255
rm -f ./PixelDiff
