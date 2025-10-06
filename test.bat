go test . -v
IF NOT EXIST "test.png" (
    magick "ico\icon.ico" "test.png"
    RENAME "test-0.png" "test.png"
    DEL "test-*.png"
)
IF EXIST "PixelDiff.exe" (
    DEL PixelDiff.exe
)
go build .
PixelDiff.exe -i "test.png" -x 100 -y 100 -r 255 -g 0 -b 0
PixelDiff.exe -i "test.png" -x 100 -y 100 -r 0 -g 255 -b 0
PixelDiff.exe -i "test.png" -x 100 -y 100 -r 0 -g 0 -b 255
DEL PixelDiff.exe
