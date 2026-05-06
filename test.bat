go test . -v
IF NOT EXIST "test.png" (
    magick "ico\icon.ico" "test.png"
    RENAME "test-0.png" "test.png"
    DEL "test-*.png"
)
IF EXIST "pixeldiff.exe" (
    DEL pixeldiff.exe
)
go build . -o pixeldiff.exe
pixeldiff.exe -i "test.png" -x 100 -y 100 -r 255 -g 0 -b 0
pixeldiff.exe -i "test.png" -x 100 -y 100 -r 0 -g 255 -b 0
pixeldiff.exe -i "test.png" -x 100 -y 100 -r 0 -g 0 -b 255
DEL pixeldiff.exe
