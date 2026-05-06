# [PixelDiff](https://github.com/kagurazakayashi/PixelDiff)

[English](README.md) | [简体中文](README.zh-Hans.md) | [繁體中文](README.zh-Hant.md) | **日本語**

`PixelDiff` は、画像内の指定した座標点の色値を検出し、対象の RGB 色との差異を定量的に計算する軽量なコマンドラインツールです。

このツールは、UI 自動化テストの検証、AI アシスタントからの呼び出し、画像処理のデバッグ、または簡易的なビジュアル監視の用途に適しています。

## 機能

- **高精度な位置指定**：画像内の任意のピクセル座標を指定できます。
- **色比較**：コマンドラインで対象の RGB 値を入力し、実際の色値と対象色との差分を出力します。
- **差異の定量化**：ユークリッド距離アルゴリズムにより、色の差異を出力します。

### 原理

3D ユークリッド距離（Euclidean Distance）アルゴリズムを使用して、RGB 色空間における 2 つの色の直線距離を計算します。

$$Distance = \sqrt{(R_1 - R_2)^2 + (G_1 - G_2)^2 + (B_1 - B_2)^2}$$

最終的な **差異値** は、この距離を色空間内で取り得る最大距離、つまり黒 `(0,0,0)` と白 `(255,255,255)` の間の距離（約 `441.67`）で割ることで算出されます。

## 使用方法

`./pixeldiff -i [画像パス] -x [X座標] -y [Y座標] -r [赤] -g [緑] -b [青]`

### パラメーター説明

| パラメーター | 説明                                  | 例         |
| ------------ | ------------------------------------- | ---------- |
| `-i`         | 画像ファイルの絶対パスまたは相対パス  | `test.png` |
| `-x`         | 対象ピクセルの X 座標（`0` から開始） | `100`      |
| `-y`         | 対象ピクセルの Y 座標（`0` から開始） | `200`      |
| `-r`         | 対象色の R（Red）値（`0`-`255`）      | `255`      |
| `-g`         | 対象色の G（Green）値（`0`-`255`）    | `0`        |
| `-b`         | 対象色の B（Blue）値（`0`-`255`）     | `0`        |

### 出力結果の例

`0.363805`

値が `1` に近いほど、差異が大きいことを示します。

## インストール

Release からダウンロードしてそのまま使用できます。

## ビルド

まず、[Go 言語環境](https://go.dev/doc/install) がインストールされていることを確認してください。

1. このプロジェクトをクローンまたはダウンロードします。
2. プロジェクトディレクトリで初期化してビルドします。

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
