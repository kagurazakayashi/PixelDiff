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

// rgbaTo8Bit 將 16 位元 RGBA 值 (0-65535) 轉換為 8 位元 (0-255)。
func rgbaTo8Bit(r, g, b uint32) (int, int, int) {
	return int(r >> 8), int(g >> 8), int(b >> 8)
}

// colorDistance 計算兩個 RGB 顏色之間的歐幾里得距離。
func colorDistance(r1, g1, b1, r2, g2, b2 int) float64 {
	return math.Sqrt(
		math.Pow(float64(r1-r2), 2) +
			math.Pow(float64(g1-g2), 2) +
			math.Pow(float64(b1-b2), 2),
	)
}

// maxColorDistance 是兩個 RGB 顏色之間的最大可能歐幾里得距離。
// √(255² + 255² + 255²) ≈ 441.67
const maxColorDistance = 441.6729559300637

// normalizedColorDiff 回傳兩個 RGB 顏色之間的標準化差異 (0.0 到 1.0)。
func normalizedColorDiff(r1, g1, b1, r2, g2, b2 int) float64 {
	return colorDistance(r1, g1, b1, r2, g2, b2) / maxColorDistance
}

func main() {
	// 定義命令列引數
	imgPath := flag.String("i", "", "image file path")
	targetX := flag.Int("x", 0, "target pixel X coordinate")
	targetY := flag.Int("y", 0, "target pixel Y coordinate")
	targetR := flag.Int("r", 0, "target R value (0-255)")
	targetG := flag.Int("g", 0, "target G value (0-255)")
	targetB := flag.Int("b", 0, "target B value (0-255)")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
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
	currR, currG, currB := rgbaTo8Bit(r, g, b)

	diffval := normalizedColorDiff(currR, currG, currB, *targetR, *targetG, *targetB)

	// 輸出結果
	fmt.Printf("%f", diffval)
}
