# [PixelDiff](https://github.com/kagurazakayashi/PixelDiff)

[English](README.md) | **简体中文** | [繁體中文](README.zh-Hant.md) | [日本語](README.ja.md)

`PixelDiff` 是一个轻量级命令行工具，用于检测图片中指定坐标点的颜色值，并计算其与目标 RGB 颜色的量化差异度。

该工具非常适合用于 UI 自动化测试校验、 AI 助理调用、图像处理调试或简单的视觉监控场景。

## 功能

- **精准定位**：支持指定图片的任意像素坐标。
- **色彩对比**：在命令行输入目标 RGB 值，输出计算实际色值与目标的偏差。
- **差异量化**：通过欧几里得距离算法，输出色彩差异。

### 原理

使用 3D 欧几里得距离 (Euclidean Distance) 算法来计算两种颜色在 RGB 色彩空间中的直线距离。

$$Distance = \sqrt{(R_1 - R_2)^2 + (G_1 - G_2)^2 + (B_1 - B_2)^2}$$

最终的 **差异值** 是将该距离除以色彩空间中可能的最大距离（即黑色 `(0,0,0)` 与白色 `(255,255,255)` 之间的距离，约等于 `441.67`）计算得出的。

## 使用方法

`./pixeldiff -i [图片路径] -x [X坐标] -y [Y坐标] -r [红] -g [绿] -b [蓝]`

### 参数说明

| 参数 | 说明                                | 示例       |
| ---- | ----------------------------------- | ---------- |
| `-i` | 图片文件的绝对或相对路径            | `test.png` |
| `-x` | 目标像素的 X 坐标（从 `0` 开始）    | `100`      |
| `-y` | 目标像素的 Y 坐标（从 `0` 开始）    | `200`      |
| `-r` | 目标颜色的 R (Red) 值 (`0`-`255`)   | `255`      |
| `-g` | 目标颜色的 G (Green) 值 (`0`-`255`) | `0`        |
| `-b` | 目标颜色的 B (Blue) 值 (`0`-`255`)  | `0`        |

### 输出结果示例

`0.363805`

数值越接近于 `1` 差异越大。

## 安装说明

从 [Release](https://github.com/kagurazakayashi/PixelDiff/releases) 下载使用即可。

## 编译

首先确保你已经安装了 [Go 语言环境](https://go.dev/doc/install)。

1. 克隆或下载本项目
2. 在项目目录下初始化并编译：

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
