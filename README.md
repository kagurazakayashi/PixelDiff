# [PixelDiff](https://github.com/kagurazakayashi/PixelDiff)

**English** | [简体中文](README.zh-Hans.md) | [繁體中文](README.zh-Hant.md) | [日本語](README.ja.md)

`PixelDiff` is a lightweight command-line tool for detecting the color value of a specified pixel coordinate in an image and calculating the quantified difference between that color and a target RGB color.

This tool is well suited for UI automation test verification, AI assistant calls, image processing debugging, or simple visual monitoring scenarios.

## Features

- **Precise positioning**: Supports specifying any pixel coordinate in an image.
- **Color comparison**: Enter a target RGB value in the command line and output the deviation between the actual color value and the target color.
- **Quantified difference**: Outputs the color difference using the Euclidean distance algorithm.

### Principle

It uses the 3D Euclidean Distance algorithm to calculate the straight-line distance between two colors in the RGB color space.

$$Distance = \sqrt{(R_1 - R_2)^2 + (G_1 - G_2)^2 + (B_1 - B_2)^2}$$

The final **difference value** is calculated by dividing this distance by the maximum possible distance in the color space, namely the distance between black `(0,0,0)` and white `(255,255,255)`, which is approximately `441.67`.

## Usage

`./pixeldiff -i [image path] -x [X coordinate] -y [Y coordinate] -r [red] -g [green] -b [blue]`

### Parameters

| Parameter | Description                                         | Example    |
| --------- | --------------------------------------------------- | ---------- |
| `-i`      | Absolute or relative path to the image file         | `test.png` |
| `-x`      | X coordinate of the target pixel, starting from `0` | `100`      |
| `-y`      | Y coordinate of the target pixel, starting from `0` | `200`      |
| `-r`      | R (Red) value of the target color (`0`-`255`)       | `255`      |
| `-g`      | G (Green) value of the target color (`0`-`255`)     | `0`        |
| `-b`      | B (Blue) value of the target color (`0`-`255`)      | `0`        |

### Example Output

`0.363805`

The closer the value is to `1`, the greater the difference.

## Installation

Download it from the Release page and use it directly.

## Build

First, make sure you have installed the [Go environment](https://go.dev/doc/install).

1. Clone or download this project.
2. Initialize and build it in the project directory:

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
