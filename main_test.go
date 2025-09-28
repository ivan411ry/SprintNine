package main

// Пишите тесты в этом файле
import "testing"

func TestGenerateRandomElements_ZeroAndNegative(t *testing.T) {
	if got := generateRandomElements(0); len(got) != 0 {
		t.Fatalf("size=0: expected empty slice, got len=%d", len(got))
	}
	if got := generateRandomElements(-5); len(got) != 0 {
		t.Fatalf("size<0: expected empty slice, got len=%d", len(got))
	}
}
func TestGenerateRandomElements_SizeAndPositive(t *testing.T) {
	const n = 1000
	got := generateRandomElements(n)
	if len(got) != n {
		t.Fatalf("expected len=%d, got %d", n, len(got))
	}
	for i, v := range got {
		if v <= 0 {
			t.Fatalf("expected positive values, got %d at index %d", v, i)
		}
	}
}
func TestMaximum_Empty(t *testing.T) {
	if max := maximum(nil); max != 0 {
		t.Fatalf("empty slice: expected 0, got %d", max)
	}
	if max := maximum([]int{}); max != 0 {
		t.Fatalf("empty slice (non-nil): expected 0, got %d", max)
	}
}
func TestMaximum_Single(t *testing.T) {
	if max := maximum([]int{42}); max != 42 {
		t.Fatalf("single element: expected 42, got %d", max)
	}
}
func TestMaximum_Many(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{5, 1, 2}, 5},
		{[]int{1, 9, 3, 9, 2}, 9},
		{[]int{7, 7, 7}, 7},
		{[]int{1_000_000, 999_999}, 1_000_000},
	}
	for i, tc := range tests {
		if got := maximum(tc.in); got != tc.want {
			t.Fatalf("case %d: expected %d, got %d", i, tc.want, got)
		}
	}
}
