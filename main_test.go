package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "zero size", size: 0, wantLen: 0},
		{name: "negative size", size: -10, wantLen: 0},
		{name: "single element", size: 1, wantLen: 1},
		{name: "many elements", size: 1000, wantLen: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if len(data) != tt.wantLen {
				t.Fatalf("len(generateRandomElements(%d)) = %d, want %d", tt.size, len(data), tt.wantLen)
			}
			for i, value := range data {
				if value < 0 {
					t.Fatalf("generateRandomElements(%d)[%d] = %d, want non-negative value", tt.size, i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{42}, want: 42},
		{name: "single zero", data: []int{0}, want: 0},
		{name: "max at start", data: []int{100, 1, 2, 3}, want: 100},
		{name: "max in middle", data: []int{1, 100, 2, 3}, want: 100},
		{name: "max at end", data: []int{1, 2, 3, 100}, want: 100},
		{name: "all equal", data: []int{7, 7, 7, 7}, want: 7},
		{name: "duplicates", data: []int{5, 9, 9, 2}, want: 9},
		{name: "negative values", data: []int{-5, -1, -3}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{42}, want: 42},
		{name: "fewer than chunks", data: []int{3, 1, 2}, want: 3},
		{name: "exactly one chunk", data: []int{1, 8, 2, 3, 7, 4, 5, 6}, want: 8},
		{name: "chunks with remainder", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 100}, want: 100},
		{name: "max in last chunk", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 900, 17}, want: 900},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Fatalf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	sizes := []int{0, 1, 2, 7, 8, 9, 16, 17, 1000, 100_000}

	for _, size := range sizes {
		data := generateRandomElements(size)
		want := maximum(data)
		if got := maxChunks(data); got != want {
			t.Fatalf("maxChunks(size=%d) = %d, maximum = %d", size, got, want)
		}
	}
}
