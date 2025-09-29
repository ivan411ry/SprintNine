package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("size <= 0 -> empty slice", func(t *testing.T) {
		for _, sz := range []int{0, -1, -10} {
			got := generateRandomElements(sz)
			if got == nil {
				t.Fatalf("got nil slice for size=%d; want non-nil empty slice", sz)
			}
			if len(got) != 0 {
				t.Fatalf("len(generateRandomElements(%d)) = %d; want 0", sz, len(got))
			}
		}
	})

	t.Run("size > 0 -> slice of exact length", func(t *testing.T) {
		for _, sz := range []int{1, 5, 100} {
			got := generateRandomElements(sz)
			if got == nil {
				t.Fatalf("got nil slice for size=%d; want non-nil", sz)
			}
			if len(got) != sz {
				t.Fatalf("len(generateRandomElements(%d)) = %d; want %d", sz, len(got), sz)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{5}, 5},
		{"all equal", []int{7, 7, 7}, 7},
		{"increasing", []int{1, 2, 3, 4}, 4},
		{"decreasing", []int{9, 7, 5, 3}, 9},
		{"mixed signs", []int{-10, 0, 10}, 10},
		{"all negative", []int{-5, -2, -10}, -2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.in)
			if got != tt.want {
				t.Fatalf("maximum(%v) = %d; want %d", tt.in, got, tt.want)
			}
		})
	}
}
