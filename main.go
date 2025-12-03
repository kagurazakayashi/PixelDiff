//go:generate goversioninfo -o=resource_windows_386.syso -64=false -icon=ico/icon.ico -manifest=main.exe.manifest
//go:generate goversioninfo -o=resource_windows_amd64.syso -64=true -icon=ico/icon.ico -manifest=main.exe.manifest
//go:generate goversioninfo -o=resource_windows_arm64.syso -arm=true -icon=ico/icon.ico -manifest=main.exe.manifest

package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

func main() {
	// 定義命令列引數
	imgPath := flag.String("path", "", "Image file path")
	targetX := flag.Int("x", 0, "X")
	targetY := flag.Int("y", 0, "Y")
	targetR := flag.Int("r", 0, "R(0-255)")
	targetG := flag.Int("g", 0, "G(0-255)")
	targetB := flag.Int("b", 0, "B(0-255)")
	flag.Parse()

	if *imgPath == "" {
		fmt.Println("NO IMAGE: -path <path>")
		return
	}

	// 開啟並解碼圖片
	file, err := os.Open(*imgPath)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// 獲取指定座標的色值
	// image.At 產生的 RGBA 是 16 位深 (0-65535)，需要右移 8 位轉回 0-255
	c := img.At(*targetX, *targetY)
	r, g, b, _ := c.RGBA()
	currR, currG, currB := int(r>>8), int(g>>8), int(b>>8)

	// 計算歐幾里得距離
	// 距離公式: √((r1-r2)² + (g1-g2)² + (b1-b2)²)
	diff := math.Sqrt(
		math.Pow(float64(currR-*targetR), 2) +
			math.Pow(float64(currG-*targetG), 2) +
			math.Pow(float64(currB-*targetB), 2),
	)

	// 標準化差異度 (最大可能距離為 √(255²+255²+255²) ≈ 441.67)
	maxDiff := math.Sqrt(3 * math.Pow(255, 2))
	diffval := diff / maxDiff

	// 輸出結果
	// fmt.Printf("%d, %d, %d ?\n", *targetR, *targetG, *targetB)
	// fmt.Printf("%d, %d = %d, %d, %d\n", *targetX, *targetY, currR, currG, currB)
	fmt.Printf("%f", diffval)
}
