package main

import (
	"math"
	"testing"
)

func TestRgbaTo8Bit(t *testing.T) {
	tests := []struct {
		name     string
		r, g, b  uint32
		wantR, wantG, wantB int
	}{
		{"zero values", 0, 0, 0, 0, 0, 0},
		{"max 16-bit white", 65535, 65535, 65535, 255, 255, 255},
		{"mid value (128)", 32768, 32768, 32768, 128, 128, 128},
		{"pure red 8-bit equivalent", 65535, 0, 0, 255, 0, 0},
		{"pure green 8-bit equivalent", 0, 65535, 0, 0, 255, 0},
		{"pure blue 8-bit equivalent", 0, 0, 65535, 0, 0, 255},
		{"128 16-bit maps to 0 after shift", 128, 128, 128, 0, 0, 0},
		{"256 16-bit maps to 1 after shift", 256, 256, 256, 1, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB := rgbaTo8Bit(tt.r, tt.g, tt.b)
			if gotR != tt.wantR || gotG != tt.wantG || gotB != tt.wantB {
				t.Errorf("rgbaTo8Bit(%d, %d, %d) = (%d, %d, %d), want (%d, %d, %d)",
					tt.r, tt.g, tt.b, gotR, gotG, gotB, tt.wantR, tt.wantG, tt.wantB)
			}
		})
	}
}

func TestColorDistance(t *testing.T) {
	tests := []struct {
		name             string
		r1, g1, b1       int
		r2, g2, b2       int
		want             float64
	}{
		{"identical colors (black)", 0, 0, 0, 0, 0, 0, 0},
		{"identical colors (white)", 255, 255, 255, 255, 255, 255, 0},
		{"identical colors (red)", 255, 0, 0, 255, 0, 0, 0},
		{"black vs white", 0, 0, 0, 255, 255, 255, maxColorDistance},
		{"red vs green", 255, 0, 0, 0, 255, 0, math.Sqrt(255*255 + 255*255)},
		{"single channel diff (R)", 0, 0, 0, 100, 0, 0, 100},
		{"single channel diff (G)", 0, 0, 0, 0, 100, 0, 100},
		{"single channel diff (B)", 0, 0, 0, 0, 0, 100, 100},
		{"50,50,50 vs 150,150,150", 50, 50, 50, 150, 150, 150, math.Sqrt(3 * 100.0 * 100.0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := colorDistance(tt.r1, tt.g1, tt.b1, tt.r2, tt.g2, tt.b2)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("colorDistance(%d,%d,%d, %d,%d,%d) = %v, want %v",
					tt.r1, tt.g1, tt.b1, tt.r2, tt.g2, tt.b2, got, tt.want)
			}
		})
	}
}

func TestNormalizedColorDiff(t *testing.T) {
	tests := []struct {
		name       string
		r1, g1, b1 int
		r2, g2, b2 int
		want       float64
	}{
		{"identical colors", 0, 0, 0, 0, 0, 0, 0},
		{"identical colors non-black", 128, 64, 32, 128, 64, 32, 0},
		{"max difference (black vs white)", 0, 0, 0, 255, 255, 255, 1.0},
		{"max difference (red vs cyan)", 255, 0, 0, 0, 255, 255, 1.0},
		{"half distance on single channel", 0, 0, 0, 128, 0, 0, 128.0 / maxColorDistance},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizedColorDiff(tt.r1, tt.g1, tt.b1, tt.r2, tt.g2, tt.b2)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("normalizedColorDiff(%d,%d,%d, %d,%d,%d) = %v, want %v",
					tt.r1, tt.g1, tt.b1, tt.r2, tt.g2, tt.b2, got, tt.want)
			}
		})
	}
}

func TestNormalizedColorDiffRange(t *testing.T) {
	// Verify the result is always in [0, 1] for all 8-bit color pairs.
	colors := []int{0, 64, 128, 192, 255}
	for _, r1 := range colors {
		for _, g1 := range colors {
			for _, b1 := range colors {
				for _, r2 := range colors {
					for _, g2 := range colors {
						for _, b2 := range colors {
							got := normalizedColorDiff(r1, g1, b1, r2, g2, b2)
							if got < 0 || got > 1.0 {
								t.Errorf("normalizedColorDiff(%d,%d,%d, %d,%d,%d) = %v, out of [0,1] range",
									r1, g1, b1, r2, g2, b2, got)
								return
							}
						}
					}
				}
			}
		}
	}
}

func TestColorDistanceSymmetry(t *testing.T) {
	// colorDistance should be symmetric.
	cases := [][6]int{
		{0, 0, 0, 255, 255, 255},
		{50, 100, 150, 200, 100, 50},
		{255, 0, 0, 0, 255, 0},
		{10, 20, 30, 40, 50, 60},
	}
	for _, c := range cases {
		d1 := colorDistance(c[0], c[1], c[2], c[3], c[4], c[5])
		d2 := colorDistance(c[3], c[4], c[5], c[0], c[1], c[2])
		if d1 != d2 {
			t.Errorf("colorDistance not symmetric: d(%v,%v,%v → %v,%v,%v) = %v, reverse = %v",
				c[0], c[1], c[2], c[3], c[4], c[5], d1, d2)
		}
	}
}

func TestMaxColorDistanceConstant(t *testing.T) {
	expected := math.Sqrt(3 * 255.0 * 255.0)
	if math.Abs(maxColorDistance-expected) > 1e-9 {
		t.Errorf("maxColorDistance = %v, want %v", maxColorDistance, expected)
	}
}
