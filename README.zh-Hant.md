# [PixelDiff](https://github.com/kagurazakayashi/PixelDiff)

[English](README.md) | [简体中文](README.zh-Hans.md) | **繁體中文** | [日本語](README.ja.md)

`PixelDiff` 是一款輕量級命令列工具，用於偵測圖片中指定座標點的色彩值，並計算其與目標 RGB 色彩的量化差異程度。

此工具非常適合用於 UI 自動化測試驗證、AI 助理呼叫、影像處理除錯，或簡易的視覺監控場景。

## 功能

- **精準定位**：支援指定圖片中的任意像素座標。
- **色彩比對**：在命令列輸入目標 RGB 值，輸出實際色彩值與目標色彩之間的偏差。
- **差異量化**：透過歐幾里得距離演算法，輸出色彩差異。

### 原理

使用 3D 歐幾里得距離（Euclidean Distance）演算法，計算兩種色彩在 RGB 色彩空間中的直線距離。

$$Distance = \sqrt{(R_1 - R_2)^2 + (G_1 - G_2)^2 + (B_1 - B_2)^2}$$

最終的 **差異值** 是將該距離除以色彩空間中可能的最大距離，也就是黑色 `(0,0,0)` 與白色 `(255,255,255)` 之間的距離，約等於 `441.67`，計算而得。

## 使用方式

`./pixeldiff -i [圖片路徑] -x [X座標] -y [Y座標] -r [紅色] -g [綠色] -b [藍色]`

### 參數說明

| 參數 | 說明                                 | 範例       |
| ---- | ------------------------------------ | ---------- |
| `-i` | 圖片檔案的絕對或相對路徑             | `test.png` |
| `-x` | 目標像素的 X 座標（從 `0` 開始）     | `100`      |
| `-y` | 目標像素的 Y 座標（從 `0` 開始）     | `200`      |
| `-r` | 目標色彩的 R（Red）值（`0`-`255`）   | `255`      |
| `-g` | 目標色彩的 G（Green）值（`0`-`255`） | `0`        |
| `-b` | 目標色彩的 B（Blue）值（`0`-`255`）  | `0`        |

### 輸出結果範例

`0.363805`

數值越接近 `1`，代表差異越大。

## 安裝說明

從 Release 下載後即可使用。

## 編譯

首先，請確認你已安裝 [Go 語言環境](https://go.dev/doc/install)。

1. 複製或下載本專案
2. 在專案目錄下初始化並編譯：

`go mod tidy && go build .`

## LICENSE

```LICENSE
Copyright (c) 2026 KagurazakaYashi
PixelDiff is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
```
